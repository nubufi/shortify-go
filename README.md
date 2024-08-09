# shortify-go
shortify-go is a simple, efficient URL shortener built with Golang and htmx. It allows users to quickly generate short, shareable links from long URLs, featuring dynamic frontend interactions. Perfect for anyone looking to deploy a lightweight URL shortening service.

## Stack

- Golang(echo) for backend
- PostgreSql for database
- Htmx + TailwindCss for frontend

## Features

- **URL Shortening:** Convert long URLs into shorter ones.
- **Redirection:** Redirect users from the short URL to the original URL.
- **Dynamic Frontend:** Interact with the service using a responsive interface without full-page reloads, thanks to htmx.
- **Database Support:** Store URLs and their corresponding short codes in a relational database (PostgreSQL, MySQL, SQLite, etc.).

## Getting Started

### Prerequisites

- Docker & Docker compose.

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/nubufi/shortify.git
   cd shortify
   ```
2. **Run docker instances:**

    ```bash
    docker-compose up --build
    ```

3. **Open localhost:**
    
    Open http://localhost:8080 on your browser
