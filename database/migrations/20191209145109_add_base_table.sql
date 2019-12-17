-- +goose Up
-- +goose StatementBegin
CREATE TABLE history (
  id uuid primary key default gen_random_uuid() UNIQUE,
  user_id varchar(255)  NOT NULL,
  latitude varchar(255)  NOT NULL,
  longitude varchar(255)  NOT NULL,
  created_at timestamptz default now(),
  updated_at timestamptz default now(),
  deleted_at timestamptz
);

create trigger update_history_update_at
before update on history for each row execute procedure update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS history;
-- +goose StatementEnd
