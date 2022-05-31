#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/../.." || exit 1

sudo docker pull iamsashko/web2kindle

sudo docker rm web2kindle

sudo docker run -d -p 80:80 --name web2kindle \
  -v "$scriptDir/../../tmp":/storage/web2kindle/tmp \
  -v "$scriptDir/../../config":/storage/web2kindle/config \
  -v "$scriptDir/../../logs":/storage/web2kindle/logs \
  iamsashko/web2kindle

cd "$currentDir" || exit 1

exit 0
