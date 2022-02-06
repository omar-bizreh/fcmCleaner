#!/bin/bash
CGO_ENABLED=0 go build
zip fcmCleaning.zip fcmCleaner .env
mv fcmCleaning.zip /home/dev/Desktop/CleaningService/