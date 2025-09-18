#!/bin/bash

# This script has to be run as postgres user

# This script creates the directories and initializes them for postgres.
export DATAPATH=/usr/local/PGData/GoBackend
export PATH=$PATH:/usr/lib/postgresql/16/bin/

if [ ! -d $DATAPATH/Authsvc ]; then
	initdb -D $DATAPATH/Authsvc
fi

if [ ! -d $DATAPATH/DatasvcShard1 ]; then
	initdb -D $DATAPATH/DatasvcShard1
fi

if [ ! -d $DATAPATH/DatasvcShard2 ]; then
	initdb -D $DATAPATH/DatasvcShard2
fi
