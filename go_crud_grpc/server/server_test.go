package main

import (
	"context"
	"testing"

	pb "example/go_crud_grpc/proto"
	ser "example/go_crud_grpc/service"

	"github.com/DATA-DOG/go-sqlmock"
)

type TestServer struct {
	server *server
	mock   sqlmock.Sqlmock
}

func newTestServer(t *testing.T) *TestServer {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	service := ser.NewService(db)
	return &TestServer{
		server: &server{service: service},
		mock:   mock,
	}
}

func TestCreateStudent(t *testing.T) {
	ts := newTestServer(t)
	defer ts.mock.ExpectClose()

	ts.mock.ExpectExec("INSERT INTO studentTwo").
		WithArgs("John Doe", "1", 10, "john.doe@example.com", "123 Main St").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.Student{
		StudentId: "1",
		Name:      "John Doe",
		Class:     10,
		Email:     "john.doe@example.com",
		Address:   "123 Main St",
	}

	_, err := ts.server.CreateStudent(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateStudent failed: %v", err)
	}
}

func TestReadStudent(t *testing.T) {
	ts := newTestServer(t)
	defer ts.mock.ExpectClose()

	rows := sqlmock.NewRows([]string{"name", "studentId", "class", "email", "address"}).
		AddRow("John Doe", "1", 10, "john.doe@example.com", "123 Main St")

	ts.mock.ExpectQuery("SELECT \\* FROM studentTwo WHERE studentId = ?").
		WithArgs("1").
		WillReturnRows(rows)

	req := &pb.ID{Id: "1"}

	resp, err := ts.server.ReadStudent(context.Background(), req)
	if err != nil {
		t.Fatalf("ReadStudent failed: %v", err)
	}

	if resp.StudentId != "1" || resp.Name != "John Doe" || resp.Class != 10 || resp.Email != "john.doe@example.com" || resp.Address != "123 Main St" {
		t.Fatalf("ReadStudent returned unexpected result: %v", resp)
	}
}
func TestUpdateStudent(t *testing.T) {
	ts := newTestServer(t)
	defer ts.mock.ExpectClose()

	ts.mock.ExpectExec(`UPDATE studentTwo SET name = \?, class = \?, email = \?, address = \? WHERE studentId = \?`).
		WithArgs("John Doe", 11, "john.doe@newexample.com", "456 New St", "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.Student{
		StudentId: "1",
		Name:      "John Doe",
		Class:     11,
		Email:     "john.doe@newexample.com",
		Address:   "456 New St",
	}

	_, err := ts.server.UpdateStudent(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateStudent failed: %v", err)
	}

	if err := ts.mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteStudent(t *testing.T) {
	ts := newTestServer(t)
	defer ts.mock.ExpectClose()

	ts.mock.ExpectExec("DELETE FROM studentTwo WHERE studentId = ?").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.ID{Id: "1"}

	_, err := ts.server.DeleteStudent(context.Background(), req)
	if err != nil {
		t.Fatalf("DeleteStudent failed: %v", err)
	}
}
