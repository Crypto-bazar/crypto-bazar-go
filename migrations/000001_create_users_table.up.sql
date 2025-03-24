CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    eth_address VARCHAR(42) UNIQUE NOT NULL
);
