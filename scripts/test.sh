#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

greenColor='\033[0;32m'
yellowColor='\033[0;33m'
redColor='\033[0;31m'
blueColor='\033[0;34m'
noColor='\033[0m'

cd "$scriptDir" || exit 1

clear

if test -f ../bin/app; then
  rm ../bin/app
fi

if test -f document.md; then
  rm document.md
fi

echo -en "${yellowColor}Building...${redColor}"

/bin/bash build.sh

echo -e "${yellowColor}OK${noColor}"

echo -e "${greenColor}Run application${blueColor}"

/bin/bash run.sh -m test

echo -e "${greenColor}OK${noColor}"

cd "$currentDir" || exit 1

exit 0
