## students_api

A practice project to learn api development with golang

## Description

Student management system, only admin can login with username/email and password. Admin can add new student data, edit them, delete them and see the list of students.
It is a simple project to learn the fundamentals backend development.

## Topics I will learn

- CRUD
- Database connection
- Standard golang project structure

## API endpoints

- POST /login (admin)
- POST /student
- GET /students
- GET /students/{id}
- PATCH /students/{id}
- DELETE /students/{id}

## Schema

```go
type Student struct {
    ID          uuid.UUID `json:"id" db:"id"`
    Name        string    `json:"name" db:"name"`
    Email       string    `json:"email" db:"email"`
    Age         int       `json:"age" db:"age"`
    Class       string    `json:"class" db:"class"`
    Roll        int       `json:"roll" db:"roll"`
    BloodGroup  string    `json:"blood_group" db:"blood_group"`
    PhoneNumber string    `json:"phone_number" db:"phone_number"`
}

```
