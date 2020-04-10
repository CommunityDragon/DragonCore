# hydra
#
# this image is used in order to initialize the database
# of hydra before using any hydra commands
FROM oryd/hydra:v1.3.2-alpine

# imports a start script
COPY start.sh /usr/local/bin/hydra-start

ENTRYPOINT ["hydra-start"]