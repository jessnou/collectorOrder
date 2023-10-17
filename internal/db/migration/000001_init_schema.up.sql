-- Таблица для хранения информации о стеллажах
CREATE TABLE shelves (
    shelf_id SERIAL PRIMARY KEY,
    shelf_name VARCHAR(255)
);

-- Таблица для хранения заказов
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(255),
    order_date DATE
);

-- Таблица для хранения товаров на складе
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255)
);

-- Таблица для связи продуктов и стеллажей
CREATE TABLE product_shelf (
    product_id INT,
    shelf_id INT,
    main_shelf BOOLEAN,
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (shelf_id) REFERENCES shelves(shelf_id)
);

-- Таблица для связи продуктов и заказаов
CREATE TABLE products_orders
(
    order_id         INT,
    product_id       INT,
    quantity         INT,
    FOREIGN KEY (order_id) REFERENCES orders (order_id),
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

INSERT INTO shelves (shelf_name) VALUES
    ('A'),
    ('Б'),
    ('В'),
    ('Г'),
    ('Д'),
    ('Е'),
    ('Ж'),
    ('З');


INSERT INTO products (product_name) VALUES
    ('Ноутбук'),
    ('Телевизор'),
    ('Телефон'),
    ('Системный блок'),
    ('Часы'),
    ('Микрофон');

INSERT INTO orders (customer_name, order_date) VALUES
    ('Иван Иванов', '2023-10-10'),
    ('Петр Петров', '2023-10-11'),
    ('Анна Сидорова', '2023-10-14'),
    ('Мария Козлова', '2023-10-15'),
    ('Александр Смирнов', '2023-10-16'),
    ('Екатерина Иванова', '2023-10-17'),
    ('Андрей Кузнецов', '2023-10-18'),
    ('Татьяна Павлова', '2023-10-19'),
    ('Иван Иванов', '2023-10-10'),
    ('Петр Петров', '2023-10-11'),
    ('Анна Сидорова', '2023-10-14'),
    ('Мария Козлова', '2023-10-15'),
    ('Александр Смирнов', '2023-10-16'),
    ('Екатерина Иванова', '2023-10-17'),
    ('Андрей Кузнецов', '2023-10-18'),
    ('Татьяна Павлова', '2023-10-19');

INSERT INTO product_shelf (product_id, shelf_id,main_shelf) VALUES
    (1, 1,true),
    (2, 1,true),
    (3, 2,true),
    (3, 3,false),
    (3, 8,false),
    (4, 7,true),
    (5, 7,true),
    (5, 1,false),
    (6,7,true);

INSERT INTO products_orders (order_id, product_id, quantity) VALUES
    (10, 1,2),
    (10,3,1),
    (10,6,1),
    (11, 2,3),
    (14, 1,3),
    (14, 4,4),
    (15, 5,1);