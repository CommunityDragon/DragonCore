# genmig
#
# this image is used in order to create a migration
# using a temporary database
FROM alpine:3.11

# updates package listing
RUN apk update

# installs packages
RUN apk add \
        git \
        libpq \
        build-base \
        python3-dev \
        postgresql-dev \
        postgresql-client

# installs pip
RUN python3 -m ensurepip
RUN rm -r /usr/lib/python*/ensurepip

# installs pip packages
RUN pip3 install --no-cache --upgrade \
        pip \
        wheel \
        PyYAML \
        psycopg2 \
        pgdbconn \
        setuptools

# symlinks python and pip
RUN if [ ! -e /usr/bin/python ]; then ln -sf python3 /usr/bin/python; fi
RUN if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi

# installing pyrseas
WORKDIR /tmp/pyrseas
RUN git clone git://github.com/perseas/Pyrseas.git /tmp/pyrseas
RUN python setup.py install
WORKDIR /
RUN rm -rf /tmp/pyrseas

# imports the genmig bash script
COPY genmig.sh /usr/local/bin/genmig
RUN chmod +x /usr/local/bin/genmig

# default environmental variables
ENV DB_USER=postgres
ENV DB_PASS=postgres
ENV DB_HOST=db

ENTRYPOINT ["genmig"]
