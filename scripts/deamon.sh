#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir" || exit 1

/bin/bash run.sh -m loop

cd "$currentDir" || exit 1

exit 0
