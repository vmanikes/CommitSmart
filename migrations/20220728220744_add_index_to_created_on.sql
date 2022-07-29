-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS account_transaction_created_on_idx
    ON account_transaction(created_on);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS account_transaction_created_on_idx;
-- +goose StatementEnd
