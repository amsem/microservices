package main

import (
	"context"

	"github.com/amsem/encryptService/proto"
)




type Encrypter struct{}

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, res *proto.Response) error {
    res.Result = EncryptString(req.Key, req.Message)
    return nil
}

func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, res *proto.Response) error {
    res.Result = DecryptString(req.Key, req.Message)
    return nil
}
