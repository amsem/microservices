package main

import (
        "fmt"
       proto "github.com/amsem/encryptService/proto"
        micro "go-micro.dev/v4"
)

func main()  {
    service := micro.NewService(
        micro.Name("encrypter"),
    )
    service.Init()

    proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))
    err := service.Run()
    if err != nil  {
        fmt.Println(err)
    }
}
