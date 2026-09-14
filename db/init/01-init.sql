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

INSERT INTO products (name, description, price, stock)
SELECT 'Souris sans fil', 'Souris ergonomique pour travailler confortablement', 39.90, 40
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'Souris sans fil'
);

INSERT INTO products (name, description, price, stock)
SELECT 'SSD 1 To', 'Stockage SSD pour applications et machines virtuelles', 89.90, 30
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'SSD 1 To'
);

INSERT INTO products (name, description, price, stock)
SELECT 'Casque audio', 'Casque avec microphone pour les reunions', 59.90, 20
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'Casque audio'
);

INSERT INTO products (name, description, price, stock)
SELECT 'Station USB-C', 'Station multiports pour ordinateur portable', 79.90, 18
WHERE NOT EXISTS (
    SELECT 1 FROM products WHERE name = 'Station USB-C'
);
