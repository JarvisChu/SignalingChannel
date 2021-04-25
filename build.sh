#!/bin/sh

tag="v0.0.2"
docker build -t signalingchannel:$tag . || exit
docker tag signalingchannel:$tag jarvischu/signalingchannel:$tag
docker push jarvischu/signalingchannel:$tag