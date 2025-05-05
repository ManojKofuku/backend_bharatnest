package main

import (
	"backend/controllers"
	"backend/initializers"
	pb "backend/proto/hotel" // Generated from protobuf
	"backend/routes"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Start gRPC server in a goroutine
	go startGRPCServer()

	// Start HTTP server
	r := gin.Default()
	
	// Add middleware to provide DB to all routes
	r.Use(func(c *gin.Context) {
		c.Set("db", initializers.DB)
		c.Next()
	})
	
	routes.RegisterUserRoutes(r)
	routes.RegisterHotelRoutes(r)
	
	r.Run("0.0.0.0:8080")
}

func startGRPCServer() {
	// Create a TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the hotel controller
	hotelController := controllers.NewHotelController(initializers.DB)
	pb.RegisterHotelServiceServer(grpcServer, hotelController)

	log.Println("gRPC server started on :50051")
	
	// Start serving gRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
