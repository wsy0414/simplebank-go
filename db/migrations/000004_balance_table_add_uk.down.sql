ALTER TABLE balance
DROP CONSTRAINT uk_user_id_currency UNIQUE (user_id, currency);