#!/bin/bash
psql -U postgres -d contracts -c "TRUNCATE TABLE schemas;"
exec "$@"
