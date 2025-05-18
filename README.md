
---

# Sales Analytics Project

## Project Overview

This project processes sales data from CSV files and loads it into a MariaDB/MySQL database using Golang 1.24. It manages customers, products, orders, and order items, and offers REST APIs to fetch sales analytics such as top products overall, by category, and by region. The system also supports automated periodic data refresh from CSV via a scheduler.

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

| Field                                       | Type            | Details                |
| ------------------------------------------- | --------------- | ---------------------- |
| order\_id                                   | bigint unsigned | Foreign Key (Orders)   |
| product\_id                                 | varchar(191)    | Foreign Key (Products) |
| quantity\_sold                              | bigint          |                        |
| unit\_price                                 | double          |                        |
| discount                                    | double          |                        |
| **Primary Key:** (`order_id`, `product_id`) |                 |                        |

---

## Key Features

* **CSV Data Upload:** Parses and bulk loads large CSV sales data into the database.
* **Auto Refresh Scheduler:** Periodically refreshes database data every configurable interval (default every 10 hours).
* **REST API Endpoints:**

  * `POST /api/refresh-data` → Manually trigger CSV data refresh.
  * `GET /api/top-products` → Retrieve top-selling products overall.
  * `GET /api/top-products/category` → Retrieve top-selling products by category.
  * `GET /api/top-products/region` → Retrieve top-selling products by region.
* **Logging:** Custom logging for refresh operations and errors.

---

## Technologies Used

* Golang 1.24
* Gorilla Mux (HTTP routing)
* GORM (ORM for database interaction)
* MariaDB/MySQL (Relational database)
* TOML (Configuration management)
* gocsv (CSV parsing)
* Custom logging helpers

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

Make sure your MariaDB/MySQL server is running and the database specified in `dbname` exists.

Run any necessary migrations (if you have migration scripts), or create tables as per the schema above.

### 4. Run the application

```bash
go run cmd/main.go
```

The server will start on the configured port (default: 8080).

---

## API Usage

* **POST** `/api/refresh-data`
  Trigger manual refresh of sales data from the CSV file.

* **GET** `/api/top-products`
  Fetch overall top N products (optionally filter by date range and limit).

* **GET** `/api/top-products/category`
  Fetch top N products within a specific category.

* **GET** `/api/top-products/region`
  Fetch top N products sold in a specific region.

Query parameters example for top products APIs:

```
?start=2023-01-01&end=2023-01-31&limit=10&category=Electronics&region=North
```

---

## Additional Notes

* The application auto-refreshes data every N hours (default 10 hours) by reloading the CSV file. This is configurable in `config.toml`.
* Logging files are generated to track processing and errors.
