#!/bin/bash

# Run this script as postgres user. It terminates all three instances of the DB.

export DATAPATH=/usr/local/PGData/GoBackend
export PATH=$PATH:/usr/lib/postgresql/16/bin/

pg_ctl stop -D $DATAPATH/Authsvc
pg_ctl stop -D $DATAPATH/DatasvcShard1
pg_ctl stop -D $DATAPATH/DatasvcShard2
