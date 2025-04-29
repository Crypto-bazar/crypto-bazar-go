CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    eth_address VARCHAR(42) UNIQUE NOT NULL,
    avatar_url VARCHAR(255) DEFAULT 'uploads/avatars/default-avatar.jpg'
);

INSERT INTO
    users (eth_address)
VALUES (
        '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266'
    ),
    (
        '0x70997970C51812dc3A010C7d01b50e0d17dc79C8'
    );