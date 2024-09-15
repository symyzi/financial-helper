CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "currency" varchar NOT NULL
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "description" varchar,
  "category_id" bigint NOT NULL,
  "transaction_date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "budget" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "transactions" ("user_id");

CREATE INDEX ON "transactions" ("category_id");

CREATE INDEX ON "transactions" ("transaction_date");

CREATE INDEX ON "transactions" ("user_id", "transaction_date");

CREATE INDEX ON "budget" ("user_id");

CREATE INDEX ON "budget" ("category_id");

CREATE INDEX ON "budget" ("user_id", "category_id");

COMMENT ON COLUMN "transactions"."amount" IS 'must be positive';

COMMENT ON COLUMN "budget"."amount" IS 'must be positive';

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "budget" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "budget" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

