CREATE DATABASE products_rest_api;
USE products_rest_api;

CREATE TABLE products(
    id INTEGER AUTO_INCREMENT,
    sku VARCHAR(255) NOT NULL,
    pname VARCHAR(255) NOT NULL,
    category VARCHAR(20),
    price INTEGER NOT NULL,
    UNIQUE KEY (sku),
    KEY (category),
    PRIMARY KEY (id)
);