-- CREATE DATABASE cylo;
-- USE cylo;

ALTER DATABASE cylo SET timezone TO 'Asia/Ho_Chi_Minh';

CREATE TABLE Products (
    id SERIAL NOT NULL PRIMARY KEY ,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    images TEXT [],
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL
);

INSERT INTO Products (name, price, images) VALUES
    ('Apple', 1.00, '{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F0%2F0.jpg?alt=media&token=7b40f3cc-37a2-414a-8550-cc8b2389dda3", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F0%2F1.jpg?alt=media&token=8faa9613-e27d-4193-91ec-615d2d53dfe5"}'),
    ('Orange', 1.50, '{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F1%2F0.jpg?alt=media&token=cfb5ea77-7215-45a1-946c-78f281fdd4b6", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F1%2F1.jpg?alt=media&token=7ed979aa-2992-4d97-a862-c98007312ad6"}'),
    ('Banana', 2.00, '{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F10%2F0.jpg?alt=media&token=9baece8e-5bb2-458f-b454-9308c6afebf2", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F10%2F1.jpg?alt=media&token=fb028a35-6ddf-4c3c-9634-2db90d202756"}'),
    ('Pineapple', 3.00, '{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F11%2F0.jpg?alt=media&token=909e31c1-38bc-4de1-ad08-4f5b52f9bad1", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F11%2F1.jpg?alt=media&token=9f2a9325-4c9e-4d04-85f2-a455add0b901"}'),
    ('Mango', 4.00, '{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F12%2F0.jpg?alt=media&token=a4379223-24db-4184-89de-dcf4cb089a15", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F12%2F1.jpg?alt=media&token=9d705a38-8bcf-4daa-9043-5e073e4e0d5a"}');