-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS account_user_acc_num_idx
ON account_user(account_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS account_user_acc_num_idx;
-- +goose StatementEnd
