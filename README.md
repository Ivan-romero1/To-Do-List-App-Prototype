# To-Do-List-App-Prototype
This app is a simple iteration of a To-Do App that takes care of CRUD operations. It uses 3 different Docker containers, including a container for: 1. The Frontend, 2. The Backend, 3. PostgreSQL database. 

The purpose of this application is to demonstrate communication between a React frontend, a Go REST API, and a PostgreSQL database using Docker containers. 

# To-Do-List-App-Prototype

## Overview

This project is a simple To-Do List application designed to demonstrate the implementation of **CRUD (Create, Read, Update, Delete)** operations using a containerized, full-stack architecture.

The application consists of three separate Docker containers:

1. **Frontend:** A React and TypeScript application that provides the user interface.
2. **Backend:** A Go REST API that handles HTTP requests, application logic, and database interactions.
3. **Database:** A PostgreSQL database that stores and manages To-Do items.

## Project Purpose

The primary purpose of this application is to demonstrate communication between a **React frontend, a Go REST API, and a PostgreSQL database** using Docker containers.

The project also serves as a practical introduction to full-stack development, RESTful API design, database integration, and container orchestration using Docker Compose.

## Technology Stack

- **Frontend:** React, TypeScript, Vite
- **Backend:** Go (Golang), REST API
- **Database:** PostgreSQL
- **Containerization:** Docker and Docker Compose

## Application Architecture

The application follows a three-tier architecture:

```text
React Frontend
      |
      | HTTP Requests
      v
Go REST API
      |
      | SQL Queries
      v
PostgreSQL Database
```

Each component runs in its own Docker container, allowing the services to communicate while remaining independently managed.

## CRUD Operations

The application is designed to support the following operations:

- **Create:** Add new To-Do items.
- **Read:** Retrieve existing To-Do items.
- **Update:** Modify existing To-Do items.
- **Delete:** Remove To-Do items.

