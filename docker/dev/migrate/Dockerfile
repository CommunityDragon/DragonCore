# migrate
#
# this image is used in order to execute go-migrate
# commands without having to specify a database location
FROM migrate/migrate:latest

# imports the migrate bash script
COPY migrate.sh /dbmigrate
RUN chmod +x /dbmigrate

# sets migrations director as work directory
WORKDIR /migrations

# default environmental variables
ENV DB_USER=postgres
ENV DB_PASS=postgres
ENV DB_HOST=db
ENV DB_NAME=app
ENV DB_SSL=disable

ENTRYPOINT ["/dbmigrate"]