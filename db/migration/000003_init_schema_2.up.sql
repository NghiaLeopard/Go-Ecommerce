CREATE TABLE "City" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT (now()),
    "update_at"timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

ALTER TABLE "Users" ADD CONSTRAINT "fk_UserCity" FOREIGN KEY ("city") REFERENCES "City"("id") ON DELETE CASCADE;