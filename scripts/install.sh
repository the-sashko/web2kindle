#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/.." || exit 1

mkdir logs
chmod -R 755 logs

mkdir bin
chmod -R 755 bin

mkdir tmp
chmod -R 755 tmp

touch tmp/telegram_last_update_id.txt
chmod -R 755 tmp/telegram_last_update_id.txt

cp config/config_sample.json config/config.json
chmod 755 config/config.json

cp config/credentials_sample.json config/credentials.json
chmod 755 config/credentials.json

cd "$scriptDir" || exit 1

/bin/bash build.sh

cd "$currentDir" || exit 1

exit 0
