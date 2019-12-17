package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/cagodoy/tenpo-history-api/database"
	historySvc "github.com/cagodoy/tenpo-history-api/rpc/history"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/cagodoy/tenpo-challenge/lib/proto"
	_ "github.com/lib/pq"

	nats "github.com/nats-io/nats.go"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5040"
		log.Println("missing env variable PORT, using default value...")
	}

	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		postgresDSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
		fmt.Println("missing env variable POSTGRES_DSN, using default value")
	}

	natsHost := os.Getenv("NATS_HOST")
	if natsHost == "" {
		natsHost = "nats"
		log.Println("missing env variable NATS_HOST, using default value...")
	}

	natsPort := os.Getenv("NATS_PORT")
	if natsPort == "" {
		natsPort = "4222"
		log.Println("missing env variable NATS_PORT, using default value...")
	}

	pgSvc, err := database.NewPostgres(postgresDSN)
	if err != nil {
		log.Println("PG DSN ", postgresDSN)
		log.Fatalf("Failed connect to postgres: %v", err)
	}

	srv := grpc.NewServer()

	nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		fmt.Printf("Got disconnected! Reason: %q\n", err)
	})
	nats.ReconnectHandler(func(nc *nats.Conn) {
		fmt.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
	})
	nats.ClosedHandler(func(nc *nats.Conn) {
		fmt.Printf("Connection closed. Reason: %q\n", nc.LastError())
	})
	nc, err := nats.Connect("nats://"+natsHost+":"+natsPort, nats.MaxReconnects(15), nats.ReconnectWait(3*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect to NATS servdr: %v", err)
	}
	log.Println("Connected to NATS server", "nats://"+natsHost+":"+natsPort)

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	svc := historySvc.New(pgSvc, c)

	pb.RegisterHistoryServiceServer(srv, svc)
	reflection.Register(srv)

	log.Println("Starting History service...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to list: %v", err)
	}

	log.Println(fmt.Sprintf("History service, Listening on: %v", port))

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Fatal to serve: %v", err)
	}
}
