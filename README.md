# Backend_SDAB

This project is the backend service for the SDAB. It provides APIs for data processing, storage, and retrieval. The backend is built using Go and follows best practices for scalability and maintainability.

## Features
- RESTful API endpoints for data operations.
- Secure authentication and authorization mechanisms (*Up coming*).
- Scalable architecture with support for horizontal scaling.
- Comprehensive logging and monitoring (*almost there/Up coming*).

## Requirements
- Go 1.18+
- PostgreSQL 12+
- Redis (optional, for caching)

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/userNanni/Backend_SDAB.git
    ```
2. Navigate to the project directory:
    ```bash
    cd Backend_SDAB
    ```
3. Build the project:
    ```bash
    go build
    ```
4. Configure environment variables as per `.env`.

`.env` Example:
host: "192.0.0.1"
port: "3306"
user: "user"
password: "password"
dbname: "your_db"

## Usage
Start the development server:
```bash
./Backend_SDAB
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
