#!/bin/bash -e
# wait-for-postgres.sh
NEXT_WAIT_TIME=0
END_TIME=20
echo "wait for postgres timeout $END_TIME"

until docker-compose exec postgres pg_isready; do
  echo >&2 "$(date +%Y%m%dt%H%M%S) Postgres is unavailable - sleeping"
  echo "wait time $((NEXT_WAIT_TIME++))"
  echo

  if [ $NEXT_WAIT_TIME -eq $END_TIME ]; then 
    echo "Error: Timed out waiting for the postgres to start." >&2
    exit 1
  fi

  sleep 1
done

echo >&2 "$(date +%Y%m%dt%H%M%S) Postgres is up - executing command"