CREATE TABLE "Product_Type" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "slug" varchar NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT (now()),
    "update_at" date
);

ALTER TABLE "Product" ADD CONSTRAINT "fk_ProductType" FOREIGN KEY ("type") REFERENCES "Product_Type"("id") ON DELETE CASCADE;