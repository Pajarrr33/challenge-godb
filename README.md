
# Challange Golang Database
This project is a simple Go-based application for managing a laundry shop. It provides CRUD (Create, Read, Update, Delete) operations for customers and services and includes validations for data consistency.



## Installation

To Run This project, ensure you have Go and SQL Database installed

1. Clone this repository

```bash
	https://git.enigmacamp.com/enigma-20/ahmad-fajar-shidik/challenge-godb.git
```

2. Create a database 

```bash
	CREATE DATABASE example_name
```

3. Run this DDL query or copy it from DDL.sql File

```bash
	CREATE TABLE customer (
		customer_id INT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		phone VARCHAR(255) NOT NULL,
		address VARCHAR(255) DEFAULT '',
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE service (
		service_id INT PRIMARY KEY,
		service_name VARCHAR(255) NOT NULL,
		unit VARCHAR(255) NOT NULL,
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE "order" (
		order_id INT PRIMARY KEY,
		customer_id INT NOT NULL,
		order_date TIMESTAMP NOT NULL,
		completion_date TIMESTAMP,
		received_by VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (customer_id) REFERENCES customer(customer_id)
	);

	CREATE TABLE order_detail (
		order_detail_id SERIAL PRIMARY KEY,
		order_id INT NOT NULL,
		service_id INT NOT NULL,
		qty INT NOT NULL,
		FOREIGN KEY (order_id) REFERENCES "order"(order_id),
		FOREIGN KEY (service_id) REFERENCES service(service_id)
	);
```

4. Run this DML query or copy it from DML.sql File
```bash
	INSERT INTO customer (customer_id, name, phone, address, created_at, updated_at)
	VALUES
	(1, 'John Doe', '+1234567890', '123 Main St, Anytown, CA 12345', '2023-11-22 10:00:00', '2023-11-22 10:00:00'),
	(2, 'Jane Smith', '+1987654321', '456 Elm St, Anytown, CA 12345', '2023-11-23 14:30:00', '2023-11-23 14:30:00'),
	(3, 'Michael Johnson', '+1555121212', '789 Oak Ave, Anytown, CA 12345', '2023-11-24 09:15:00', '2023-11-24 09:15:00'),
	(4, 'Emily Davis', '+1111222233', '101 Pine St, Anytown, CA 12345', '2023-11-25 16:45:00', '2023-11-25 16:45:00'),
	(5, 'Danuel Iskandar', '+6289612345678', '101 Metro Duta Depok 12312', '2024-09-13 11:11:11', '2024-09-13 11:11:11');

	INSERT INTO service (service_id, service_name, unit, price, created_at, updated_at)
	VALUES
	(1, 'Jasa Cuci Mobil', 'buah', 10000, '2023-11-22 10:00:00', '2023-11-22 10:00:00'),
	(2, 'Service AC', 'unit', 250000, '2023-11-23 14:30:00', '2023-11-23 14:30:00'),
	(3, 'Perbaikan Komputer', 'jam', 150000, '2023-11-24 09:15:00', '2023-11-24 09:15:00'),
	(4, 'Potong Rambut', 'orang', 50000, '2023-11-25 16:45:00', '2023-11-25 16:45:00'),
	(5, 'Les Privat', 'jam', 100000, '2023-11-26 12:00:00', '2023-11-26 12:00:00');

	INSERT INTO "order" (order_id, customer_id, order_date, completion_date,received_by,created_at,updated_at)
	VALUES 
	(1,1,'2024-09-13 11:11:11',NULL,'John Doe','2023-11-25 16:45:00','2023-11-25 16:45:00'),
	(2,2,'2024-09-13 11:11:11',NULL,'Jane Smith','2023-11-25 16:45:00','2023-11-25 16:45:00'),
	(3,3,'2024-09-13 11:11:11',NULL,'Micheal Jonshon','2023-11-25 16:45:00','2023-11-25 16:45:00'),
	(4,4,'2024-09-13 11:11:11',NULL,'Emily Davis','2023-11-25 16:45:00','2023-11-25 16:45:00'),
	(5,5,'2024-09-13 11:11:11',NULL,'Danuel Iskandar','2023-11-25 16:45:00','2023-11-25 16:45:00');

	INSERT INTO order_detail (order_detail_id, order_id, service_id, qty)
	VALUES 
	(1,1,1,3),
	(2,2,2,5),
	(3,3,3,1),
	(4,4,4,6),
	(5,5,5,11);
```

5. Configure Your database in env file and change the env file name to .env
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=username
DB_PASSWORD=password
DB_NAME=example_name
```

6. Navigate to the project directory
```bash
cd challenge-godb/enigma-laundry.go
```

7. Install necessary dependencies
```bash
go mod tidy
```

7. Run the application
```bash
go run main.go
```
    
## Features

- Main Menu
    - Customer
    - Service
    - Order

- Customer Menu
    - Create Customer
    - View List Of Customer
    - View Customer By Id
    - Update Customer
    - Delete Customer

- Service Menu
    - Create Service
    - View List Of Service
    - View Service by Id
    - Update Service
    - Delete Service

- Order Menu
    - Create Order
    - Complete Order
    - View List Of Order
    - View Order By Id
