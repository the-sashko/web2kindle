#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

helpFlag=0
modeFlag="default"

while getopts m:h flag
do
    case "${flag}" in
        m) modeFlag=${OPTARG};;
        h) helpFlag=1;;
        *) exit 1;;
    esac
done

if [ $helpFlag == 1 ]; then
  echo "Usage: run.sh [options]"
  echo "Options:"
  echo "    -m    set running mode: default, test or loop"
  echo "    -h    show this help message and exit"
  exit 1
fi

cd "$scriptDir" || exit 1

if [ ! -f ../bin/app ]; then
  echo "app binary not exists"
  exit 1;
fi

cd "$scriptDir/../bin" || exit 1

if [ ! -d ../tmp ]; then
  mkdir ../tmp
  chmod -R 775 ../tmp
  exit 1;
fi

if [ ! -d ../logs ]; then
  mkdir ../tmp
  chmod -R 775 ../tmp
  exit 1;
fi

if test -f ../tmp/document.md; then
  rm ../tmp/document.md
fi

./app --mode="$modeFlag"

cd "$currentDir" || exit 1

exit 0
