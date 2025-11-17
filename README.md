# Olist Test - Shipping Management API

REST API developed in Go for managing packages, transport companies, and shipping quotes. The project uses Gin as the web framework, GORM as the ORM, and PostgreSQL as the database.

## 📋 Table of Contents

- [Technologies](#technologies)
- [Project Structure](#project-structure)
- [Models](#models)
- [API Routes](#api-routes)
- [How to Run](#how-to-run)
- [Environment Variables](#environment-variables)
- [Usage Examples](#usage-examples)
- [Architecture](#architecture)

## 🛠 Technologies

- **Go 1.25.4** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM for Go
- **PostgreSQL 16** - Relational database
- **Docker & Docker Compose** - Containerization

## 📁 Project Structure

```
olist-test/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # Application configuration
│   │   └── config.go            # Environment variables loading
│   ├── database/                # Database configuration
│   │   └── postgres.go          # PostgreSQL initialization and migrations
│   ├── entities/                # Domain models (entities)
│   │   ├── package.go           # Package entity
│   │   ├── shipment.go          # Shipment entity
│   │   └── transport_company.go # TransportCompany and TransportCompanyRegion entities
│   ├── dto/                     # Data Transfer Objects
│   │   ├── package.go           # Package-related DTOs
│   │   ├── shipment.go          # Shipment-related DTOs
│   │   ├── shipping_quote.go    # Shipping quote DTOs
│   │   └── transport_company*.go # Transport company-related DTOs
│   ├── handlers/                # HTTP handlers (controllers)
│   │   ├── package.go           # Package handlers
│   │   ├── shipment.go          # Shipment handlers
│   │   ├── shipping_quote.go    # Shipping quote handler
│   │   ├── transport_company.go # Transport company handlers
│   │   └── transport_company_region.go # Region handlers
│   ├── repositories/            # Data access layer
│   ├── services/                # Business logic
│   ├── router/                  # Route configuration
│   │   └── router.go            # API routes setup
│   └── utils.go                 # Utilities
├── docker/                      # Docker scripts
│   └── postgres/
│       └── init.sql             # PostgreSQL initialization script
├── docker-compose.yaml          # Docker Compose configuration
├── Dockerfile                   # Docker image for the application
├── go.mod                       # Project dependencies
└── go.sum                       # Dependency checksums
```

## 🗄 Models

### Package
Represents a package to be shipped.

**Fields:**
- `ID` (UUID) - Unique identifier
- `Status` (string) - Package status (default: "pending")
- `Product` (string) - Product name
- `Weight` (float32) - Weight in kg
- `DestinationState` (string) - Destination state
- `CreatedAt` (timestamp) - Creation date
- `UpdatedAt` (timestamp) - Update date

### TransportCompany
Represents a transport company.

**Fields:**
- `ID` (UUID) - Unique identifier
- `Name` (string) - Company name
- `Regions` ([]TransportCompanyRegion) - Served regions
- `CreatedAt` (timestamp) - Creation date
- `UpdatedAt` (timestamp) - Update date

### TransportCompanyRegion
Represents a region served by a transport company with its prices and delivery times.

**Fields:**
- `ID` (uint) - Unique identifier
- `TransportCompanyID` (UUID) - Transport company ID
- `Name` (string) - Region name (e.g., "SP", "RJ")
- `EstimatedDays` (int) - Estimated delivery days
- `PricePerKg` (float64) - Price per kg
- `CreatedAt` (timestamp) - Creation date
- `UpdatedAt` (timestamp) - Update date

### Shipment
Represents a shipment of a package by a transport company.

**Fields:**
- `ID` (UUID) - Unique identifier
- `PackageID` (UUID) - Package ID
- `TransportCompanyID` (UUID) - Transport company ID
- `TrackingCode` (string) - Tracking code (unique)
- `Price` (float64) - Shipping price
- `EstimatedDays` (int) - Estimated delivery days
- `EstimatedDeliveryAt` (timestamp) - Estimated delivery date
- `CreatedAt` (timestamp) - Creation date

## 🛣 API Routes

All routes are prefixed with `/api/v1`.

### Packages

#### Create Package
```
POST /api/v1/packages
```

**Body:**
```json
{
  "product": "Notebook",
  "weight": 2.5,
  "destination_state": "SP"
}
```

**Response:** 201 Created
```json
{
  "id": "uuid",
  "status": "pending",
  "product": "Notebook",
  "weight": 2.5,
  "destination_state": "SP",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### List Packages
```
GET /api/v1/packages?status=pending&product=Notebook
```

**Query Parameters (optional):**
- `status` - Filter by status
- `product` - Filter by product

**Response:** 200 OK
```json
[
  {
    "id": "uuid",
    "status": "pending",
    "product": "Notebook",
    "weight": 2.5,
    "destination_state": "SP",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Package by ID
```
GET /api/v1/packages/:id
```

**Response:** 200 OK (same format as create)

#### Update Package
```
PUT /api/v1/packages/:id
```

**Body (all fields optional):**
```json
{
  "status": "shipped",
  "product": "Notebook Pro",
  "peso_kg": 3.0,
  "destination_state": "RJ"
}
```

**Response:** 200 OK (same format as create)

#### Delete Package
```
DELETE /api/v1/packages/:id
```

**Response:** 204 No Content

### Transport Companies

#### Create Transport Company
```
POST /api/v1/transport-companies
```

**Body:**
```json
{
  "name": "Correios"
}
```

**Response:** 201 Created
```json
{
  "id": "uuid",
  "name": "Correios",
  "regions": [],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### List Transport Companies
```
GET /api/v1/transport-companies
```

**Response:** 200 OK
```json
[
  {
    "id": "uuid",
    "name": "Correios",
    "regions": [...],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Get Transport Company by ID
```
GET /api/v1/transport-companies/:id
```

**Response:** 200 OK (same format as create)

#### Update Transport Company
```
PUT /api/v1/transport-companies/:id
```

**Body:**
```json
{
  "name": "Correios Express"
}
```

**Response:** 200 OK (same format as create)

#### Delete Transport Company
```
DELETE /api/v1/transport-companies/:id
```

**Response:** 204 No Content

### Transport Company Regions

#### Create Region for Transport Company
```
POST /api/v1/transport-companies/:id/regions
```

**Body:**
```json
{
  "name": "SP",
  "estimated_days": 5,
  "price_per_kg": 10.50
}
```

**Response:** 201 Created

#### List Regions of a Transport Company
```
GET /api/v1/transport-companies/:id/regions
```

**Response:** 200 OK

### Shipping Quotes

#### Calculate Quote
```
POST /api/v1/shipping/quote
```

**Body:**
```json
{
  "destination_state": "SP",
  "weight": 2.5
}
```

**Response:** 200 OK
```json
[
  {
    "transport_company_id": "uuid",
    "transport_company": "Correios",
    "price": 26.25,
    "estimated_days": 5
  }
]
```

Returns all transport companies that serve the destination state, ordered by price.

### Shipments

#### Create Shipment
```
POST /api/v1/shipments
```

**Body:**
```json
{
  "package_id": "uuid",
  "transport_company_id": "uuid",
  "price": 26.25,
  "estimated_days": 5
}
```

**Response:** 201 Created
```json
{
  "id": "uuid",
  "package_id": "uuid",
  "transport_company_id": "uuid",
  "price": 26.25,
  "estimated_days": 5,
  "estimated_delivery_at": "2024-01-06T00:00:00Z",
  "tracking_code": "ABC123XYZ"
}
```

## 🚀 How to Run

### Prerequisites

- **Docker** (version 20.10 or higher)
- **Docker Compose** (version 2.0 or higher)
- **Go 1.25.4+** (only if running locally without Docker)

### Option 1: Running with Docker Compose (Recommended)

This is the easiest way to run the project. Docker Compose will handle all the setup automatically.

#### Step 1: Clone the repository
```bash
git clone <repository-url>
cd olist-test
```

#### Step 2: Build and start the containers
```bash
docker-compose up --build
```

This command will:
- Build the Go application Docker image
- Start the PostgreSQL container
- Start the API container
- Run database migrations automatically
- Wait for PostgreSQL to be healthy before starting the API

#### Step 3: Verify the services are running

The API will be available at: `http://localhost:8080`

The PostgreSQL database will be available at: `localhost:5431`

You can test if the API is running by making a request:
```bash
curl http://localhost:8080/api/v1/packages
```

#### Step 4: Stop the services

To stop the containers:
```bash
docker-compose down
```

To stop and remove volumes (this will delete the database data):
```bash
docker-compose down -v
```

#### Additional Docker Compose Commands

- **Run in detached mode (background):**
  ```bash
  docker-compose up -d --build
  ```

- **View logs:**
  ```bash
  docker-compose logs -f api
  ```

- **Restart a specific service:**
  ```bash
  docker-compose restart api
  ```

- **Rebuild without cache:**
  ```bash
  docker-compose build --no-cache
  docker-compose up
  ```

### Option 2: Running Locally (without Docker)

If you prefer to run the application locally without Docker, follow these steps:

#### Step 1: Install PostgreSQL

Make sure PostgreSQL 16 is installed and running on your system.

#### Step 2: Create the database

Connect to PostgreSQL and create the database:
```bash
psql -U postgres
```

Then run:
```sql
CREATE DATABASE "olist-db";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
\q
```

#### Step 3: Set environment variables

Create a `.env` file or export the variables:
```bash
export PORT=8080
export DATABASE_URL="postgres://postgres:your_password@localhost:5432/olist-db?sslmode=disable"
```

Or create a `.env` file:
```env
PORT=8080
DATABASE_URL=postgres://postgres:your_password@localhost:5432/olist-db?sslmode=disable
```

#### Step 4: Install dependencies

```bash
go mod download
```

#### Step 5: Run the application

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`.

### Troubleshooting

#### Port already in use

If port 8080 is already in use, you can change it by:
- **Docker Compose:** Modify the port mapping in `docker-compose.yaml` (e.g., `"8081:8080"`)
- **Local:** Change the `PORT` environment variable

#### Database connection errors

- Make sure PostgreSQL is running
- Verify the `DATABASE_URL` is correct
- Check if the database exists
- Ensure the PostgreSQL user has the necessary permissions

#### Container build errors

- Make sure Docker is running
- Try rebuilding without cache: `docker-compose build --no-cache`
- Check if you have enough disk space

## 🔧 Environment Variables

The application uses the following environment variables:

- `PORT` - Port where the API will run (default: 8080)
- `DATABASE_URL` - PostgreSQL connection string

**Format:**
```
postgres://username:password@host:port/database?sslmode=disable
```

**Example:**
```bash
export PORT=8080
export DATABASE_URL="postgres://postgres:123456PW@localhost:5431/olist-db?sslmode=disable"
```

In Docker Compose, these variables are automatically configured.

## 📝 Usage Examples

### 1. Create a transport company
```bash
curl -X POST http://localhost:8080/api/v1/transport-companies \
  -H "Content-Type: application/json" \
  -d '{"name": "Correios"}'
```

### 2. Add a region to a transport company
```bash
curl -X POST http://localhost:8080/api/v1/transport-companies/{id}/regions \
  -H "Content-Type: application/json" \
  -d '{
    "name": "SP",
    "estimated_days": 5,
    "price_per_kg": 10.50
  }'
```

### 3. Create a package
```bash
curl -X POST http://localhost:8080/api/v1/packages \
  -H "Content-Type: application/json" \
  -d '{
    "product": "Notebook",
    "weight": 2.5,
    "destination_state": "SP"
  }'
```

### 4. Get shipping quotes
```bash
curl -X POST http://localhost:8080/api/v1/shipping/quote \
  -H "Content-Type: application/json" \
  -d '{
    "destination_state": "SP",
    "weight": 2.5
  }'
```

### 5. Create a shipment
```bash
curl -X POST http://localhost:8080/api/v1/shipments \
  -H "Content-Type: application/json" \
  -d '{
    "package_id": "uuid-do-pacote",
    "transport_company_id": "uuid-da-transportadora",
    "price": 26.25,
    "estimated_days": 5
  }'
```

## 🏗 Architecture

The project follows a layered architecture:

1. **Handlers** - Receive HTTP requests and return responses
2. **Services** - Contain business logic
3. **Repositories** - Manage data access
4. **Entities** - Domain models
5. **DTOs** - Data Transfer Objects

This separation of concerns makes the code more maintainable, testable, and scalable.

## 📄 License

This project is licensed under the terms specified in the LICENSE file.
