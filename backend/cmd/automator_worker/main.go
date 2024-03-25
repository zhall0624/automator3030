package main

import (
	"context"
	"log"
	"net"

	automator_worker "github.com/zhall0624/automator3030/protos"
	"google.golang.org/grpc"
)

type server struct {
	automator_worker.WorkerServer
}

func (s *server) EnqueueJob(ctx context.Context, j *automator_worker.Job) (*automator_worker.Job, error) {
	log.Printf("Received: %v", j.GetName())
	return j, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	s := grpc.NewServer()
	automator_worker.RegisterWorkerServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
