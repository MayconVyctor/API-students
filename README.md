# API-students
Students API
Description:
This project was developed during the Golang from Zero course. It is a RESTful API built to manage student records, providing basic CRUD operations like creating, retrieving, updating, and deleting students.

Technologies Used:

Go (Golang) — main programming language

Echo — lightweight and fast web framework for building APIs

SQLite — lightweight, embedded database for data persistence

Git — version control system

✅ Features:

Create new student records

Retrieve all students or by ID

Update student information

Delete student records

 Project Goal:
The goal of this project is to put into practice the concepts learned in the course, focusing on Go and the fundamentals of RESTful API development. The code is structured using clean practices to ensure readability, maintainability, and ease of learning.



Routes: 
-GET /students - List all students
-POST /students - Create student
-GET /students/:id - Get infos from a specific student
-PUT /students/:id - Update student
-DELETE /students/id: - Delete student
-GET /students?active=<true/false> - List all active/non-active students

Struct Student:
-Name (string)
-CPF (int)
-Email (string)
-Age (int)
-Active (bool)
