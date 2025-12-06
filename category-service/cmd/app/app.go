package app

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	storageclient "category-service/internal"
	"category-service/pkg/pb"
)

func Init() error {
	storageAddr := os.Getenv("STORAGE_ADDR")

	storageClient, err := storageclient.NewStorageClient(storageAddr)
	if err != nil {
		log.Fatalf("Failed to connect to storage: %s", err.Error())
		return err
	}

	defer storageClient.Close()

	grpcServer := grpc.NewServer()

	categoryServer := NewCategoryServer(storageClient)

	pb.RegisterAdvertisementsStorageServer(grpcServer, categoryServer)

	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Failed listen 8001 port: %s", err)
		return err
	}

	log.Printf("category-service started on %d port", 8001)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")

	grpcServer.GracefulStop()

	log.Println("Server stopped")

	return nil
}
