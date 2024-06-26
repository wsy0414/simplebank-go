CREATE TABLE "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "birthdate" timestamp,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "balance" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer NOT NULL,
  "currency" varchar NOT NULL,
  "balance" decimal(50,8) NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" integer NOT NULL,
  "amount" decimal(50,8) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "activity" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "from_user_id" integer NOT NULL,
  "to_user_id" integer NOT NULL,
  "amount" decimal(50,8) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "balance" ("user_id");

CREATE INDEX ON "balance" ("balance");

CREATE INDEX ON "balance" ("currency");

CREATE INDEX ON "transfer" ("user_id");

CREATE INDEX ON "activity" ("from_user_id");

CREATE INDEX ON "activity" ("to_user_id");

CREATE INDEX ON "activity" ("from_user_id", "to_user_id");

COMMENT ON COLUMN "users"."name" IS '帳號';

COMMENT ON COLUMN "users"."password" IS '密碼';

COMMENT ON COLUMN "transfer"."amount" IS 'postive or negative';

COMMENT ON COLUMN "activity"."amount" IS 'just postive';

ALTER TABLE "balance" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "activity" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("id");

ALTER TABLE "activity" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("id");
