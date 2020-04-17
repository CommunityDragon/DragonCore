# vue
#
# this image is used in order to run nuxt
# or vue in production mode
FROM node:13

# these args need to be set for lerna to work
ARG project_dir

# creates a one line lerna project starter
RUN echo '#!/bin/sh\nlerna exec -- yarn start' >> /usr/local/bin/start && chmod +x /usr/local/bin/start

# checks if project_dir args is set
RUN bash -c 'if [[ -z "$project_dir" ]]; then exit 1; fi'

# installs lerna
RUN npm install -g lerna

# copies all the necessary files
WORKDIR /app
COPY $project_dir $project_dir
COPY package.json .
COPY lerna.json .
COPY yarn.lock .

# installs the packages and builds the project
RUN lerna bootstrap
RUN lerna exec -- yarn build;

# entrypoint
ENTRYPOINT [ "start" ]
