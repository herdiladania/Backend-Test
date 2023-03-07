CREATE DATABASE IF NOT EXISTS db_rumah_makan;

USE db_rumah_makan;

DROP TABLE IF EXISTS Payment;
DROP TABLE IF EXISTS Delivery;
DROP TABLE IF EXISTS `Order`;
DROP TABLE IF EXISTS Menu;
DROP TABLE IF EXISTS Customer;


CREATE TABLE Customer (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  address VARCHAR(255),
  email VARCHAR(255),
  phone VARCHAR(255),
  password VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE Menu (
  id INT PRIMARY KEY,
  name VARCHAR(255),
  price DECIMAL(10,2),
  description VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE `Order` (
  id INT PRIMARY KEY,
  customer_id INT,
  menu_id INT,
  quantity INT,
  total_price DECIMAL(10,2),
  order_time TIMESTAMP,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (customer_id) REFERENCES Customer(id),
  FOREIGN KEY (menu_id) REFERENCES Menu(id)
);

CREATE TABLE Payment (
  id INT PRIMARY KEY,
  order_id INT,
  method VARCHAR(255),
  total DECIMAL(10,2),
  status VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES `Order`(id)
);

CREATE TABLE Delivery (
  id INT PRIMARY KEY,
  order_id INT,
  `time` TIMESTAMP,
  driver VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES `Order`(id)
);
