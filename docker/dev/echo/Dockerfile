# hydra
#
# this image is used in order to run the echo
# server with hot reload
FROM golang:1.14

# installs air
RUN go get -u github.com/cosmtrek/air
WORKDIR $GOPATH/src/github.com/cosmtrek/air
RUN git checkout tags/v1.12.1
RUN make ci && make install

# set workdir to /app
WORKDIR /app

# default environmental variables
ENV DB_USER=postgres
ENV DB_PASS=postgres
ENV DB_HOST=db
ENV DB_NAME=app
ENV DB_SSL=disable

# run air
ENTRYPOINT ["/go/bin/air"]
CMD ["-c", ".air.conf"]