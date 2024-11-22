-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0
);

CREATE TABLE lots (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    start_price DECIMAL(10, 2) NOT NULL,
    seller_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE bids (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(10, 2) NOT NULL,
    lot_id INT REFERENCES lots(id) ON DELETE CASCADE,
    bidder_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE auctions (
    id SERIAL PRIMARY KEY,
    lot_id INT UNIQUE REFERENCES lots(id) ON DELETE CASCADE,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop TABLE users;
drop TABLE lots;
drop TABLE bids;
drop TABLE auctions;
-- +goose StatementEnd
