package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"

	p "../proto"
	"google.golang.org/grpc"
)

func main() {
	port := flag.String("p", "127.0.0.1:4040", "port number")
	flag.Parse()
	println(*port)
	conn, err := grpc.Dial(*port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error connecting: ", err)
	}
	defer conn.Close()
	client := p.NewTTSClient(conn)

	body := &p.Text{
		Input: flag.Arg(0),
	}
	response, err := client.Process(context.Background(), body)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("output.wav", response.Output, 0666); err != nil {
		log.Fatal(err)
	}

	// cmd := exec.Command("gsutil", "cp", "output.wav", "gs://randbu")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	// out, err := os.Create("output.wav")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer out.Close()

	// resp, err := http.Get("https://storage.googleapis.com/randbu/output.wav")
	// if err != nil {
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// defer resp.Body.Close()

	// _, err = io.Copy(out, resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
