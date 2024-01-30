package main

import (
	"context"
	"fmt"
	"io"
	"log"
	wearablepb "test/test/protobuf/wearable"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		log.Println("fail to dial:", err)
	}
	defer conn.Close()

	client := wearablepb.NewWearableServiceClient(conn)
	if err := ui.Init(); err != nil {
		log.Println("Couldn't init UI:", err)
	}
	defer ui.Close()

	lc := widgets.NewPlot()
	lc.Title = "Heartbeat Per Minute"
	lc.SetRect(0, 0, 70, 20)
	lc.Data = make([][]float64, 1)
	lc.DataLabels = []string{"Minute", "Value"}
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorGreen
	lc.Marker = widgets.MarkerDot

	data := make([]float64, 60)
	ui.Render(lc)

	res, err := client.BeatsPerMinute(context.Background(), &wearablepb.BeatsPerMinutesRequest{Id: "tejas"})
	if err != nil {
		log.Println("Couldn't request", err)
	}

	go func() {
		for {
			resp, err := res.Recv()
			if err == io.EOF {
				return
			}

			if err != nil {
				log.Println("Receiving", err)

				return
			}

			data[int(resp.GetMinute())] = float64(resp.GetValue())
			lc.Data[0] = data

			ui.Render(lc)
		}
	}()

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents: // Listen for Keyboard
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-res.Context().Done():
			fmt.Println("All done, possible error", res.Context().Err())
			break
		}
	}

}
