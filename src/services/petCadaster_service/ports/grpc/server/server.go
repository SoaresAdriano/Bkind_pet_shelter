package grpc

import (
	pb "bkind_pet_shelter/src/services/petCadaster_service/ports/grpc/protos"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":50051"
)

var pets []*pb.PetInfo

type petCadasterServer struct {
	pb.UnimplementedPetCadasterServer
}

func Run() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPetCadasterServer(s, &petCadasterServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Printf("server listening at %v", lis.Addr())
	log.Printf("server listening at %v", lis.Addr())
}

func (s *petCadasterServer) ListPets(input *pb.Empty, stream pb.PetCadaster_ListPetsServer) error {
	fmt.Printf("list request received: %v", input)
	log.Printf("list request received: %v", input)

	for _, pet := range pets {
		if err := stream.Send(pet); err != nil {
			log.Fatalf("failed to send: %v", err)
			return err
		}
	}
	return nil
}

func (s *petCadasterServer) GetPet(ctx context.Context, input *pb.ID) (*pb.PetInfo, error) {
	fmt.Printf("get request received: %v", input)

	log.Printf("get request received: %v", input)

	res := &pb.PetInfo{}

	for _, pet := range pets {
		if pet.GetId().Value == input.GetValue() { //TODO: database get
			res = pet
			break
		}
	}
	return res, nil
}

func (s *petCadasterServer) CreatePetCadaster(ctx context.Context, input *pb.PetInfo) (*pb.ID, error) {
	fmt.Printf("create request received: %v", input)
	log.Printf("create request received: %v", input)

	res := pb.ID{}
	res.Value.Value = uuid.New().String()
	input.Id.Value = res.GetValue()
	pets = append(pets, input) // TODO: database insert
	return &res, nil
}

func (s *petCadasterServer) UpdatePetCadaster(ctx context.Context, input *pb.PetInfo) (*pb.Status, error) {
	fmt.Printf("update request received: %v", input)
	log.Printf("update request received: %v", input)

	res := pb.Status{}

	for index, pet := range pets {
		if pet.GetId() == input.GetId() {
			pets = append(pets[:index], pets[index+1:]...)
			input.Id = pet.GetId()
			pets = append(pets, input)
			res.Value = true //TODO: Get a better response
		}
	}
	return &res, nil
}

// func (s *petCadasterServer) RenamePet(ctx context.Context, input *pb.RenameInfo) (*pb.Status, error) {
// }

func (s *petCadasterServer) DeletePetCadaster(ctx context.Context, input *pb.ID) (*pb.Status, error) {
	fmt.Printf("delete request received: %v", input)
	log.Printf("delete request received: %v", input)

	res := pb.Status{}

	for index, pet := range pets {
		if pet.GetId().Value == input.Value {
			pets = append(pets[:index], pets[index+1:]...)
			res.Value = true
			break
		}
	}
	return &res, nil
}
