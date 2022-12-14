CREATE DATABASE cylo;
USE cylo;

CREATE TABLE Products (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO Products (name, price) VALUES
    ('Apple', 1.00),
    ('Orange', 1.50),
    ('Banana', 2.00),
    ('Pineapple', 3.00),
    ('Mango', 4.00);