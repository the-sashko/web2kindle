#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/../../docker" || exit 1

sudo docker build -t iamsashko/web2kindle .

sudo docker login

sudo docker push iamsashko/web2kindle:latest

cd "$currentDir" || exit 1

exit 0
