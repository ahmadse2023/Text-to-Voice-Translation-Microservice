package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os/exec"

	p "../proto"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	p.RegisterTTSServer(srv, server{})
	err = srv.Serve(listener)

}

func (server) Process(ctx context.Context, text *p.Text) (*p.Speech, error) {

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("trans", "-b", "-p", "-download-audio-as", tempFile.Name(), text.Input)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	return &p.Speech{Output: data}, nil
}
