CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    eth_address VARCHAR(42) UNIQUE NOT NULL
);

INSERT INTO users(eth_address)
VALUES ('0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266')