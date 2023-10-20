package main

import (
	"context"
	"log"
	"time"

	proto "github.com/amsem/asyncServer/proto"
	micro "go-micro.dev/v4"
)

func main() {
	service := micro.NewService(
		micro.Name("weather"),
	)
	p := micro.NewEvent("alerts", service.Client())
	go func() {
		for now := range time.Tick(10 * time.Second) {
			log.Println("Publishing weather alert to Topic: alerts")
			p.Publish(context.TODO(), &proto.Event{
				City:        "Munich",
				Timestamp:   now.UTC().Unix(),
				Temperature: 2,
			})
		}
	}()
    
	service.Init()

    err := service.Run(); 
    if err != nil {
		log.Println(err)
	}
}
