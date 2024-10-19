# Registration System

This project is a Registration System built using React for the frontend and Go with Gin for the backend. It allows users to create, read, update, and delete registration entries.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
  - [Running with Docker Compose](#running-with-docker-compose)
  - [Running Manually](#running-manually)
- [License](#license)

## Features

- User-friendly interface for managing registrations.
- CRUD operations for registration entries.
- Responsive design using Tailwind CSS.
- API built with Go and Gin framework.

## Technologies

- **Frontend**: React, TypeScript, Tailwind CSS
- **Backend**: Go, Gin, GORM
- **Database**: PostgreSQL
- **Containerization**: Docker, Docker Compose

## API Endpoints

### Registration API

- **POST** `/api/registrations`: Create a new registration.
- **GET** `/api/registrations`: Retrieve all registrations.
- **GET** `/api/registrations/:id`: Retrieve a specific registration by ID.
- **PUT** `/api/registrations/:id`: Update a specific registration by ID.
- **DELETE** `/api/registrations/:id`: Delete a specific registration by ID.

## Getting Started

### Running with Docker Compose

1. Ensure you have [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/) installed on your machine.
2. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

3. Create a `.env` file in the root directory and set the following environment variables:

   ```env
   DATABASE_URL=postgres://postgres:postgres@db:5432/registration_db?sslmode=disable
   PORT=8080
   ```

4. Run the following command to start the application:

   ```bash
   docker-compose up --build
   ```

5. Access the frontend at `http://localhost:8000` and the backend API at `http://localhost:3000/api`.

### Running Manually

#### Prerequisites

- Install [Go](https://golang.org/dl/) and [Node.js](https://nodejs.org/en/download/) on your machine.
- Install PostgreSQL and ensure it is running.

#### Backend Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>/backend
   ```

2. Create a PostgreSQL database named `registration_db` and set up the environment variables in a `.env` file:

   ```env
   DATABASE_URL=postgres://postgres:postgres@localhost:5432/registration_db?sslmode=disable
   PORT=8080
   ```

3. Run the database migrations:

   ```bash
   go run cmd/api/main.go
   ```

#### Frontend Setup

1. Open a new terminal and navigate to the frontend directory:

   ```bash
   cd <repository-directory>/frontend
   ```

2. Install the dependencies:

   ```bash
   npm install
   ```

3. Start the frontend application:

   ```bash
   npm start
   ```

4. Access the application at `http://localhost:8000`.

## License

This project is licensed under the MIT License.
