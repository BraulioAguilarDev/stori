CREATE TABLE IF NOT EXISTS "profile" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "firebase" VARCHAR(28) NOT NULL,
  "name" VARCHAR(100) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id),
  UNIQUE(email, firebase)
);


CREATE TABLE IF NOT EXISTS "bank_account" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "name" VARCHAR NOT NULL,
  "type" VARCHAR NOT NULL,
  "number" INTEGER NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "profile_id" uuid NOT NULL,
  PRIMARY KEY(id),
  CONSTRAINT fk__bank_account__profile_id__profile__id
    FOREIGN KEY(profile_id)
    REFERENCES profile(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS "transaction" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "bank_account_id" uuid NOT NULL,
  "date" TIMESTAMP NOT NULL,
  "debit_amount" NUMERIC(10,2) DEFAULT 0.0,
  "credit_amount" NUMERIC(10,2) DEFAULT 0.0,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY(id),
  CONSTRAINT fk__transaction__account_id__bank_account__id
    FOREIGN KEY(bank_account_id)
    REFERENCES bank_account(id) ON DELETE CASCADE
);