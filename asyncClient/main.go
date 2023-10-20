package main

import (
	"context"
	"log"

	"github.com/amsem/asyncClient/proto"
	micro "go-micro.dev/v4"
)

func ProcessEvent(ctx context.Context, event *proto.Event) error {
    log.Println("Got alert:", &event)
    return nil
}

func main()  {
    service := micro.NewService(micro.Name("weather_client")) 
    service.Init()
    
    err := micro.RegisterSubscriber("alerts", service.Server(), ProcessEvent)
    if err != nil {
        log.Fatal(err)
    }
    err = service.Run()
    if err != nil {
        log.Fatal(err)
    }
}
