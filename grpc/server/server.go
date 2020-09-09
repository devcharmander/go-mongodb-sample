package server

import (
	"context"
	"log"
	"net"

	"github.com/devcharmander/100-day-habits/database"

	pb "github.com/devcharmander/100-day-habits/grpc/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTimetableServiceServer
}

func (srv *server) AddHabit(ctx context.Context, in *pb.HabitRequest) (*pb.HabitResponse, error) {
	log.Printf("%s\n%v\n%q", in.Habit.Name, in.Habit.Startdate, in.Habit.Track)

	habit := &database.Habit{
		Name:      in.Habit.Name,
		Track:     in.Habit.Track,
		Reward:    in.Habit.Reward,
		Startdate: in.Habit.Startdate,
	}
	req := &database.MongoRequest{}
	req.Data = make([]*database.Habit, 0)
	req.Data = append(req.Data, habit)
	_, err := database.Client.Create(req)
	if err != nil {
		log.Fatalf("Could not insert record. Error %v", err)
	}
	return &pb.HabitResponse{
		Status: true,
	}, nil

}

// Start - Starts gRPC server for habit tracker
func Start() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not listed to tcp. Error %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimetableServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Error serving gRPC server. Error %v", err)
	}
}
