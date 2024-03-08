CREATE TABLE IF NOT EXISTS "profile" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR(100) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id),
  UNIQUE(name, email)
);

CREATE TABLE IF NOT EXISTS "account" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "owner" VARCHAR NOT NULL,
  "bank" VARCHAR NOT NULL,
  "type" VARCHAR NOT NULL,
  "number" INTEGER NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "profile_id" uuid NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk__account__profile_id__profile__id
    FOREIGN KEY(profile_id)
    REFERENCES profile(id) ON DELETE CASCADE,
  UNIQUE(bank, number)
);

CREATE TABLE IF NOT EXISTS "transaction" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "account_id" uuid NOT NULL,
  "date" TIMESTAMP NOT NULL,
  "debit_amount" NUMERIC(10,2) DEFAULT 0.0,
  "credit_amount" NUMERIC(10,2) DEFAULT 0.0,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id),
  CONSTRAINT fk__transaction__account_id__account__id
    FOREIGN KEY(account_id)
    REFERENCES account(id) ON DELETE CASCADE
);

ALTER TABLE account ADD CONSTRAINT constraint__bank__number UNIQUE (bank, number);

CREATE TABLE IF NOT EXISTS "account_s3" (
  "account_id" uuid NOT NULL,
  "url" VARCHAR NOT NULL,
  "filename" VARCHAR,
  CONSTRAINT fk__account_s3__account_id__account__id
    FOREIGN KEY(account_id)
    REFERENCES account(id) ON DELETE CASCADE
);
