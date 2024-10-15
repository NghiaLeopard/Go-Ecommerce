CREATE TABLE "Delivery_Type" (
    "_id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "price" integer NOT NULL,
    "createAt" timestamptz NOT NULL DEFAULT (now()),
    "update_at"timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);
