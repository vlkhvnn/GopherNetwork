# GopherNetwork

GopherNetwork is a distributed network system built with Go. It is designed to provide a robust and scalable platform for network communication and data exchange. The project leverages modern technologies—including Go for the backend and TypeScript, JavaScript, and CSS for the frontend—to offer an integrated solution for networked applications.

## Project Structure

The repository is organized as follows:

- **cmd/**  
  Contains the main application entry points.
- **internal/**  
  Houses the core business logic and internal packages.
- **docs/**  
  Documentation files and guides.
- **web/**  
  Frontend application assets including HTML, CSS, TypeScript, and JavaScript.
- **scripts/**  
  Utility scripts for various development tasks.
- **.air.toml**  
  Configuration for live reloading using [Air](https://github.com/cosmtrek/air).
- **docker-compose.yml**  
  Docker Compose configuration for containerized deployment.
- **Makefile**  
  Build and automation tasks.
- **go.mod & go.sum**  
  Go module files for dependency management.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Air](https://github.com/cosmtrek/air) (optional, for live reloading during development)

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/vlkhvnn/GopherNetwork.git
   cd GopherNetwork
2. **Install Dependencies:**

   Make sure Go is installed, then run:

   ```bash
   go mod download
3. **Running the Application:**

   Ensure Docker is running on your system, then execute:

   ```bash
   docker-compose up
   cd cmd
   cd api
   air
   ```
4. **Migrations:**
   ```bash
   make migrate-up
   ```
5. **UI links:**  
   http://localhost:8080/v1/swagger/index.html  
   http://127.0.0.1:8081/  
Make sure you have set up any necessary environment variables or configuration files as required by the service.

## Configuration

Each component of GopherNetwork may have its own configuration settings. Check the following for details:

- **Backend:**  
  Environment variables and configuration files for service settings.

- **Frontend:**  
  Configuration files located in the `web` directory.

- **Docker:**  
  Environment overrides defined in the `docker-compose.yml` file.

Ensure that any required configuration files or environment variables are set up before running the application.
