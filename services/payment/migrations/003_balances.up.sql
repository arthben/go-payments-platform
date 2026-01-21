CREATE TABLE balances (
    account_id  UUID                    PRIMARY KEY REFERENCES accounts(id),
    currency    CHAR(3)                 NOT NULL,
    amount      BIGINT                  NOT NULL DEFAULT 0, -- smallest unit
    created_at  TIMESTAMP DEFAULT now() NOT NULL,
    updated_at  TIMESTAMP DEFAULT now() NOT NULL
);

CREATE INDEX idx_balances_currency
ON balances (currency);

CREATE TRIGGER balances_set_updated_at
    BEFORE UPDATE
    ON balances
    FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_at();