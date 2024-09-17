CREATE TABLE "Product_liked" (
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (product_id, user_id),
    CONSTRAINT fk_ProductLiked FOREIGN KEY(product_id)  REFERENCES "Product"(id),
    CONSTRAINT fk_UserLiked FOREIGN KEY(user_id)  REFERENCES "Users"(id)
)
