CREATE TYPE users_type AS ENUM (
  '1',
  '2',
  '3'
);

CREATE TYPE users_status AS ENUM (
  '0',
  '1'
);

CREATE TABLE "Users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "userType" users_type DEFAULT '3',
  "status" users_status DEFAULT '1',
  "address" varchar,
  "avatar" varchar,
  "image" varchar,
  "phoneNumber" bigint,
  "role" bigint DEFAULT 2,
  "firstName" varchar,
  "lastName" varchar,
  "middleName" varchar,
  "city" bigint,
  "likeProducts" bigint[],
  "viewedProducts" bigint[],
  "deviceToken" varchar[],
  "addresses" jsonb,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);
