CREATE TYPE users_status AS ENUM (
  '1',
  '2',
  '3'
);

CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "resetToken" varchar,
  "status" users_status DEFAULT '3',
  "address" varchar,
  "avatar" varchar,
  "phoneNumber" bigint,
  "role" bigint,
  "firstName" varchar,
  "lastName" varchar,
  "middleName" varchar,
  "city" bigint,
  "likeProducts" bigint,
  "viewedProducts" bigint,
  "addresses" jsonb,
  "resetTokenExpiration" timestamptz,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);
