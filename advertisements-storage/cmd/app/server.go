package app

import avdertisement_storage "advertisement-storage/pkg/pb"

type Server struct {
	avdertisement_storage.UnimplementedAdvertisementsStorageServer
}

func NewServer() *Server {
	return &Server{}
}
