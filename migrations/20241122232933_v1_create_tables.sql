-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'users') THEN
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            balance DECIMAL(10, 2) NOT NULL DEFAULT 0
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'lots') THEN
        CREATE TABLE lots (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            description TEXT,
            start_price DECIMAL(10, 2) NOT NULL,
            seller_id INT REFERENCES users(id) ON DELETE CASCADE
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'bids') THEN
        CREATE TABLE bids (
            id SERIAL PRIMARY KEY,
            amount DECIMAL(10, 2) NOT NULL,
            lot_id INT REFERENCES lots(id) ON DELETE CASCADE,
            bidder_id INT REFERENCES users(id) ON DELETE CASCADE
        );
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'auctions') THEN
        CREATE TABLE auctions (
            id SERIAL PRIMARY KEY,
            lot_id INT UNIQUE REFERENCES lots(id) ON DELETE CASCADE,
            start_time TIMESTAMP NOT NULL,
            end_time TIMESTAMP NOT NULL
        );
    END IF;
END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS lots;
DROP TABLE IF EXISTS bids;
DROP TABLE IF EXISTS auctions;
-- +goose StatementEnd