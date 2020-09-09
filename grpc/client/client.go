package client

import (
	"context"
	"log"

	pb "github.com/devcharmander/100-day-habits/grpc/pb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

type client struct{}

//Init sets up the client
func Init() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to the server. Error %v", err)
	}
	c := pb.NewTimetableServiceClient(cc)
	trk := make(map[int32]bool, 0)
	trk[1] = false
	trk[2] = true

	req := &pb.HabitRequest{
		Habit: &pb.Habit{
			Name:      "Kiss Anusha everyday",
			Reward:    "Sex",
			Startdate: ptypes.TimestampNow(),
			Track:     trk,
		},
	}
	//TODO should be moved to a new function
	resp, err := c.AddHabit(context.Background(), req)
	if err != nil {
		log.Fatalf("Error adding habit. Error %v", err)

	}
	log.Printf("Status of response: %v", resp.Status)
}
