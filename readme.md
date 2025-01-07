# Todo Application API

A simple Todo API built with Go, Fiber, and GORM. The app allows users to create, update, delete, and toggle todos, with Bearer token (Not implemented).

## Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/md-bz/Todo-go
   cd Todo-go
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the application**:
   ```bash
   go run main.go
   ```
   The server starts on `localhost:3000`.

## API Endpoints

### 1. Get all todos
- **GET** `/`
- Fetch all todos for the authenticated user.

### 2. Create a new todo
- **POST** `/`
- **Body**: `{ "description": "todo description" }`

### 3. Update a todo
- **PATCH** `/`
- **Body**: `{ "oldDescription": "old desc", "newDescription": "new desc" }`

### 4. Delete a todo
- **DELETE** `/`
- **Body**: `{ "description": "todo description" }`

### 5. Toggle todo completion
- **POST** `/toggle`
- **Body**: `{ "description": "todo description" }`