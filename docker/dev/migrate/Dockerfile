FROM migrate/migrate:latest

COPY migrate.sh /dbmigrate
RUN chmod +x /dbmigrate

WORKDIR /migrations

ENTRYPOINT ["/dbmigrate"]