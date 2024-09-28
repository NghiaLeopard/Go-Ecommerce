

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
    "status"integer NOT NULL,
    "slug" varchar NOT NULL,
    "price" integer NOT NULL,
    "location" integer NOT NULL,
    "views" integer DEFAULT 0,
    "create_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT "fk_ProductCity" FOREIGN KEY("location") REFERENCES "City"(id) ON DELETE CASCADE
)

