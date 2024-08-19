CREATE TABLE "Role" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "permission" varchar[]
);

ALTER TABLE "Users" ADD CONSTRAINT "Fk_UserRole" FOREIGN KEY ("role") REFERENCES "Role"("id")