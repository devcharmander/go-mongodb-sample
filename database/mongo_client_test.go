package database

import (
	"log"
	"testing"

	"github.com/golang/protobuf/ptypes"
)

func TestCreate(t *testing.T) {
	trk := make(map[int32]bool, 0)
	trk[1] = false
	trk[2] = true

	habit := &Habit{
		Name:      "test",
		Track:     trk,
		Reward:    "PS4",
		Startdate: ptypes.TimestampNow(),
	}

	req := &MongoRequest{
		Filter: nil,
	}
	req.Data = make([]*Habit, 0)
	req.Data = append(req.Data, habit)
	_, err := Client.Create(req)
	if err != nil {
		log.Fatalf("Could not insert record. Error %v", err)
	}
}
