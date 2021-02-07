#!/bin/sh

tag="v0.0.2"
docker build -t signalchannel:$tag . || exit
docker tag signalchannel:$tag jarvischu/signalchannel:$tag
docker push jarvischu/signalchannel:$tag