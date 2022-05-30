#!/bin/bash

currentDir=$(pwd)
scriptDir="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null && pwd)"

cd "$scriptDir/../tmp" || exit 1

if [ ! -f document.md ]; then
  echo "Markdown document not exists"
  exit 1
fi

if test -f document.mobi; then
  rm document.mobi
fi

ebook-convert document.md document.mobi >/dev/null

if test -f document.md; then
  rm document.md
fi

if [ ! -f document.mobi ]; then
  echo "Conversion error"
  exit 1
fi

cd "$currentDir" || exit 1

exit 0
