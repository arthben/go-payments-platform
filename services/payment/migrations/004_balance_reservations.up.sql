CREATE TYPE balance_reservation_statuses as ENUM (
    'RESERVED',
    'CONFIRMED',
    'RELEASED'
);

CREATE TABLE balance_reservations (
    id          UUID                            PRIMARY KEY,
    account_id  UUID                            NOT NULL REFERENCES accounts(id),
    order_id    TEXT                            NOT NULL,
    currency    CHAR(3)                         NOT NULL,
    amount      BIGINT                          NOT NULL,
    status      balance_reservation_statuses    NOT NULL,
    created_at  TIMESTAMP DEFAULT now()         NOT NULL,
    updated_at  TIMESTAMP DEFAULT now()         NOT NULL
);

CREATE UNIQUE INDEX idx_reservation_order
ON balance_reservations (order_id);

CREATE INDEX idx_reservation_account
ON balance_reservations (account_id);

CREATE TRIGGER balance_reservations_set_updated_at
    BEFORE UPDATE
    ON balance_reservations
    FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_at();
