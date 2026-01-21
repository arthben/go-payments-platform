CREATE TYPE payment_transaction_statuses as ENUM (
    'SUCCESS',
    'FAILED'
);

CREATE TABLE payment_transactions (
    id              UUID                            PRIMARY KEY,
    order_id        TEXT                            NOT NULL,
    account_id      UUID                            NOT NULL REFERENCES accounts(id),
    currency        CHAR(3)                         NOT NULL,
    amount          BIGINT                          NOT NULL,
    status          payment_transaction_statuses    NOT NULL,
    failure_reason  TEXT,
    created_at      TIMESTAMP DEFAULT now()         NOT NULL,
    updated_at      TIMESTAMP DEFAULT now()         NOT NULL
);

CREATE UNIQUE INDEX idx_payment_order
ON payment_transactions (order_id);

CREATE TRIGGER payment_transactions_set_updated_at
    BEFORE UPDATE
    ON payment_transactions
    FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_at();
