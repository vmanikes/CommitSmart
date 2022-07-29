-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS account_transaction_transac_type_idx
    ON account_transaction(transaction_type);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS account_transaction_transac_type_idx;
-- +goose StatementEnd
