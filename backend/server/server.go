package main

import (
	"context"
	"log"
	"net"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"google.golang.org/grpc"
	pb "../src"
)

const (
	port = ":50051"
)
var db *sql.DB

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) AddRestaraunt(ctx context.Context, in *pb.RestarauntInfo) (*pb.RestarauntID, error) {
	log.Printf("Adding: %v", in.Name)

	var query = "INSERT INTO Restaraunts values ('" + in.Name + "','" + in.Rating + "','" + in.Cuisines + "','" + in.Address + "','" + in.Time + "','" + in.Cft + "','" + in.ID + "')"

	_, err := db.Query (query)

	if err != nil {
		log.Fatal(err)
	}

	return &pb.RestarauntID{ID: "Add " + in.Name}, nil
}

func (s *server) EditRestaraunt(ctx context.Context, in *pb.RestarauntInfo) (*pb.RestarauntID, error) {
	log.Printf("Editing: %v", in.Name)

	var query = "UPDATE Restaraunts SET name = '" + in.Name + "' WHERE ID = '" + in.ID + "'"

	_, err := db.Query (query)

	if err != nil {
		log.Fatal(err)
	}

	return &pb.RestarauntID{ID: "Edit " + in.Name}, nil
}
func (s *server) DeleteRestaraunt(ctx context.Context, in *pb.RestarauntID) (*pb.Response, error) {
	var query = "DELETE FROM Restaraunts WHERE ID = '" + in.ID + "'"

	_, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	return &pb.Response{Status: "Deleted " + in.ID}, nil
}

func (s *server) GetRestaraunt(ctx context.Context, in *pb.Response) (*pb.ListRestaraunt, error) {
	log.Printf("Getting:")
	var List pb.ListRestaraunt
	var query = "SELECT * FROM Restaraunts"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var restaraunt pb.RestarauntInfo
		err := rows.Scan(&restaraunt.Name, &restaraunt.Rating, &restaraunt.Cuisines, &restaraunt.Address, &restaraunt.Time, &restaraunt.Cft, &restaraunt.ID)
		if err != nil {
			log.Fatal(err)
		}
		List.List=append(List.List, &restaraunt)
		log.Println(restaraunt.Name)
	}

	return &List, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err = sql.Open("mysql", "root:zomato@tcp(127.0.0.1)/test")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

	s := grpc.NewServer()
	pb.RegisterAddeditServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}