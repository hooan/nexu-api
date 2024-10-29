CREATE TABLE models (
    id SERIAL PRIMARY KEY,
    brand_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    average_price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (brand_id) REFERENCES brands(id),
    UNIQUE (brand_id, name)
);