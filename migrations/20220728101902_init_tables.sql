-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account_user (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_number BIGINT UNIQUE NOT NULL,
    first_name VARCHAR DEFAULT '',
    middle_name VARCHAR DEFAULT '',
    last_name VARCHAR DEFAULT '',
    balance NUMERIC DEFAULT 0,
    created_on TIMESTAMP DEFAULT NOW(),
    updated_on TIMESTAMP DEFAULT NULL,
    deleted_on TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS account_transaction (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid,
    balance NUMERIC NOT NULL,
    amount NUMERIC NOT NULL,
    created_on TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id)
        REFERENCES account_user (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account_transaction;
DROP TABLE IF EXISTS account_user;
-- +goose StatementEnd
