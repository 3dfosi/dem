#!/bin/bash

VERSION="v0.1"

POSITIONAL=()

if [ "$#" != 2 ] && [ "$1" != "-h" ]; then
    echo -e '\nSomething is missing... Type "./dist.sh -h" without the quotes to find out more...\n'
    exit 0
fi

while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -v|--version)
    VERSION="$2"
    shift # past argument
    shift # past value
    ;;
    -h|--help)
    echo -e "\n3DF install.sh $VERSION\n\nOPTIONS:\n\n-v: Version Number\n-h, --help: Help\n\n"
    exit 0
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

DIR="3DF-DEM-v$VERSION-x64-MacOS"

cp -R ../../build/bin/3DF\ Development\ Environment\ Manager.app .
./make_dmg \
    -image dmgback.png \
    -file 144,144 3DF\ Development\ Environment\ Manager.app \
    -symlink 416,144 /Applications \
    -convert UDBZ \
    $DIR.dmg


