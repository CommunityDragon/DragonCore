FROM golang:1.14

RUN go get github.com/volatiletech/sqlboiler
RUN go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql
RUN mkdir /sqlboiler
WORKDIR /sqlboiler

COPY sqlb.sh /usr/local/bin/sqlb
RUN chmod +x /usr/local/bin/sqlb

ENTRYPOINT ["sqlb"]
CMD ["--wipe", "--add-panic-variants", "psql"]