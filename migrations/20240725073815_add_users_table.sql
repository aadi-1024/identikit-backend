-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Users (
    Id SERIAL PRIMARY KEY,
    Email VARCHAR UNIQUE NOT NULL,
    Password VARCHAR NOT NULL,
    Role VARCHAR NOT NULL,
    CHECK (Role in ('admin', 'user')) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE Users;
-- +goose StatementEnd
