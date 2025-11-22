-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TYPE order_type AS ENUM ('limit', 'market');

CREATE TYPE side_type AS ENUM('bid', 'ask')

CREATE TABLE IF NOT EXISTS (
	id 	   BIGSERIAL PRIMARY KEY,
	user_id    BIGINT NOT NULL,
	price      NUMERIC(18,8) NOT NULL,
	amount     NUMERIC(18,8) NOT NULL,
	remaining  NUMERIC(18,8) NOT NULL,
	side       side_type NOT NULL,
	type       order_type NOT NULL,
	created_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPZ NOT NULL DEFAULT NOW(),
)

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
