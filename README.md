# API-students
API made in Golang to manage students of a course

Routes: 
-GET /students - List all students
-POST /students - Create student
-GET /students/:id - Get infos from a specific student
-PUT /students/:id - Update student
-DELETE /students/id: - Delete student

Struct Student:
-Name (string)
-CPF (int)
-Email (string)
-Age (int)
-Active (bool)