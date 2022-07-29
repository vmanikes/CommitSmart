-- +goose Up
-- +goose StatementBegin
INSERT INTO account_user (account_number, first_name, last_name)
    VALUES (1234, 'John', 'Doe'), (5678, 'John', 'Wick');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- No down migration as we do not want to delete these users
-- +goose StatementEnd
