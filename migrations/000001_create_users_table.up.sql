CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    eth_address VARCHAR(42) UNIQUE NOT NULL,
    avatar_url VARCHAR(255)
);

INSERT INTO users(eth_address)
VALUES ('0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266')