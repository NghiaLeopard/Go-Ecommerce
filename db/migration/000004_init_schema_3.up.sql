CREATE TABLE "Product" (
    "_id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "image" varchar NOT NULL,
    "countInStock" integer NOT NULL,
    "description" varchar NOT NULL,
    "sold" integer NOT NULL DEFAULT 0,
    "discount" integer NOT NULL DEFAULT 0,
    "discountStartDate" date NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "discountEndDate" date NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    "type" integer NOT NULL,
    "status" integer NOT NULL,
    "slug" varchar NOT NULL,
    "price" integer NOT NULL,
    "location" integer NOT NULL,
    "views" integer NOT NULL DEFAULT 0,
    "create_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT "fk_ProductCity" FOREIGN KEY("location") REFERENCES "City"("_id") ON DELETE CASCADE
)

