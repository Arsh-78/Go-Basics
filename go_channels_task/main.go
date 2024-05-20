package main

import (
	"fmt"
	"log"
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
	// Open the Excel file
	xl, err := xlsxreader.OpenFile("student_records.xlsx")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer xl.Close()

	// Read student records from the Excel file
	students := readStudentRecords(xl)
	fmt.Printf("Total students read: %d\n", len(students))

	// Create a channel for student batches
	studentChannel := make(chan []Student, 5)

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, studentChannel, &wg)
	}

	// Distribute student records to workers in batches
	distributeStudents(students, studentChannel)

	// Close the channel and wait for all workers to finish
	close(studentChannel)
	wg.Wait()
}

// readStudentRecords reads student records from the first sheet of the Excel file
func readStudentRecords(xl *xlsxreader.XlsxFileCloser) []Student {
	var students []Student
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
	return students
}

// distributeStudents distributes student records to the channel in batches of 20
func distributeStudents(students []Student, studentChannel chan<- []Student) {
	for i := 1; i < len(students); i += 20 {
		end := i + 20
		if end > len(students) {
			end = len(students)
		}
		studentChannel <- students[i:end]
	}
}

// worker processes batches of student records from the channel
func worker(id int, studentChannel <-chan []Student, wg *sync.WaitGroup) {
	defer wg.Done()
	for batch := range studentChannel {
		fmt.Printf("Worker %d processing batch with %d records\n", id, len(batch))
		for _, student := range batch {
			processStudent(id, student)
		}
	}
}

// processStudent simulates processing a single student record
func processStudent(workerID int, student Student) {
	fmt.Printf("Worker %d processing student ID %s\n", workerID, student.StudentID)
}
