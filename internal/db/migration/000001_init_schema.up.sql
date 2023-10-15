-- Таблица для хранения информации о стеллах
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
    product_name VARCHAR(255),
    main_shelf_id INT,
    FOREIGN KEY (main_shelf_id) REFERENCES shelves(shelf_id)
);

-- Таблица для связи заказов и товаров
CREATE TABLE order_products (
    order_id INT,
    product_id INT,
    quantity INT,
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id)
);

-- Вставьте данные в таблицу "shelves" - информация о стеллах
INSERT INTO shelves (shelf_name) VALUES
    ('A'),
    ('B'),
    ('C'),
    ('D'),
    ('E'),
    ('F'),
    ('G'),
    ('H');

-- Вставьте данные в таблицу "products" - информация о товарах на стеллах
INSERT INTO products (product_name, main_shelf_id) VALUES
    ('Ноутбук', 1),
    ('Монитор', 1),
    ('Телефон', 2),
    ('Системный блок', 3),
    ('Часы', 4),
    ('Микрофон', 4),
    ('Клавиатура', 1),
    ('Мышь', 1),
    ('Наушники', 2),
    ('Компьютерный стол', 3),
    ('Камера', 4),
    ('Принтер', 4),
    ('Флеш-диск', 5),
    ('Видеокарта', 6),
    ('SSD накопитель', 6),
    ('Беспроводная клавиатура', 5),
    ('Беспроводная мышь', 5);

-- Вставьте данные в таблицу "orders" - информация о заказах
INSERT INTO orders (customer_name, order_date) VALUES
    ('Иван Иванов', '2023-10-10'),
    ('Петр Петров', '2023-10-11'),
    ('Анна Сидорова', '2023-10-14'),
    ('Мария Козлова', '2023-10-15'),
    ('Александр Смирнов', '2023-10-16'),
    ('Екатерина Иванова', '2023-10-17'),
    ('Андрей Кузнецов', '2023-10-18'),
    ('Татьяна Павлова', '2023-10-19');

-- Вставьте данные в таблицу "order_products" - связь заказов и товаров
-- Здесь необходимо указать product_id и quantity для каждого товара в заказе
INSERT INTO order_products (order_id, product_id, quantity) VALUES
    (1, 1, 2),
    (1, 3, 1),
    (2, 2, 1),
    (2, 4, 1),
    (3, 3, 3),
    (3, 5, 1),
    (4, 1, 1),
    (4, 2, 1),
    (4, 3, 1),
    (5, 1, 2),
    (5, 4, 1),
    (6, 2, 1),
    (6, 5, 1),
    (7, 3, 3),
    (7, 1, 1),
    (8, 1, 2),
    (8, 2, 1),
    (8, 3, 1);