#!/bin/sh

# migrates the
hydra migrate sql $DSN --yes
hydra clients create \
    --id dragoncore \
    --scope openid,offline \
    --secret dragoncoresecret \
    --endpoint http://0.0.0.0:4445 \
    --response-types code,id_token \
    --grant-types refresh_token,authorization_code \
    --callbacks http://api.communitydragon.localhost/oauth2/callback &>/dev/null &
hydra serve all --dangerous-force-http
