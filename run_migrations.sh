#!/bin/bash
# Wait for the PostgreSQL server to start
sleep 10
echo "Running migrations..."
for i in /migrations/*.up.sql
do
  echo "Running $i"
  PGPASSWORD=Secret psql -h the_monkeys_db -U root -d the_monkeys_user_dev -f $i
done
echo "Migrations finished."
