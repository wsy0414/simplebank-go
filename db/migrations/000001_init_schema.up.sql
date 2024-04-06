CREATE TABLE "accounts" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "balance" decimal(50,8) NOT NULL DEFAULT 0,
  "currency" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "account_id" integer NOT NULL,
  "amount" decimal(50,8) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "from_account_id" integer NOT NULL,
  "to_account_id" integer NOT NULL,
  "amount" decimal(50,8) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "accounts" ("balance");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "accounts"."name" IS '帳號';

COMMENT ON COLUMN "accounts"."password" IS '密碼';

COMMENT ON COLUMN "entries"."amount" IS 'postive or negative';

COMMENT ON COLUMN "transfers"."amount" IS 'just postive';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
