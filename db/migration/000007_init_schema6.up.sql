CREATE TABLE "Product_UniqueView" (
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    view_date date NOT NULL DEFAULT (now()),
    PRIMARY KEY (product_id, user_id),
    CONSTRAINT fk_ProductView FOREIGN KEY(product_id)  REFERENCES "Product"(id) ON DELETE CASCADE,
    CONSTRAINT fk_UserView FOREIGN KEY(user_id)  REFERENCES "Users"(id) ON DELETE CASCADE
)