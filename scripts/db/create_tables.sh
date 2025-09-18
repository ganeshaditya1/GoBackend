#!/bin/bash

psql -h localhost -U postgres -d postgres -p 5433 -c "CREATE DATABASE users"
psql -h localhost -U postgres -d postgres -p 5433 -f "../../schemas/users_table.sql"

psql -h localhost -U postgres -d postgres -p 5434 -c "CREATE DATABASE data_shard"
psql -h localhost -U postgres -d postgres -p 5435 -c "CREATE DATABASE data_shard"
psql -h localhost -U postgres -d postgres -p 5434 -f "../../schemas/data_table.sql"
psql -h localhost -U postgres -d postgres -p 5435 -f "../../schemas/data_table.sql"

