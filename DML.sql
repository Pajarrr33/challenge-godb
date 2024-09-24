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