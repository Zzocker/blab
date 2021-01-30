#/bin/bash

TESTREDIS_DOCKER="testredis"
TESTREDIS_DOCKER_PORT=6378

TESTMONGO_DOCKER="testmongo"
TESTMONGO_DOCKEER_PORT=27018
TESTMONGO_USERNAME="root"
TESTMONGO_PASSWORD="password"
GITROOT=$(git rev-parse --show-toplevel)


# start redis docker container
docker run --name $TESTREDIS_DOCKER -p $TESTREDIS_DOCKER_PORT:6379 -d redis

# start mongo docker container
docker run --name $TESTMONGO_DOCKER -p $TESTMONGO_DOCKEER_PORT:27017 -d -e MONGO_INITDB_ROOT_USERNAME=$TESTMONGO_USERNAME -e MONGO_INITDB_ROOT_PASSWORD=$TESTMONGO_PASSWORD mongo
sleep 1

# test datastore
go test $GITROOT/pkg/datastore/*.go -count 1 -v



# clean up
docker rm -f -v $TESTREDIS_DOCKER $TESTMONGO_DOCKER
