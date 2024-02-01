package servertypes

import (
	"errors"
	"fmt"
	"io"
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

func (w *Wearable) CalculateBeatsPerMinute(stream wearablepb.WearableService_CalculateBeatsPerMinuteServer) error {
	var count, total uint32
	for {
		req, err := stream.Recv()
	
		if err == io.EOF {
			fmt.Println("end of thew file")
			return nil
		}
		if err != nil {
			return err
		}

		total += req.GetValue()
		fmt.Println("recueved", req.GetValue())

		count++
		if count%5 == 0 {
			fmt.Println("total", total, "average", float32(total/5))
			err := stream.Send(&wearablepb.CalculateResponse{Average: float32(total / 5)})
			if err != nil {
				return nil
			}
			total = 0
		}

	
	}
}

// type WearableServiceServer interface {
// 	BeatsPerMinute(*BeatsPerMinutesRequest, WearableService_BeatsPerMinuteServer) error
// 	mustEmbedUnimplementedWearableServiceServer()
// }

// CalculateBeatsPerMinute(WearableService_CalculateBeatsPerMinuteServer) error
