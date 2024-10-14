1. **Project Structure Setup**
2. **Setting Up the Go Backend with Gin-Gonic**
3. **Integrating SQLite3 with Go**
4. **Creating API Endpoints**
5. **Setting Up Netlify Functions for Serverless Deployment**
6. **Connecting Frontend with Backend**
7. **Deploying to Netlify**
8. **Testing Your Setup**

Let's dive in!

---

## 1. Project Structure Setup

Organizing your project structure is crucial for maintainability and scalability. Here's a recommended structure for your **BigFish** project:

```
BigFish/
├── backend/
│   ├── functions/
│   │   ├── add_rsvp/
│   │   │   └── main.go
│   │   ├── get_rsvps/
│   │   │   └── main.go
│   │   └── delete_rsvp/
│   │       └── main.go
│   ├── db/
│   │   └── rsvp.db
│   ├── models/
│   │   └── rsvp.go
│   ├── main.go
│   └── go.mod
├── frontend/
│   ├── index.html
│   ├── styles.css
│   └── scripts.js
├── README.md
└── netlify.toml
```

- **backend/**: Contains all backend-related code.
  - **functions/**: Contains Netlify serverless functions.
  - **db/**: Stores the SQLite3 database file.
  - **models/**: Contains data models.
  - **main.go**: Entry point if you decide to run a local server.
- **frontend/**: Contains frontend assets.
- **netlify.toml**: Configuration file for Netlify.

---

## 2. Setting Up the Go Backend with Gin-Gonic

### a. Initialize the Go Module

Navigate to your `backend/` directory and initialize a Go module:

```bash
cd BigFish/backend
go mod init github.com/yourusername/BigFish/backend
```

Replace `github.com/yourusername/BigFish/backend` with your actual module path.

### b. Install Dependencies

Install Gin-Gonic and SQLite3 driver:

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/mattn/go-sqlite3
```

### c. Create the `main.go` File

This file is optional if you're primarily using serverless functions. However, it's useful for local development and testing.

```go
// backend/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to BigFish API",
        })
    })

    // Define your routes here

    router.Run(":8080") // Run on port 8080
}
```

You can run this local server using:

```bash
go run main.go
```

Access `http://localhost:8080` to see the welcome message.

---

## 3. Integrating SQLite3 with Go

### a. Create the Database and Table

Ensure you have SQLite3 installed. Navigate to the `backend/db/` directory and create the `rsvp.db` database.

```bash
cd backend/db
sqlite3 rsvp.db
```

Inside the SQLite shell, create the `rsvp` table:

```sql
CREATE TABLE rsvp (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    event TEXT NOT NULL
);
.exit
```

### b. Define the RSVP Model

Create a model to interact with the `rsvp` table.

```go
// backend/models/rsvp.go
package models

import (
    "database/sql"
    "log"
)

type RSVP struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Event string `json:"event"`
}

func GetDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./db/rsvp.db")
    if err != nil {
        log.Fatal(err)
    }
    return db
}
```

---

## 4. Creating API Endpoints

We'll create three main serverless functions:

1. **Add RSVP**: To submit a new RSVP.
2. **Get RSVPs**: To fetch all RSVPs for an event.
3. **Delete RSVP**: To delete an RSVP by ID.

### a. Add RSVP Function

```go
// backend/functions/add_rsvp/main.go
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    _ "github.com/mattn/go-sqlite3"
    "net/http"
)

type RSVP struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Event string `json:"event"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var rsvp RSVP
    err := json.Unmarshal([]byte(request.Body), &rsvp)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
    }

    db, err := sql.Open("sqlite3", "db/rsvp.db")
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
    }
    defer db.Close()

    stmt, err := db.Prepare("INSERT INTO rsvp(name, email, event) VALUES (?, ?, ?)")
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
    }
    defer stmt.Close()

    _, err = stmt.Exec(rsvp.Name, rsvp.Email, rsvp.Event)
    if err != nil {
        return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
    }

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       fmt.Sprintf("RSVP added for %s", rsvp.Name),
    }, nil
}

func main() {
    lambda.Start(handler)
}
```

**Note**: Netlify uses its own way to handle serverless functions. To simplify, we'll adjust the approach to fit Netlify's serverless functions using Go.

### b. Adjusting for Netlify Functions

Netlify expects serverless functions to follow a specific structure. We'll use the `netlify-functions-go` package to adapt Go functions for Netlify.

First, install the necessary package:

```bash
go get github.com/netlify/functions
```

However, as of my knowledge cutoff in September 2021, Netlify primarily supports Node.js, but they have experimental support for Go. Alternatively, you can use a build step to compile your Go functions to executables.

For simplicity, we'll use HTTP-triggered functions compatible with Netlify's serverless functions.

### c. Implementing the Add RSVP Function for Netlify

Create the `add_rsvp` function:

```go
// backend/functions/add_rsvp/main.go
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    _ "github.com/mattn/go-sqlite3"
)

type RSVP struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Event string `json:"event"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var rsvp RSVP
    err := json.Unmarshal([]byte(request.Body), &rsvp)
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusBadRequest,
            Body:       "Invalid request body",
        }, nil
    }

    db, err := sql.Open("sqlite3", "../db/rsvp.db")
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Database connection failed",
        }, nil
    }
    defer db.Close()

    stmt, err := db.Prepare("INSERT INTO rsvp(name, email, event) VALUES (?, ?, ?)")
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Failed to prepare statement",
        }, nil
    }
    defer stmt.Close()

    _, err = stmt.Exec(rsvp.Name, rsvp.Email, rsvp.Event)
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Failed to execute statement",
        }, nil
    }

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       fmt.Sprintf("RSVP added for %s", rsvp.Name),
    }, nil
}

func main() {
    lambda.Start(handler)
}
```

**Note**: Ensure that your functions can access the `rsvp.db` file. Depending on Netlify's deployment environment, you might need to adjust paths or consider using an external database.

### d. Create Get RSVPs Function

Similarly, create a function to retrieve RSVPs:

```go
// backend/functions/get_rsvps/main.go
package main

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    _ "github.com/mattn/go-sqlite3"
)

type RSVP struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Event string `json:"event"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    eventName := request.QueryStringParameters["event"]

    db, err := sql.Open("sqlite3", "../db/rsvp.db")
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Database connection failed",
        }, nil
    }
    defer db.Close()

    var rows *sql.Rows
    if eventName != "" {
        rows, err = db.Query("SELECT id, name, email, event FROM rsvp WHERE event = ?", eventName)
    } else {
        rows, err = db.Query("SELECT id, name, email, event FROM rsvp")
    }

    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Failed to query database",
        }, nil
    }
    defer rows.Close()

    var rsvps []RSVP
    for rows.Next() {
        var rsvp RSVP
        err := rows.Scan(&rsvp.ID, &rsvp.Name, &rsvp.Email, &rsvp.Event)
        if err != nil {
            continue
        }
        rsvps = append(rsvps, rsvp)
    }

    response, err := json.Marshal(rsvps)
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Failed to marshal response",
        }, nil
    }

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       string(response),
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
    }, nil
}

func main() {
    lambda.Start(handler)
}
```

### e. Create Delete RSVP Function

Finally, create a function to delete an RSVP by ID:

```go
// backend/functions/delete_rsvp/main.go
package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "strconv"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    _ "github.com/mattn/go-sqlite3"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    idStr := request.PathParameters["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusBadRequest,
            Body:       "Invalid ID",
        }, nil
    }

    db, err := sql.Open("sqlite3", "../db/rsvp.db")
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Database connection failed",
        }, nil
    }
    defer db.Close()

    stmt, err := db.Prepare("DELETE FROM rsvp WHERE id = ?")
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Failed to prepare statement",
        }, nil
    }
    defer stmt.Close()

    res, err := stmt.Exec(id)
    if err != nil {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusInternalServerError,
            Body:       "Failed to execute statement",
        }, nil
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil || rowsAffected == 0 {
        return events.APIGatewayProxyResponse{
            StatusCode: http.StatusNotFound,
            Body:       "RSVP not found",
        }, nil
    }

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       fmt.Sprintf("RSVP with ID %d deleted", id),
    }, nil
}

func main() {
    lambda.Start(handler)
}
```

---

## 5. Setting Up Netlify Functions for Serverless Deployment

Netlify supports serverless functions, but integrating Go requires specific configurations. Here's how to set it up:

### a. Install Netlify CLI

First, install the Netlify CLI globally if you haven't already:

```bash
npm install -g netlify-cli
```

### b. Configure `netlify.toml`

Create a `netlify.toml` file in the root of your project (`BigFish/`):

```toml
[build]
  functions = "backend/functions"
  publish = "frontend"

[[redirects]]
  from = "/api/*"
  to = "/.netlify/functions/:splat"
  status = 200
```

This configuration tells Netlify where to find your serverless functions and how to handle API routes.

### c. Building Go Functions for Netlify

Netlify expects functions to be compiled binaries. Here's how to prepare your Go functions:

1. **Set Up Each Function for Compilation**

   For each function (`add_rsvp`, `get_rsvps`, `delete_rsvp`), ensure the `main.go` is properly set up to handle HTTP requests as Netlify expects.

   Here's an example for `add_rsvp`:

   ```go
   // backend/functions/add_rsvp/main.go
   package main

   import (
       "database/sql"
       "encoding/json"
       "fmt"
       "net/http"
       "os"

       "github.com/aws/aws-lambda-go/events"
       "github.com/aws/aws-lambda-go/lambda"
       _ "github.com/mattn/go-sqlite3"
   )

   type RSVP struct {
       Name  string `json:"name"`
       Email string `json:"email"`
       Event string `json:"event"`
   }

   func handler(w http.ResponseWriter, r *http.Request) {
       var rsvp RSVP
       err := json.NewDecoder(r.Body).Decode(&rsvp)
       if err != nil {
           http.Error(w, "Invalid request body", http.StatusBadRequest)
           return
       }

       dbPath := os.Getenv("LAMBDA_TASK_ROOT") + "/../db/rsvp.db"
       db, err := sql.Open("sqlite3", dbPath)
       if err != nil {
           http.Error(w, "Database connection failed", http.StatusInternalServerError)
           return
       }
       defer db.Close()

       stmt, err := db.Prepare("INSERT INTO rsvp(name, email, event) VALUES (?, ?, ?)")
       if err != nil {
           http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
           return
       }
       defer stmt.Close()

       _, err = stmt.Exec(rsvp.Name, rsvp.Email, rsvp.Event)
       if err != nil {
           http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
           return
       }

       fmt.Fprintf(w, "RSVP added for %s", rsvp.Name)
   }

   func main() {
       http.HandleFunc("/", handler)
       port := os.Getenv("PORT")
       if port == "" {
           port = "8080"
       }
       http.ListenAndServe(":"+port, nil)
   }
   ```

2. **Build the Functions**

   For each function, navigate to its directory and build the binary:

   ```bash
   cd backend/functions/add_rsvp
   GOOS=linux GOARCH=amd64 go build -o add_rsvp
   ```

   Repeat this for `get_rsvps` and `delete_rsvp`.

3. **Place Binaries in Functions Directory**

   Ensure that each function directory contains the compiled binary (without the `.exe` extension) and any necessary files like `rsvp.db`. However, Netlify's serverless functions have limited support for writing to the filesystem. Consider using an external database service like **Firebase**, **AWS RDS**, or **Supabase** for production. For development, you can use SQLite.

4. **Set Executable Permissions**

   Ensure that the binaries are executable:

   ```bash
   chmod +x add_rsvp
   ```

### d. Deploying Functions

With the functions built and `netlify.toml` configured, deploy to Netlify:

```bash
netlify deploy
```

Follow the prompts to link your project to a Netlify site. Choose the appropriate build and publish directories as per your `netlify.toml`.

---

## 6. Connecting Frontend with Backend

Your frontend will interact with the backend through the API endpoints you've set up. Here's how to connect them:

### a. Submitting an RSVP

In your `scripts.js`, add a function to handle form submissions:

```javascript
// frontend/scripts.js

async function submitRSVP(event) {
    event.preventDefault();

    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const eventName = document.getElementById('event').value; // Assuming you have an event selector

    const response = await fetch('/api/add_rsvp', {
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

Ensure your HTML form has the correct IDs:

```html
<!-- frontend/index.html -->
<form id="rsvp-form">
    <input type="text" id="name" name="name" placeholder="Your Name" required />
    <input type="email" id="email" name="email" placeholder="Your Email" required />
    <select id="event" name="event">
        <!-- Populate with events -->
    </select>
    <button type="submit">Submit</button>
</form>
```

### b. Fetching RSVPs

To display the number of RSVPs:

```javascript
// frontend/scripts.js

async function fetchRSVPs(eventName = '') {
    let url = '/api/get_rsvps';
    if (eventName) {
        url += `?event=${encodeURIComponent(eventName)}`;
    }

    const response = await fetch(url);
    if (response.ok) {
        const rsvps = await response.json();
        document.getElementById('rsvp-count').innerText = rsvps.length;
    } else {
        console.error('Failed to fetch RSVPs.');
    }
}

// Call this function on page load and when events change
document.addEventListener('DOMContentLoaded', () => {
    fetchRSVPs();
});
```

Ensure you have an element to display the count:

```html
<!-- frontend/index.html -->
<p>RSVP'd: <span id="rsvp-count">0</span></p>
```

---

## 7. Deploying to Netlify

With your backend functions and frontend ready, deploy your project to Netlify.

### a. Initial Deployment

From the root of your project (`BigFish/`), run:

```bash
netlify deploy --prod
```

Follow the prompts to deploy your site.

### b. Environment Variables

If your functions require environment variables (e.g., database paths), set them in Netlify:

1. Go to your site's dashboard on Netlify.
2. Navigate to **Site settings** > **Build & deploy** > **Environment**.
3. Add the necessary variables.

**Note**: For SQLite3, it's best suited for development. For production, consider using a managed database service to ensure data persistence and scalability.

---

## 8. Testing Your Setup

After deployment, test your application to ensure everything works as expected.

1. **Submit an RSVP**: Fill out the form and submit. Check if the RSVP count updates.
2. **View RSVPs**: Ensure that the RSVPs are correctly fetched and displayed.
3. **Delete an RSVP**: If you have a frontend option to delete RSVPs, test this functionality.

---

## Additional Recommendations

1. **Use an External Database for Production**: SQLite is great for development but not recommended for production serverless environments. Consider using **PostgreSQL**, **MySQL**, **Firebase Firestore**, or **Supabase**.

2. **Implement Authentication**: To secure your API endpoints, especially for actions like deleting RSVPs.

3. **Input Validation and Sanitization**: Ensure all inputs are validated to prevent SQL injection and other security vulnerabilities.

4. **Error Handling**: Enhance error handling in your backend functions to provide meaningful feedback to the frontend.

5. **Logging**: Implement logging for better debugging and monitoring.

6. **CORS Configuration**: Ensure your serverless functions have the correct CORS settings if your frontend and backend are on different domains.

---

## Troubleshooting Tips

- **Function Not Found**: Ensure your functions are correctly named and placed in the designated functions directory.
- **Database Connection Issues**: Verify the path to `rsvp.db` and ensure the database file is included in the deployment. Remember that serverless functions may have ephemeral file systems.
- **Permission Errors**: Ensure your function binaries have executable permissions.
- **CORS Issues**: If the frontend cannot communicate with the backend, check CORS settings in your functions.

---
