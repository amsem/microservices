package main

import (
	"context"
	"fmt"

	proto "github.com/amsem/client-RPC/proto"
	micro "go-micro.dev/v4"
)


func main()  {
    service := micro.NewService(micro.Name("encrypter.client"))
    service.Init()
    encrypter := proto.NewEncrypterService("encrypter", service.Client())
     r, err := encrypter.Encrypt(context.TODO(), &proto.Request{
        Message: "Hello from home",
        Key: "111023043350789514532147",
     })

     if err != nil {
        fmt.Println(err)
     }
     fmt.Println("Encrypted Message : ",r.Result)
}
