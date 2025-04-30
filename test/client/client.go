package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "backend/proto/hotel"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewHotelServiceClient(conn)

	// Set context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Request to get all hotels with pagination
	request := &pb.GetAllHotelsRequest{
		Page:     1,
		PageSize: 10,
	}

	// Call the service
	response, err := client.GetAllHotels(ctx, request)
	if err != nil {
		log.Fatalf("Error calling GetAllHotels: %v", err)
	}

	// Display results
	fmt.Printf("Total hotels: %d\n", response.TotalCount)
	for i, hotel := range response.Hotels {
		fmt.Printf("[%d] - %s (Rating: %.1f)\n", i+1, hotel.Name, hotel.Rating)
		fmt.Printf("    Address: %s\n", hotel.FullAddress)
		fmt.Printf("    Price: $%.2f\n", hotel.Price)
		fmt.Printf("    Amenities: %v\n", hotel.Amenities)
		fmt.Printf("    Rooms: %d\n", len(hotel.Rooms))
		fmt.Println("-----------------------------------")
	}
} 