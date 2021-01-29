#/bin/bash

TESTREDIS_DOCKER="testredis"
TESTREDIS_DOCKER_PORT=6378
GITROOT=$(git rev-parse --show-toplevel)


# start redis docker container
docker run --name $TESTREDIS_DOCKER -p $TESTREDIS_DOCKER_PORT:6379 -d redis
sleep 1

# test datastore
go test $GITROOT/pkg/datastore/*.go -count 1



# clean up
docker rm -f -v $TESTREDIS_DOCKER