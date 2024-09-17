CREATE TYPE product_status AS ENUM (
  '0',
  '1'
);

CREATE TABLE "Product" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "image" varchar NOT NULL,
    "countInStock" integer NOT NULL,
    "description" varchar NOT NULL,
    "sold" integer DEFAULT 0,
    "discount" integer,
    "discountStartDate" date,
    "discountEndDate" date,
    "type" integer NOT NULL,
    "status" product_status DEFAULT '1',
    "slug" varchar NOT NULL,
    "views" integer DEFAULT 0,
    "price" integer NOT NULL,
    "location" varchar NOT NULL
)