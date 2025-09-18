#!/bin/bash
# The following script sets up the Postgres DB.
# Redis instance.
# Brings up all the Go services like the authentication service,
# the DB service, and the driver service.
# Launches them in their respective Kubernetes pods.

# Create the directories needed for postgres
export DATAPATH=/usr/local/PGData/GoBackend
export LOGPATH=/var/local/log
sudo mkdir -p $DATAPATH
sudo mkdir -p $LOGPATH
sudo chown postgres $DATAPATH
sudo chown postgres $LOGPATH

sudo -u postgres ./setup_dbs.sh

sudo -u postgres ./start_dbs.sh

sudo ./create_tables.sh