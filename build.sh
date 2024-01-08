sudo docker network create monkeys-network
sudo docker rm -f the_monkeys_db
sudo docker run --network=monkeys-network --name the_monkeys_db -e POSTGRES_PASSWORD=Secret -e POSTGRES_USER=root -e POSTGRES_DB=the_monkeys_user_dev -p 1234:5432 -d postgres;

# Write the migration script

sudo docker rm -f the_monkeys_dev
sudo docker build -t my-golang-app . && sudo docker run --network=monkeys-network -p 8081:8081 -it --name the_monkeys_dev my-golang-app
