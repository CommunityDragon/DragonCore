# pgadmin
#
# this image is used in order to run pgadmin
# in desktop mode
FROM dpage/pgadmin4:4

ENV PGADMIN_CONFIG_SERVER_MODE=False
ENV PGADMIN_DEFAULT_EMAIL=admin@dragoncore
ENV PGADMIN_DEFAULT_PASSWORD=admin

COPY config.py .
COPY servers.json .

COPY pgpass.conf /
USER root
RUN chown pgadmin:pgadmin /pgpass.conf
USER pgadmin
RUN chmod 600 /pgpass.conf