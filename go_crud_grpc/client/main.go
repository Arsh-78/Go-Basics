package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example/go_crud_grpc/proto"

	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50052"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect to the server: %v", err)
	}
	defer conn.Close()

	c := pb.NewCRUDClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Menu-based program allowing the user to choose from CRUD
	fmt.Println("\nWelcome to a simple gRPC/MongoDB based app that performs CRUD",
		" operations!")
	fmt.Println("Enter the one of the folliwing choices below:")
	fmt.Print("1 to create an item; 2 to read; 3 to update and 4 to remove: ")

	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		// CreateItem operation
		// Read the name
		fmt.Print("\nEnter the name: ")
		var name string
		fmt.Scanln(&name)

		// Read the ID
		fmt.Print("Enter the ID: ")
		var id string
		fmt.Scanln(&id)

		// Read the category
		fmt.Print("Enter the Class: ")
		var class int32
		fmt.Scanln(&class)

		fmt.Println("Enter the email: ")
		var email string
		fmt.Scanln(&email)

		fmt.Println("Enter the address: ")
		var addr string
		fmt.Scanln(&addr)

		// Populate the Employee struct
		item, err := c.CreateStudent(ctx, &pb.Student{Name: name, StudentId: id, Class: class, Email: email, Address: addr})

		//c.CreateItem(ctx, &pb.Employee{Name: n, Id: i,
		//Category: int32(catInt)})

		if err != nil {
			log.Fatalf("Could not create a new item: %v", err)
		}
		fmt.Println("\nInserted", name, "with the ID", item.Id, "and class", class)

	default:
		fmt.Println("\nWrong option!")
	}

}
