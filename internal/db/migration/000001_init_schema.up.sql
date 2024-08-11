CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "resetToken" varchar,
  "resetTokenExpiration" timestamptz,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);
