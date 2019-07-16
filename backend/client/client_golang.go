package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "../src"
)

const (
	address     = "localhost:50051"
)


func main () {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	connection := pb.NewAddeditClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := connection.AddRestaraunt(ctx, &pb.RestarauntInfo {Name: "Saurabh", Rating: "1.0", Cuisines: "South", Address: "Sushant Lok", Time: "6am", Cft: "230 for two person", ID: "1"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.ID)

	r2, err := connection.EditRestaraunt(ctx, &pb.RestarauntInfo {Name: "Saurabh Sabharwal", Rating: "1.0", Cuisines: "South", Address: "Sushant Lok", Time: "6am", Cft: "230 for two person", ID: "1"})

	log.Printf("Greeting: %s", r2.ID)

	r4, err := connection.GetRestaraunt(ctx, &pb.Response {Status: "OK"})

	log.Printf("Greetings: %s", r4.Name)

	r3, err := connection.DeleteRestaraunt(ctx, &pb.RestarauntID {ID: "1"})

	log.Printf("Greeting: %s", r3.Status)
}