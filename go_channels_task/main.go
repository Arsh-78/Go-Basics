package main

import (
	"fmt"
	"sync"

	"github.com/thedatashed/xlsxreader"
)

type Student struct {
	StudentID string
	Name      string
	Class     string
	Email     string
	Address   string
}

func main() {
	// Create an instance of the reader by opening a target file
	xl, _ := xlsxreader.OpenFile("student_records.xlsx")

	// Ensure the file reader is closed once utilised
	defer xl.Close()

	fmt.Println(len(xl.ReadRows(xl.Sheets[0])))

	var students []Student

	// Iterate on the rows of data
	for row := range xl.ReadRows(xl.Sheets[0]) {
		if len(row.Cells) >= 5 {
			students = append(students, Student{
				StudentID: row.Cells[0].Value,
				Name:      row.Cells[1].Value,
				Class:     row.Cells[2].Value,
				Email:     row.Cells[3].Value,
				Address:   row.Cells[4].Value,
			})
		}
	}

	var wg sync.WaitGroup

	studentChannel := make(chan []Student, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, studentChannel, &wg)
	}

	for i := 1; i < len(students); i += 20 {
		end := i + 20

		if end > len(students) {
			end = len(students)
		}
		studentChannel <- students[i:end]
	}

	close(studentChannel)

	wg.Wait()
}

func worker(id int, studentChannel <-chan []Student, wg *sync.WaitGroup) {

	defer wg.Done()

	for batch := range studentChannel {
		fmt.Printf("Worker %d processing batch with %d records\n", id, len(batch))
		// Process the batch
		for _, student := range batch {
			// Simulate processing each student record
			fmt.Printf("Worker %d processing student ID %s\n", id, student.StudentID)
		}

	}
}
