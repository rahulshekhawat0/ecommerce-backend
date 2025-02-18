package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/config"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/models"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/routes"
	"github.com/rahulshekhawat0/ecommerce-backend/protos/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	ecommerce.UnimplementedProductServiceServer
}

func (s *server) GetProductDetails(ctx context.Context, req *ecommerce.ProductRequest) (*ecommerce.ProductResponse, error) {
	productID := req.GetProductId()

	// Define a variable to hold the product
	var product models.Product

	// Fetch product from database
	if err := config.DB.First(&product, productID).Error; err != nil {
		log.Println("Error fetching product from DB:", err)
		return nil, status.Errorf(codes.NotFound, "Product not found")
	}

	// Convert database model to proto format
	protoProduct := &ecommerce.Product{
		Id:          strconv.Itoa(int(product.ID)), // Convert uint to string
		Name:        product.Name,
		Price:       float32(product.Price), // Convert float64 to float32
		Description: product.Description,
	}

	// Return the product details wrapped in a ProductResponse
	return &ecommerce.ProductResponse{
		Message: "Product details fetched successfully",
		Product: protoProduct,
	}, nil
}

func main() {
	config.ConnectDatabase()

	// Get underlying sql.DB to properly close it
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()

	// Create a new Fiber instance
	app := fiber.New()

	// Set up routes
	routes.SetupAuthRoutes(app)
	routes.SetupProductRoutes(app)
	routes.SetupCartRoutes(app)
	routes.SetupOrderRoutes(app)
	routes.SetupAdminRoutes(app)

	// Define a basic route
	app.Get("/ecom", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to E-commerce API!")
	})

	// Get the port value from the environment or default to 8000
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Error in getting PORT value") // Default to port 8000 if PORT is not set
		port = "8000"
	}

	// Start Fiber server in a goroutine so it runs concurrently
	go func() {
		log.Fatal(app.Listen(":" + port)) // Start the Fiber web server
	}()

	// Set up the gRPC server
	lis, err := net.Listen("tcp", ":50051") // Listen on TCP port 50051
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()                                // Create a new gRPC server
	ecommerce.RegisterProductServiceServer(grpcServer, &server{}) // Register our server with gRPC

	fmt.Println("gRPC Server is running on port 50051...")

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
