#!/bin/bash

psql -h localhost -U postgres -p 5433 -c "ALTER USER AUTHSVC_ADMIN createdb"
createdb -h localhost -U authsvc_admin -p 5433  users
psql -h localhost -U postgres -p 5433 -c "ALTER USER AUTHSVC_ADMIN nocreatedb"

psql -h localhost -U authsvc_admin -d users -p 5433 -f "../schemas/users_table.sql"

psql -h localhost -U authsvc_admin -p 5433 -d users -c "GRANT SELECT, INSERT, UPDATE, DELETE on users to authsvc"
psql -h localhost -U authsvc_admin -p 5433 -d users -c "REVOKE ALL PRIVILEGES on admin from authsvc"
psql -h localhost -U authsvc_admin -p 5433 -d users -c "GRANT SELECT on admin to authsvc"
psql -h localhost -U authsvc_admin -p 5433 -d users -c "GRANT USAGE, SELECT ON SEQUENCE users_userid_seq TO authsvc"

psql -h localhost -U postgres -p 5434 -c "ALTER USER datasvc_shard1 createdb"
createdb -h localhost -U datasvc_shard1 -p 5434  data_shard
psql -h localhost -U postgres -p 5434 -c "ALTER USER datasvc_shard1 nocreatedb"

psql -h localhost -U postgres -p 5435 -c "ALTER USER datasvc_shard2 createdb"
createdb -h localhost -U datasvc_shard2 -p 5435  data_shard
psql -h localhost -U postgres -p 5435 -c "ALTER USER datasvc_shard2 nocreatedb"

psql -h localhost -U datasvc_shard1 -d data_shard -p 5434 -f "../schemas/data_table.sql"
psql -h localhost -U datasvc_shard2 -d data_shard -p 5435 -f "../schemas/data_table2.sql"

