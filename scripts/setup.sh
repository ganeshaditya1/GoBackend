#!/bin/bash
# The following script sets up the Postgres DB.
# Redis instance.
# Brings up all the Go services like the authentication service,
# the DB service, and the driver service.
# Launches them in their respective Kubernetes pods.

# Create the directories needed for postgres
export DATAPATH=/usr/local/PGData/GoBackend
export LOGPATH=/var/local/log
mkdir -p $DATAPATH
mkdir -p $LOGPATH
chown postgres $DATAPATH
chown postgres $LOGPATH

sudo -u postgres ./db/setup_dbs.sh

sudo -u postgres ./db/start_dbs.sh

# This script expects 4 environment variables.
# 1. AUTHSVC_PWORD
# 2. AUTHSVC_ADMIN_PWORD
# 3. DATASVC_SHARD1_PWORD
# 4. DATASVC_SHARD2_PWORD
./db/create_users.sh

./db/create_tables.sh
