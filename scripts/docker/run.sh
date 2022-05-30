#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/../.." || exit 1

sudo docker pull iamsashko/web2kindle

sudo docker rm web2kindle

sudo docker run -d -p 80:80 --name web2kindle -v tmp:/storage/web2kindle/tmp \
  -v config:/storage/web2kindle/config \
  iamsashko/web2kindle

cd "$currentDir" || exit 1

exit 0
