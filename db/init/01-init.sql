CREATE TABLE IF NOT EXISTS products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS orders (
    order_id VARCHAR(64) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL
);

INSERT INTO products (name, description, price, stock)
SELECT 'Laptop DevOps', 'Ordinateur pour environnement DevOps', 1299.00, 10
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'Laptop DevOps'
);

INSERT INTO products (name, description, price, stock)
SELECT 'Clavier Mecanique', 'Clavier mecanique pour developpeur', 99.90, 25
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'Clavier Mecanique'
);

INSERT INTO products (name, description, price, stock)
SELECT 'Ecran 27 pouces', 'Ecran QHD pour poste de travail', 349.00, 15
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'Ecran 27 pouces'
);
