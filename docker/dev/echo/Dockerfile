FROM golang:1.14

# install air
WORKDIR /usr/local/bin
RUN curl -fLo air https://git.io/linux_air
RUN chmod +x air

# set workdir to /app
WORKDIR /app

# run air
CMD ["air"]
