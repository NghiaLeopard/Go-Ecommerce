CREATE TABLE "Product_liked" (
    product_id INT NOT NULL,
    user_id INT NOT NULL,
    like_date timestamptz NOT NULL DEFAULT (now()),
    PRIMARY KEY (product_id, user_id),
    CONSTRAINT fk_ProductLiked FOREIGN KEY(product_id)  REFERENCES "Product"("_id") ON DELETE CASCADE,
    CONSTRAINT fk_UserLiked FOREIGN KEY(user_id)  REFERENCES "Users"("_id") ON DELETE CASCADE
)