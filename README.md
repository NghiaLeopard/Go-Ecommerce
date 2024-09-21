# Go-Ecommerce

cd /mnt/d/github.com/NghiaLeopard/Go-ecommerce-Backend

Run on background

docker run --name rdb -d -p 6379:6379 redis

docker exec -it rdb redis-cli

docker run --rm --name redis-commander -d -p 8081:8081 rediscommander/redis-commander:latest
