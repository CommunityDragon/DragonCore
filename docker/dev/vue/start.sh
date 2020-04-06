#!/bin/sh
lerna bootstrap --use-workspaces;
lerna exec $@ -- yarn dev;