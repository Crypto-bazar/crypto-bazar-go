CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    eth_address CHAR(42) UNIQUE NOT NULL
);
