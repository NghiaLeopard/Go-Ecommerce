CREATE TABLE "Role" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "permission" varchar[]
);

ALTER TABLE "Users" ADD CONSTRAINT "Fk_UserRole" FOREIGN KEY ("role") REFERENCES "Role"("id");


CREATE TABLE "City" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL
);

ALTER TABLE "Users" ADD CONSTRAINT "Fk_UserCity" FOREIGN KEY ("city") REFERENCES "City"("id");