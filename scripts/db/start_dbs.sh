#!/bin/bash

# You have to run this script as postgres user

export PATH=$PATH:/usr/lib/postgresql/16/bin/
export DATAPATH=/usr/local/PGData/GoBackend
export LOGPATH=/var/local/log

# Launch three instances of the postgres db one for Auth Service. One for each Datasvc shards.
# The config file was edited to remove the config external_pid_file to prevent creating a pid file. 
# This is needed to be able to run 3 instances of the Postgres db on the same host.

pg_ctl -D $DATAPATH/Authsvc -o "-p 5433" -l $LOGPATH/Authsvc.log start
pg_ctl -D $DATAPATH/DatasvcShard1 -o "-p 5434" -l $LOGPATH/DatasvcShard1.log start
pg_ctl -D $DATAPATH/DatasvcShard2 -o "-p 5435" -l $LOGPATH/DatasvcShard2.log start
