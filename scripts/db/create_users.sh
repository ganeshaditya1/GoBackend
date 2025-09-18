#!/bin/bash

if [ "x${AUTHSVC_PWORD}" == "x" ]; then
    echo "Password for the Authsvc user in postgres not specified."
    exit 1
fi

if [ "x${AUTHSVC_ADMIN_PWORD}" == "x" ]; then
    echo "Password for the Authsvc Admin in postgres not specified."
    exit 1
fi

if [ "x${DATASVC_SHARD1_PWORD}" == "x" ]; then
    echo "Password for the Datasvc shard1 user in postgres not specified."
    exit 1
fi

if [ "x${DATASVC_SHARD2_PWORD}" == "x" ]; then
    echo "Password for the Datasvc shard2 user in postgres not specified."
    exit 1
fi

psql -h localhost -U postgres -d postgres -p 5433 -c "CREATE USER authsvc WITH PASSWORD '${AUTHSVC_PWORD}'"
psql -h localhost -U postgres -d postgres -p 5433 -c "CREATE USER authsvc_admin WITH PASSWORD '${AUTHSVC_ADMIN_PWORD}'"

psql -h localhost -U postgres -d postgres -p 5434 -c "CREATE USER datasvc_shard1 WITH PASSWORD '${DATASVC_SHARD1_PWORD}'"
psql -h localhost -U postgres -d postgres -p 5435 -c "CREATE USER datasvc_shard2 WITH PASSWORD '${DATASVC_SHARD2_PWORD}'"

psql -h localhost -U postgres -d postgres -p 5433 -c "GRANT SELECT, INSERT, UPDATE, DELETE on users to authsvc"
psql -h localhost -U postgres -d postgres -p 5433 -c "GRANT SELECT, INSERT, UPDATE, DELETE on admin to authsvc_admin"
psql -h localhost -U postgres -d postgres -p 5433 -c "REVOKE ALL PRIVILEGES on admin from authsvc"
psql -h localhost -U postgres -d postgres -p 5433 -c "GRANT SELECT on admin to authsvc"