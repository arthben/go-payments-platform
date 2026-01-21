CREATE TYPE owner_types as ENUM (
    'USER',
    'MERCHANT',
    'SYSTEM'
);

CREATE TABLE accounts (
    id          UUID                       PRIMARY KEY,
    owner_type  owner_types                NOT NULL,
    owner_id    TEXT                       NOT NULL,
    is_deleted  BOOL
    created_at  TIMESTAMP DEFAULT now()    NOT NULL,
    updated_at  TIMESTAMP DEFAULT now()    NOT NULL
);

CREATE UNIQUE INDEX idx_accounts_owner
ON accounts (owner_type, owner_id);

CREATE TRIGGER accounts_set_updated_at
    BEFORE UPDATE
    ON accounts
    FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_at();
