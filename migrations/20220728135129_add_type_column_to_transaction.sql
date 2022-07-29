-- +goose Up
-- +goose StatementBegin
CREATE TYPE transaction_type AS ENUM ('deposit', 'withdraw', 'transfer');

ALTER TABLE IF EXISTS account_transaction
    ADD COLUMN IF NOT EXISTS transaction_type transaction_type;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS account_transaction
    DROP COLUMN IF EXISTS transaction;

DROP TYPE transaction_type;
-- +goose StatementEnd
