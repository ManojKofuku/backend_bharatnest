package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/models"
	pb "backend/proto/hotel"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

// HotelController handles hotel-related gRPC requests
type HotelController struct {
	pb.UnimplementedHotelServiceServer
	DB *gorm.DB
}

// NewHotelController creates a new hotel controller
func NewHotelController(db *gorm.DB) *HotelController {
	return &HotelController{DB: db}
}

// GetAllHotels retrieves all hotels with optional pagination
func (c *HotelController) GetAllHotels(ctx context.Context, req *pb.GetAllHotelsRequest) (*pb.GetAllHotelsResponse, error) {
	var hotels []models.Hotel
	var count int64

	// Query builder
	query := c.DB.Model(&models.Hotel{})

	// Get total count
	if err := query.Count(&count).Error; err != nil {
		log.Printf("Error counting hotels: %v", err)
		return nil, err
	}

	// Add pagination if provided
	pageSize := int(req.PageSize)
	if pageSize <= 0 {
		pageSize = 10 // Default page size
	}

	page := int(req.Page)
	if page <= 0 {
		page = 1 // Default page
	}

	offset := (page - 1) * pageSize

	// Fetch hotels with preloading
	if err := query.Offset(offset).Limit(pageSize).
		Preload("Rooms").
		Preload("Manager").
		Find(&hotels).Error; err != nil {
		log.Printf("Error fetching hotels: %v", err)
		return nil, err
	}

	// Convert to protobuf response
	pbHotels := make([]*pb.Hotel, len(hotels))
	for i, hotel := range hotels {
		pbHotel, err := hotelModelToProto(&hotel)
		if err != nil {
			log.Printf("Error converting hotel to proto: %v", err)
			continue
		}
		pbHotels[i] = pbHotel
	}

	return &pb.GetAllHotelsResponse{
		Hotels:     pbHotels,
		TotalCount: int32(count),
	}, nil
}

// hotelModelToProto converts a hotel model to protobuf message
func hotelModelToProto(hotel *models.Hotel) (*pb.Hotel, error) {
	// Parse amenities from JSON
	var amenitiesList []string
	if hotel.Amenities != nil {
		if err := json.Unmarshal(hotel.Amenities, &amenitiesList); err != nil {
			log.Printf("Error parsing amenities: %v", err)
			// Continue with empty amenities rather than failing
			amenitiesList = []string{}
		}
	}

	// Convert rooms
	rooms := make([]*pb.Room, len(hotel.Rooms))
	for i, room := range hotel.Rooms {
		rooms[i] = &pb.Room{
			Id:        uint64(room.ID),
			HotelId:   uint64(room.HotelID),
			RoomType:  room.RoomType,
			BedType:   room.BedType,
			Capacity:  int32(room.Capacity),
			Price:     room.Price,
			Available: room.Available,
			CreatedAt: timestamppb.New(room.CreatedAt),
			UpdatedAt: timestamppb.New(room.UpdatedAt),
		}
	}

	return &pb.Hotel{
		Id:                 uint64(hotel.ID),
		Name:               hotel.Name,
		LocationShort:      hotel.LocationShort,
		FullAddress:        hotel.FullAddress,
		PinCode:            hotel.PinCode,
		Landmark:           hotel.Landmark,
		Price:              hotel.Price,
		Rating:             hotel.Rating,
		Image:              hotel.Image,
		Availability:       hotel.Availability,
		Description:        hotel.Description,
		LocationLink:       hotel.LocationLink,
		Superhost:          hotel.Superhost,
		HostStory:          hotel.HostStory,
		CancellationPolicy: hotel.CancellationPolicy,
		CheckinTime:        timestamppb.New(hotel.CheckinTime),
		ManagerId:          uint64(hotel.ManagerID),
		Rooms:              rooms,
		Amenities:          amenitiesList,
		CreatedAt:          timestamppb.New(hotel.CreatedAt),
		UpdatedAt:          timestamppb.New(hotel.UpdatedAt),
	}, nil
}

// HTTP handlers for REST API

// GetAllHotelsHTTP handles the GET request to fetch all hotels
func GetAllHotelsHTTP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	var hotels []models.Hotel
	var count int64

	// Query builder
	query := db.Model(&models.Hotel{})

	// Get total count
	if err := query.Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count hotels"})
		return
	}

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// Fetch hotels with preloading
	if err := query.Offset(offset).Limit(pageSize).
		Preload("Rooms").
		Preload("Manager").
		Find(&hotels).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch hotels"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hotels":      hotels,
		"total_count": count,
	})
}

// GetHotelByIDHTTP handles the GET request to fetch a single hotel by its ID
func GetHotelByIDHTTP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	var hotel models.Hotel
	if err := db.Preload("Rooms").Preload("Manager").First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

// CreateHotelHTTP handles the POST request to create a new hotel
func CreateHotelHTTP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	var hotel models.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hotel"})
		return
	}

	c.JSON(http.StatusCreated, hotel)
}

// UpdateHotelHTTP handles the PUT request to update an existing hotel
func UpdateHotelHTTP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	var hotel models.Hotel
	if err := db.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}

	// Bind JSON to the hotel model
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the hotel
	if err := db.Save(&hotel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update hotel"})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

// DeleteHotelHTTP handles the DELETE request to remove a hotel
func DeleteHotelHTTP(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	if err := db.Delete(&models.Hotel{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
} 