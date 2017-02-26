package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-file"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.file"),
		micro.Flags(
			cli.StringFlag{
				Name:  "file_dir",
				Usage: "Directory for files served",
				Value: "/tmp",
			},
		),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) {
			file.RegisterHandler(service.Server(), ctx.String("file_dir"))
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
