# Bootcamp Courses Microservice

The **Bootcamp Courses Microservice** handles the management of courses within the bootcamp ecosystem.

## Features

- Create and manage courses
- Retrieve courses based on current login user ID

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mucha-fauzy/bootcamp-course-microservice.git
   ```

2. Navigate to the project directory:

   ```bash
   cd bootcamp-course-microservice
   ```

3. Install the dependencies:

   ```bash
   go mod download
   ```

4. Create Required table (./migrations) and seed(./seeders) the table if necessary


5. Build the project:

   ```bash
   go run main.go
   ```


By default, the microservice will listen on port 8081.

## API Endpoints

- **GET /v1/courses**: Retrieve courses from current user login. Requires a valid JWT token with "teacher" role in the `Authorization` header.

- **POST /v1/courses**: Create a new course. Requires a valid JWT token with "teacher" role in the `Authorization` header.


## Authentication

The microservice uses JWT (JSON Web Tokens) for authentication. To access protected routes, include a valid JWT token in the `Authorization` header as `Bearer <token>`.

## Changelogs
- Add chaining middleware to verify JWT and validate role
- Get user_id from context (before: user_id was obtained path parameter)
- Generate uuid for course ID and change the datatype to string

## Future Updates
- Pisahkan app logic dengan db logic (service: app, repo: db)
- Set secret key untuk JWT di .env dan buat configs nya (users-management-crud-api) -> jadi nanti tinggal dipanggil tidak perlu hard code 

---
Feel free to customize and expand this README.md according to your microservice's specific details, requirements, and usage instructions.