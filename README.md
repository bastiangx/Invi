# Invi 🎣

## Table of Contents
- [Project Overview](#project-overview)
- [Features](#features)
- [Technologies](#technologies)
  - [Frontend](#frontend)
  - [Backend](#backend)
  - [Database](#database)
  - [Deployment](#deployment)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
  - [Prerequisites](#prerequisites)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Deployment to Render](#deployment-to-render)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Environment Variables](#environment-variables)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)

## Project Overview

**Invi``** is a lightweight web application designed to manage RSVPs for parties and events efficiently. Built with a Go backend using the Gin-Gonic framework, it provides a straightforward interface for users to submit their attendance and for organizers to track responses. The project leverages SQLite3 for data storage and is deployed on Render for seamless scalability and management.

## Features

- **Simple RSVP Form:** Users can submit their name and email to RSVP for events.
- **RSVP Count Display:** Shows the total number of RSVPs for each event.
- **Event Management:** Display event details (name, title, time, location) with the ability to add new events easily.
- **Unique RSVP IDs:** Each RSVP is assigned a unique ID for easy enumeration and deletion.
- **Dynamic Event Parsing:** Automatically parses event details from a database file and updates the frontend accordingly.

## Technologies

### Frontend

- **HTML5 & CSS3:** Structure and styling of the web pages.
- **JavaScript:** Handles form submissions and dynamic content updates.
- **Responsive Design:** Ensures the application is accessible on various devices.

### Backend

- **Go (Golang):** Primary language for the backend server.
- **Gin-Gonic:** Lightweight and high-performance HTTP web framework for Go.
- **Netlify Functions / Render Server:** Handles serverless functions or server deployment.

### Database

- **SQLite3:** Lightweight, file-based relational database for storing RSVPs and event details.

### Deployment

- **Render:** Platform for deploying the Go backend, managing server instances, and handling continuous deployments.

## Project Structure

```
Invi/
├── backend/
│   ├── db/
│   │   └── rsvp.db
│   ├── models/
│   │   └── rsvp.go
│   ├── handlers/
│   │   └── handlers.go
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── index.html
│   ├── styles.css
│   └── scripts.js
├── README.md
└── .gitignore
```

- **backend/**: Contains all backend-related code.
  - **db/**: Stores the SQLite3 database file.
  - **models/**: Defines data models and database interactions.
  - **handlers/**: Contains HTTP handler functions for API endpoints.
  - **main.go**: Entry point for the Go application.
  - **go.mod & go.sum**: Go module files managing dependencies.
- **frontend/**: Contains frontend assets.
  - **index.html**: Main HTML file.
  - **styles.css**: CSS styling.
  - **scripts.js**: JavaScript for frontend interactions.
- **README.md**: Project documentation.
- **.gitignore**: Specifies intentionally untracked files to ignore.

## Setup and Installation

### Prerequisites

- **Go:** Ensure Go is installed. [Download Go](https://golang.org/dl/)
- **Git:** Version control system. [Download Git](https://git-scm.com/downloads)
- **Render Account:** For deployment. [Sign Up on Render](https://render.com/)
- **SQLite3:** For managing the database. [Install SQLite](https://www.sqlite.org/download.html)

### Backend Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/Invi.git
   cd Invi/backend
   ```

2. **Initialize Go Modules:**
   Ensure you're in the `backend/` directory.
   ```bash
   go mod init github.com/yourusername/Invi/backend
   go mod tidy
   ```

3. **Install Dependencies:**
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/mattn/go-sqlite3
   go get -u github.com/gin-contrib/cors
   ```

4. **Database Setup:**
   - Navigate to the `db/` directory:
     ```bash
     cd db
     ```
   - Initialize the SQLite database and create the `rsvp` table:
     ```bash
     sqlite3 rsvp.db
     ```
     Inside the SQLite shell, run:
     ```sql
     CREATE TABLE IF NOT EXISTS rsvp (
         id INTEGER PRIMARY KEY AUTOINCREMENT,
         name TEXT NOT NULL,
         email TEXT NOT NULL,
         event TEXT NOT NULL
     );
     .exit
     ```
   - **Note:** Ensure `rsvp.db` is located at `backend/db/rsvp.db`.

5. **Configure Environment Variables:**
   - Create a `.env` file in the `backend/` directory (optional):
     ```bash
     PORT=8080
     DATABASE_URL=./db/rsvp.db
     ```
   - Alternatively, set environment variables directly in your deployment platform.

6. **Run the Backend Locally:**
   ```bash
   go run main.go
   ```
   - The server should start on `http://localhost:8080`.

### Frontend Setup

1. **Navigate to Frontend Directory:**
   ```bash
   cd ../frontend
   ```

2. **Serve the Frontend:**
   - You can use a simple HTTP server. For example, using Python:
     ```bash
     python3 -m http.server 8000
     ```
   - Access the frontend at `http://localhost:8000`.

## Deployment to Render

**Render** offers a straightforward deployment process for Go applications. Follow these steps to deploy your Invi backend.

### Step 1: Push Your Code to GitHub

Ensure all your changes are committed and pushed to your GitHub repository.

```bash
cd Invi
git add .
git commit -m "Prepare for Render deployment"
git push origin main
```

### Step 2: Create a New Web Service on Render

1. **Log in to Render:**
   - Go to [Render Dashboard](https://dashboard.render.com/) and log in.

2. **Create a New Web Service:**
   - Click on the **"New"** button and select **"Web Service"**.

3. **Connect to Your Repository:**
   - Choose your GitHub account and select the `Invi` repository.

4. **Configure the Web Service:**
   - **Name:** `bigfish-backend`
   - **Region:** Select a region closest to your users.
   - **Branch:** `main` (or your default branch)
   - **Root Directory:** `backend`
   - **Environment:** `Go`
   - **Build Command:** `go build -o main .`
   - **Start Command:** `./main`

5. **Set Environment Variables:**
   - In the **Environment** section, add any necessary variables:
     - `PORT` (optional, Render typically handles this)
     - `DATABASE_URL=./db/rsvp.db` (if using SQLite; consider external DB for persistence)

6. **Choose Plan:**
   - For a small-scale project, the **Free** plan should suffice.

7. **Create Web Service:**
   - Click **"Create Web Service"** to initiate deployment.

### Step 3: Configure the Frontend to Point to Render Backend

Update your `frontend/scripts.js` to use the Render-deployed backend URL.

```javascript
// frontend/scripts.js

const API_BASE_URL = 'https://bigfish-backend.onrender.com'; // Replace with your actual Render URL

async function submitRSVP(event) {
    event.preventDefault();

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const eventName = document.getElementById('event').value; // Assuming you have an event selector

    const response = await fetch(`${API_BASE_URL}/add_rsvp`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, email, event: eventName })
    });

    if (response.ok) {
        alert('RSVP submitted successfully!');
        // Optionally, refresh the RSVP count
    } else {
        alert('Failed to submit RSVP.');
    }
}

document.getElementById('rsvp-form').addEventListener('submit', submitRSVP);
```

### Step 4: Enable CORS in Backend

If your frontend is hosted on a different domain or port, ensure CORS is properly configured.

1. **Install CORS Middleware:**
   ```bash
   go get github.com/gin-contrib/cors
   ```

2. **Update `main.go`:**

   ```go
   // backend/main.go
   package main

   import (
       "log"
       "os"

       "github.com/gin-contrib/cors"
       "github.com/gin-gonic/gin"
       "github.com/yourusername/Invi/backend/models"
   )

   func main() {
       router := gin.Default()

       // Enable CORS
       router.Use(cors.Default())

       // Initialize the database
       db := models.GetDB()
       defer db.Close()

       // Create the RSVP table if it doesn't exist
       createTableQuery := `
       CREATE TABLE IF NOT EXISTS rsvp (
           id INTEGER PRIMARY KEY AUTOINCREMENT,
           name TEXT NOT NULL,
           email TEXT NOT NULL,
           event TEXT NOT NULL
       );
       `
       _, err := db.Exec(createTableQuery)
       if err != nil {
           log.Fatalf("Failed to create table: %v", err)
       }

       // Define API routes
       router.GET("/", func(c *gin.Context) {
           c.JSON(200, gin.H{
               "message": "Welcome to Invi API",
           })
       })

       router.POST("/add_rsvp", AddRSVPHandler)
       router.GET("/get_rsvps", GetRSVPsHandler)
       router.DELETE("/delete_rsvp/:id", DeleteRSVPHandler)

       // Get the PORT from environment variables, default to 8080 if not set
       port := os.Getenv("PORT")
       if port == "" {
           port = "8080"
       }

       // Start the server on the specified port
       if err := router.Run(":" + port); err != nil {
           log.Fatalf("Failed to run server: %v", err)
       }
   }
   ```

### Step 5: Deploy and Monitor

1. **Automatic Deployment:**
   - Render will automatically build and deploy your application whenever you push changes to the connected branch.

2. **Monitor Logs:**
   - Use the Render dashboard to view build and runtime logs for debugging and monitoring.

3. **Access Your Deployed Backend:**
   - Once deployed, Render provides a live URL (e.g., `https://bigfish-backend.onrender.com`).

## API Endpoints

### 1. Add RSVP

- **URL:** `/add_rsvp`
- **Method:** `POST`
- **Description:** Submits a new RSVP.
- **Request Body:**
  ```json
  {
      "name": "John Doe",
      "email": "john.doe@example.com",
      "event": "Birthday Party"
  }
  ```
- **Response:**
  - **Success (200):**
    ```json
    {
        "message": "RSVP added successfully"
    }
    ```
  - **Error (400/500):**
    ```json
    {
        "error": "Error message"
    }
    ```

### 2. Get RSVPs

- **URL:** `/get_rsvps`
- **Method:** `GET`
- **Description:** Retrieves all RSVPs or RSVPs for a specific event.
- **Query Parameters:**
  - `event` (optional): Filter RSVPs by event name.
- **Response:**
  - **Success (200):**
    ```json
    [
        {
            "id": 1,
            "name": "John Doe",
            "email": "john.doe@example.com",
            "event": "Birthday Party"
        },
        ...
    ]
    ```
  - **Error (500):**
    ```json
    {
        "error": "Error message"
    }
    ```

### 3. Delete RSVP

- **URL:** `/delete_rsvp/:id`
- **Method:** `DELETE`
- **Description:** Deletes an RSVP by its unique ID.
- **Path Parameters:**
  - `id`: ID of the RSVP to delete.
- **Response:**
  - **Success (200):**
    ```json
    {
        "message": "RSVP with ID 1 deleted"
    }
    ```
  - **Error (400/404/500):**
    ```json
    {
        "error": "Error message"
    }
    ```

## Database Schema

The application uses SQLite3 with a single table named `rsvp`.

```sql
CREATE TABLE IF NOT EXISTS rsvp (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    event TEXT NOT NULL
);
```

### Columns

- `id`: Unique identifier for each RSVP (Primary Key).
- `name`: Name of the person RSVPing.
- `email`: Email address of the person RSVPing.
- `event`: Name of the event the person is RSVPing to.

## Environment Variables

Ensure the following environment variables are set, especially when deploying to Render:

- `PORT`: The port number the server listens on (handled by Render).
- `DATABASE_URL`: Path to the SQLite3 database file (if using SQLite).

### Example `.env` File

```env
PORT=8080
DATABASE_URL=./db/rsvp.db
```

**Note:** For security and flexibility, avoid hardcoding sensitive information. Use Render's environment variable settings to manage these variables securely.

## Troubleshooting

### Common Issues

1. **Server Not Starting:**
   - **Check Logs:** Inspect Render's build and runtime logs for error messages.
   - **PORT Configuration:** Ensure the server listens on the `PORT` environment variable.

2. **Database Connection Errors:**
   - **Path Issues:** Verify the database path in `DATABASE_URL` is correct.
   - **Permissions:** Ensure the application has read/write permissions to the database file.

3. **CORS Errors:**
   - **Middleware Configuration:** Ensure CORS is properly configured in the backend.
   - **Frontend Requests:** Verify that the frontend is making requests to the correct backend URL.

4. **Deployment Failures:**
   - **Dependency Issues:** Run `go mod tidy` to ensure all dependencies are included.
   - **Build Command Errors:** Ensure the build command (`go build -o main .`) executes without errors locally.

### Debugging Tips

- **Local Testing:** Always test your application locally before deploying.
- **API Testing Tools:** Use tools like **Postman** or **cURL** to test API endpoints independently.
- **Verbose Logging:** Implement detailed logging in your application to trace issues.

## Contributing

Contributions are welcome! Please follow these steps to contribute to the Invi project:

1. **Fork the Repository:**
   - Click the "Fork" button on the repository page.

2. **Clone Your Fork:**
   ```bash
   git clone https://github.com/yourusername/Invi.git
   cd Invi
   ```

3. **Create a New Branch:**
   ```bash
   git checkout -b feature/YourFeatureName
   ```

4. **Make Changes and Commit:**
   ```bash
   git add .
   git commit -m "Add Your Feature"
   ```

5. **Push to Your Fork:**
   ```bash
   git push origin feature/YourFeatureName
   ```

6. **Create a Pull Request:**
   - Navigate to your fork on GitHub and click "Compare & pull request".

7. **Review and Merge:**
   - Project maintainers will review your PR and merge if everything looks good.

**Note:** Ensure your code follows the project's coding standards and passes all tests before submitting a PR.

---
