ALTER TABLE IF EXISTS "Product_UniqueView" DROP CONSTRAINT IF EXISTS "fk_ProductView";
ALTER TABLE IF EXISTS "Product_UniqueView" DROP CONSTRAINT IF EXISTS "fk_UserView";

DROP TABLE IF EXISTS "Product_UniqueView"
