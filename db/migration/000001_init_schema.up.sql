CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "wallet_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "expense_description" varchar,
  "category_id" bigint NOT NULL,
  "expense_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "budgets" (
  "id" bigserial PRIMARY KEY,
  "wallet_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallets" (
  "name" varchar NOT NULL,
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "expenses" ("wallet_id");

CREATE INDEX ON "expenses" ("category_id");

CREATE INDEX ON "expenses" ("expense_date");

CREATE INDEX ON "expenses" ("wallet_id", "expense_date");

CREATE INDEX ON "budgets" ("wallet_id");

CREATE INDEX ON "budgets" ("category_id");

CREATE INDEX ON "budgets" ("wallet_id", "category_id");

CREATE INDEX ON "wallets" ("owner");

COMMENT ON COLUMN "expenses"."amount" IS 'must be positive';

COMMENT ON COLUMN "budgets"."amount" IS 'must be positive';

ALTER TABLE "expenses" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "budgets" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "budgets" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "wallets" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

