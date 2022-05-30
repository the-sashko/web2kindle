#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/../src" || exit 1

go build -o ../bin/app

cd "../bin" || exit 1

chmod -x app
chmod 775 app

cd "$currentDir" || exit 1

exit 0
