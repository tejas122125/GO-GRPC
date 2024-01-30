package servertypes

import (
	"errors"
	"math/rand"
	wearablepb "test/test/protobuf/wearable"
	"time"

	// "google.golang.org/genproto/googleapis/rpc/code"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/internal/status"
)

type Wearable struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (w *Wearable) BeatsPerMinute(req *wearablepb.BeatsPerMinutesRequest, stream wearablepb.WearableService_BeatsPerMinuteServer) error {

	for {
		select {
		case <-stream.Context().Done():
			return errors.New("stream ended")
		default:
			time.Sleep(1 * time.Second)
			value := 30 + rand.Int31n(80)
			err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})
			if err != nil {
				return errors.New("stream ended")
			}
		}
	}

}

// type WearableServiceServer interface {
// 	BeatsPerMinute(*BeatsPerMinutesRequest, WearableService_BeatsPerMinuteServer) error
// 	mustEmbedUnimplementedWearableServiceServer()
// }
