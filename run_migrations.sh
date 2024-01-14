#!/bin/bash
for i in /scripts/*.up.sql;
do
  echo Running $i;
  PGPASSWORD=Secret psql -h the_monkeys_db -U root -d the_monkeys_user_dev -f $i;
done
