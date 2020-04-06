#!/bin/sh

if [[ -z "$1" ]]; then
    echo "no migration name supplied, not generating a migration";
    exit 0;
fi

NAME=$1
CURRENTTIME=`date +"%Y%m%d%H%M%S"`

# returns an error
function handle_error() {
    if [[ $? -ne 0 ]]; then
        echo -e "\e[31mERROR: $1";
        exit 1;
    fi
}

# connects to the default database
function psql_glob() {
    psql -v ON_ERROR_STOP=1 -q -U "$PGUSERNAME" -h db $@
}

# connects to the migrate database
function psql_migrate() {
    psql_glob -d migrate $@
}

# drops the migrate database
function drop_database() {
    psql_glob <<-EOSQL
        DROP DATABASE IF EXISTS migrate;
EOSQL
}

# creates the migrate database
function create_database() {
    psql_glob <<-EOSQL
        CREATE DATABASE migrate;
EOSQL
    psql_migrate <<-EOSQL
        REVOKE ALL ON SCHEMA public FROM postgres;
        REVOKE ALL ON SCHEMA public FROM PUBLIC;
EOSQL
}

# recreates the migrate database
function recreate_database() {
    echo "=== migrate database recreation ==="
    echo "recreating migrate database..."
    drop_database 2> /dev/null;
    handle_error "failed to drop migrate database, make sure nothing is connected to the database"
    create_database 2> /dev/null;
    handle_error "failed to create migrate database"
    echo "recreated migrate database"
    echo ""
}

# migrates all migrations up
function migrate_up() {
    echo "=== running migrations up ==="
    find /migrations -type f -name "*.up.sql" -print0 | sort | while read filename; do
        echo "running $filename";
        failed_file=""
        psql_migrate -f $filename 2> /dev/null;
        handle_error "migration $filename failed to execute"
        echo "completed $filename"
    done
    handle_error "unable to migrate up"
    echo "migrations completed"
    echo ""
}

# local yamltodb
function yamltodb_l() {
    yamltodb -H db -U $PGUSERNAME $@
}

# generates a migration
function gen_mig() {
    yamltodb_l -m migrate $@
}

# generates an up migration
function gen_mig_up() {
    echo "generating an up migration..."
    gen_mig -o "/migrations/${CURRENTTIME}_${NAME}.up.sql";
    if [[ $? -ne 0 ]]; then
        echo -e "\e[31mERROR: unable to migrate up migration";
        rm -rf "/migrations/${CURRENTTIME}_${NAME}.*.sql"
        exit 1;
    fi
    echo "generated up migration"
}

# generates a down migration
function gen_mig_down() {
    echo "generating a down migration..."
    gen_mig --revert -o "/migrations/${CURRENTTIME}_${NAME}.down.sql";
    if [[ $? -ne 0 ]]; then
        echo -e "\e[31mERROR: unable to migrate down migration";
        rm -rf "/migrations/${CURRENTTIME}_${NAME}.*.sql"
        exit 1;
    fi
    echo "generated down migration"
}

# removes migration if empty
function remove_mig_if_empty() {
    if [ ! -s /migrations/${CURRENTTIME}_${NAME}.up.sql ]; then
        echo "no changes in migration, so removing migration..."
        rm /migrations/${CURRENTTIME}_${NAME}.up.sql
        rm /migrations/${CURRENTTIME}_${NAME}.down.sql
        echo "new migration removed"
    fi
}

recreate_database
migrate_up
echo "=== migration generation ==="
gen_mig_up
gen_mig_down
remove_mig_if_empty
echo ""
echo "=== dropping migrate database ==="
echo "dropping migrate database..."
drop_database
handle_error "failed to drop migrate database"
echo "migrate database dropped"

chmod -R 777 /migrations
chmod -R 777 /metadata