#/bin/sh

MONGO_NAME="testmongo"
GITROOT=$(git rev-parse --show-toplevel)

function create_mongo(){
    docker run --name $MONGO_NAME -d -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo
}

function remove_mongo(){
    docker rm -f $MONGO_NAME -v
}

function e2e(){
    docker container inspect $MONGO_NAME > /tmp/e2e
    if [ "$(echo $?)" -eq 0 ];then
        remove_mongo
    fi
    create_mongo
    sleep 1
    # datastore test
    go test -count 1 $GITROOT/pkg/datastore/*.go
}

case $1 in
    "create_mongo")
        create_mongo
    ;;
    "e2e")
        e2e
    ;;

    *)
        echo "Default"
    ;;
esac