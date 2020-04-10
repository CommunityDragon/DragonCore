```
  ____                               ____               
 |  _ \ _ __ __ _  __ _  ___  _ __  / ___|___  _ __ ___ 
 | | | | '__/ _` |/ _` |/ _ \| '_ \| |   / _ \| '__/ _ \
 | |_| | | | (_| | (_| | (_) | | | | |__| (_) | | |  __/
 |____/|_|  \__,_|\__, |\___/|_| |_|\____\___/|_|  \___|
                  |___/                                 
```
DragonCore is a hub for CommunityDragon.
Everything related to CommunityDragon will be interlinked with DragonCore

&nbsp;

## Structure
```
ROOT
 ├─ .external/          # External repositories
 ├─ docker/             # All docker files for CI/CD and local development
 ├─ frontend/           # All front-end libraries and projects
 │   ├─ main/           # Main front-end using Nuxt
 │   ├─ auth/           # Authentication front-end using plain Vue
 │   └─ developer/      # Developer front-end using Nuxt
 │
 ├─ lib/                # The API library
 │   └─ models/         # The API models
 │
 ├─ migrations/         # All migrations for the database
 ├─ modules/            # The API is split up in modules, this is the directory for it
 ├─ ...
 ├─ lerna.json          # Lerna configuration for managing all front-end projects
 ├─ package.json        # Global NPM dependencies for all front-end projects
 ├─ go.mod              # Dependencies of the API
 └─ main.go             # Entrypoint of the API
```

&nbsp;

## Prerequisites
- [Go 1.14+](https://golang.org/) (optional)
- [Node 12+](https://nodejs.org/en/) (optional)
  - [Yarn 1.X](https://classic.yarnpkg.com/)
- [Docker](https://www.docker.com/)
  - [Compose](https://docs.docker.com/compose/install/)

Docker and Docker-Compose are used for running all the applications.
Node and Go are optional, as the only use is to install other dependencies
using a CLI. If you do want to install a dependency for node, do make sure 
to use Yarn.

&nbsp;

## Setup
### Development
We also use a proxy in development mode, so you'll need to set some /etc/hosts values
in order to be able to have access to all the projects. Copy and paste the following values:
```
127.0.0.1  developer.communitydragon.localhost
127.0.0.1    traefik.communitydragon.localhost
127.0.0.1       auth.communitydragon.localhost
127.0.0.1        www.communitydragon.localhost
127.0.0.1        api.communitydragon.localhost
127.0.0.1            communitydragon.localhost
```
Everything is set up in docker. All you need to do is start up the containers using docker-compose.
```bash
docker-compose up
```
Or run it as a daemon
```bash
docker-compose up -d
```

&nbsp;

### Production
**(not set up yet)**

We use docker-images for the production build. Just use the following command (keep 
in mind, a proxy is set up within docker. You'll need to change out the domains if 
you want to run production yourself)
```bash
docker-compose up -f docker-compose.build.yml
```
You can now spin up the production images with the production docker-compose configuration file.
```bash
docker-compose up -f docker-compose.prod.yml -d
```
