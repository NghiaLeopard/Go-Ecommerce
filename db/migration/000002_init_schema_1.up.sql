CREATE TABLE "Role" (
    "_id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "permission" varchar[],
    "create_at" timestamptz NOT NULL DEFAULT (now()),
    "update_at"timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

ALTER TABLE "Users" ADD CONSTRAINT "fk_UserRole" FOREIGN KEY ("role") REFERENCES "Role"("_id") ON DELETE CASCADE;


