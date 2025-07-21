ALTER TABLE if EXISTS "accounts" DROP CONSTRAINT if EXISTS "owner_currency_key";

ALTER TABLE if EXISTS "accounts" DROP CONSTRAINT if EXISTS "accounts_owner_fkey";

DROP TABLE if EXISTS "users" 