package app

import (
	"advertisement-storage/pkg/db"
	avdertisement_storage "advertisement-storage/pkg/pb"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func Init() error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("database successful started")

	grpcServer := grpc.NewServer()
	server := NewServer()
	avdertisement_storage.RegisterAdvertisementsStorageServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		return err
	}

	log.Printf("advertisements-storage started on %d port", 8000)

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
