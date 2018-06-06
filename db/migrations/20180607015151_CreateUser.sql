
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user(first text, last text);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user;
