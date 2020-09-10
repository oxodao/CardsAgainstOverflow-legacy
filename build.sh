#!/bin/bash

echo "Building CardsAgainstOverflow executable"
echo "\tIf you want the docker version, just use the dockerfile"
echo "Building frontend"
rm -rf $PWD/frontend/dist/ $PWD/data
rm cardsagainstoverflow

yarn --cwd $PWD/frontend
yarn --cwd $PWD/frontend build

echo "Generating assets"
mkdir $PWD/data
mv $PWD/frontend/dist/* $PWD/data/
~/go/bin/pkger

echo "Building backend"
export CGO_ENABLED=1
go build

echo "Building docker image"
docker build -t cardsagainstoverflow .

echo "Removing old files"
rm -rf $PWD/frontend/dist/ $PWD/data
rm cardsagainstoverflow
