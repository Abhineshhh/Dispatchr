# Dispatchr

Dispatchr is a fast, concurrent email dispatcher built in Go. It demonstrates the power of Go's concurrency model by utilizing the **Worker Pool pattern** to process and send bulk emails efficiently.

## Features

- **High-Performance Concurrency**: Uses Goroutines and Channels to process emails in parallel across multiple workers.
- **Worker Pool Pattern**: Controls resource usage by limiting the number of active background workers executing tasks.
- **Efficient Templating**: Parses HTML/Text templates once at startup rather than on every email, drastically reducing disk I/O and CPU overhead.
- **Robust Error Handling**: Gracefully handles failed email deliveries and malformed CSV rows without crashing the entire service.
- **CSV Processing**: Dynamically loads recipient names and emails from a CSV file.

## Technologies Used

- **Language:** Go
- **Core Concepts:** Goroutines, Channels, `sync.WaitGroup`, `net/smtp`, `html/template`, `encoding/csv`

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) installed (1.18+ recommended)
- [Docker](https://docs.docker.com/get-docker/) installed for running the local SMTP testing environment.

### Setup and Execution

1. **Clone the repository** (or navigate to the workspace):
   ```bash
   cd Dispatchr
   ```

2. **Start the Mailpit SMTP Server**:
   Mailpit acts as a dummy SMTP server to catch outgoing emails so no real emails are accidentally sent during development.
   ```bash
   docker run -d --restart unless-stopped --name=mailpit -p 8025:8025 -p 1025:1025 axllent/mailpit
   ```
   *You can view the incoming emails in its web UI at [http://localhost:8025](http://localhost:8025).*

3. **Prepare your data**:
   Ensure you have an `emails.csv` file in the root directory.
   ```csv
   Name,Email
   John Doe,john@example.com
   Jane Smith,jane@example.com
   ```

4. **Prepare your template**:
   Ensure you have an `email.tmpl` file with your email structure.
   ```html
   To: {{.Email}}
   Subject: Hello {{.Name}}!

   Welcome to Dispatchr, {{.Name}}!
   ```

5. **Run the application**:
   ```bash
   go run .
   ```

## Learning Outcomes
This project was built to master Go's concurrency primitives, specifically learning how to safely pass data between goroutines using channels and synchronously wait for task completion using WaitGroups.
