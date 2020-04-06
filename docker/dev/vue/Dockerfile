FROM node:13

# installs lerna
RUN npm install -g lerna

# create a start script
COPY start.sh /usr/local/bin/lerna-start
RUN chmod +x /usr/local/bin/lerna-start

# set /app as work directory
WORKDIR /app

# run lerna-start
ENTRYPOINT ["lerna-start"]