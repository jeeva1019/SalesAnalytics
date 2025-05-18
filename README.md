# Sales Analytics Project

## Project Overview

This project processes sales data from CSV files and loads it into a MariaDB/MySQL database using **Golang 1.24**. It manages customers, products, orders, and order items, and provides REST APIs to fetch sales analytics such as top-selling products overall, by category, and by region. The system also supports automated, periodic data refresh from CSV via a scheduler.

---

## Database Schema

### 1. Customers

| Field   | Type         | Details     |
| ------- | ------------ | ----------- |
| id      | varchar(191) | Primary Key |
| name    | longtext     |             |
| email   | longtext     |             |
| address | longtext     |             |

### 2. Products

| Field         | Type         | Details     |
| ------------- | ------------ | ----------- |
| product\_id   | varchar(191) | Primary Key |
| product\_name | longtext     |             |
| category      | longtext     |             |

### 3. Orders

| Field           | Type            | Details                 |
| --------------- | --------------- | ----------------------- |
| order\_id       | bigint unsigned | Primary Key             |
| customer\_id    | varchar(191)    | Foreign Key (Customers) |
| date\_of\_sale  | datetime(3)     |                         |
| region          | longtext        |                         |
| shipping\_cost  | double          |                         |
| payment\_method | longtext        |                         |

### 4. Order Details

| Field          | Type            | Details                |
| -------------- | --------------- | ---------------------- |
| order\_id      | bigint unsigned | Foreign Key (Orders)   |
| product\_id    | varchar(191)    | Foreign Key (Products) |
| quantity\_sold | bigint          |                        |
| unit\_price    | double          |                        |
| discount       | double          |                        |

---

## Key Features

* **CSV Data Upload:** Parses and bulk-loads large CSV sales data into the database.
* **Auto-Refresh Scheduler:** Automatically refreshes database data at configurable intervals (default: every 10 hours).
* **REST API Endpoints:**

  * `POST /api/refresh-data` → Manually trigger CSV data refresh.
  * `GET /api/top-products` → Get top-selling products overall.
  * `GET /api/top-products/category` → Get top-selling products by category.
  * `GET /api/top-products/region` → Get top-selling products by region.
* **Logging:** Logs CSV processing and error details to files.

---

## Technologies Used

* **Golang 1.24**
* **Gorilla Mux** – HTTP routing
* **GORM** – ORM for database interaction
* **MariaDB/MySQL** – Relational database
* **gocsv** – CSV parsing
* **TOML** – Configuration management
* **Custom Logging Helpers**

---

## Setup & Installation

### 1. Clone the repository

```bash
git clone https://github.com/jeeva1019/SalesAnalytics.git
cd SalesAnalytics
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Prepare your database

Ensure that your MariaDB/MySQL server is running and the database specified in `config.toml` exists.


### 4. Run the application

```bash
go run cmd/main.go
```

The server will start on the configured port (default: **8080**).

---

## API Usage

### 1. `POST /api/refresh-data`

Manually trigger a refresh of sales data from the configured CSV file.

### 2. `GET /api/top-products`

Retrieve top N selling products.

### 3. `GET /api/top-products/category`

Retrieve top N products within a specific category.

### 4. `GET /api/top-products/region`

Retrieve top N products sold in a specific region.

#### Query Parameters (optional)

```
?start=2023-01-01&end=2023-01-31&limit=10&category=Electronics&region=North
```

* `start`, `end` → Date range
* `limit` → Number of top products
* `category` → Filter by category
* `region` → Filter by region

---

## Additional Notes

* The application auto-refreshes CSV data at regular intervals (default: every 10 hours). This interval is configurable in the `config.toml` file.
* Log files are written to the `log/` directory. Ensure this folder exists before running the application.
* Use the manual `POST /api/refresh-data` endpoint to trigger a one-time refresh as needed.

