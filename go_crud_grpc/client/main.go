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
	case "2":
		// ReadItem operation
		fmt.Print("\nEnter the ID: ")
		var id string
		fmt.Scanln(&id)

		read, err := c.ReadStudent(ctx, &pb.ID{Id: id})
		if err != nil {
			log.Fatalf("Error reading the item: %v", err)
		}
		fmt.Println("\nItem found!")
		fmt.Println("Name:", read.Name)
		fmt.Println("ID:", read.StudentId)
		fmt.Println("Class:", read.Class)
		fmt.Println("Email:", read.Email)
		fmt.Println("Address:", read.Address)
	case "3":

		// UpdateItem operation
		// Read the ID
		fmt.Print("\nEnter the existing ID: ")
		var id string
		fmt.Scanln(&id)

		// Read the name
		fmt.Print("Enter the new name: ")
		var name string
		fmt.Scanln(&name)

		fmt.Print("Enter the new Class: ")
		var class int32
		fmt.Scanln(&class)

		fmt.Println("Enter the new email: ")
		var email string
		fmt.Scanln(&email)

		fmt.Println("Enter the new address: ")
		var addr string
		fmt.Scanln(&addr)

		up, err := c.UpdateStudent(ctx, &pb.Student{Name: name, StudentId: id, Class: class, Email: email, Address: addr})
		if err != nil {
			log.Fatalf("Error updating the item: %v", err)
		}
		log.Printf("\nItem updated with the ID: %s", up.Id)
	case "4":
		// DeleteItem operation
		// Ignoring the error - should always be successful regardless of the
		// implicit find result
		// Read the ID
		fmt.Print("\nEnter the existing ID: ")
		var id string
		fmt.Scanln(&id)

		del, _ := c.DeleteStudent(ctx, &pb.ID{Id: id})
		if del != nil {
			log.Printf("\nItem with the ID %s deleted", del.Id)
		} else {
			log.Printf("Successful delete (even though ID didn't exist)")
		}

	default:
		fmt.Println("\nWrong option!")
	}

}
