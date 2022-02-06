#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
NEWLINE="------------------------------------------"
now=`date`

echo "$now"
read -p "Enter release tag: " tag

imageName="localhost:32000/fcm-cleaner:$tag"

echo "$NEWLINE"
echo -e "${GREEN}Publishing FCM Cleaner API${NC}"
echo -e "${RED}Destination: Local Kubernetes${NC}"
echo "$NEWLINE"
echo "Building API"
echo "$NEWLINE"
go build
echo "$NEWLINE"
if [ $? -eq 0 ]; then
    echo -e "${GREEN} API successfully built${NC}"
else
    echo -e "${RED}Failed to build. Aborting${NC}"
    exit 999
fi

echo "$NEWLINE"
echo -e "${GREEN}Building Docker Image${NC}"
docker build -t $imageName .

if [ $? -eq 0 ]; then
    echo -e "${GREEN} Docker Image successfully built${NC}"
else
    echo -e"${RED}Failed to build. Aborting${NC}"
    exit 998
fi


echo "$NEWLINE"
echo -e "${GREEN}Pushing image to registery${NC}"
docker image push $imageName

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Docker image published to registery"
else
    docker start registry
    docker image push $imageName
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Docker image published to registery"
    else
        echo -e "${RED}Failed to publish image"
        exit 987
    fi
fi