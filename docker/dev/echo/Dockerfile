FROM golang:1.14

# installs air
RUN go get -u github.com/cosmtrek/air
WORKDIR $GOPATH/src/github.com/cosmtrek/air
RUN git checkout tags/v1.12.1
RUN make ci && make install

# set workdir to /app
WORKDIR /app

# run air
ENTRYPOINT ["/go/bin/air"]
CMD ["-c", ".air.conf"]