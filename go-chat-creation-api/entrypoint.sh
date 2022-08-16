#!/bin/bash

# /usr/bin/wait-for-it.sh redis:6379 -t 0
cd /go/src/github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/cmd/go-chat-creation-api

go build

/usr/bin/wait-for-it.sh chat-api:3000 -t 0
# Then exec the container's main process (what's set as CMD in the Dockerfile).
exec "$@"
