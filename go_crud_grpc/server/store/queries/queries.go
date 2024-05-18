package my_queries

const InsertStudentQuery = `
INSERT INTO studentTwo (name, studentId, class, email, address)
VALUES (?, ?, ?, ?, ?)
`

const ReadStudentQuery = ` SELECT * FROM studentTwo WHERE studentId = ?`

const UpdateStudentQuery = `UPDATE studentTwo SET name = ?, class = ? , email = ? ,address = ? WHERE studentId = ?`

const DeleteStudentQuery = `DELETE FROM studentTwo WHERE studentId = ?`
