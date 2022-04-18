#!/bin/bash

VERSION="v0.01"

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
    echo -e "\ndist.sh $VERSION\n\nOPTIONS:\n\n-v: Version Number\n-h, --help: Help\n\n"
    exit 0
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

DIR="3DF-DEM-v$VERSION-x64-Linux"

mkdir $DIR

cp ../../build/bin/dem $DIR/
cp desktop-icon.png $DIR/
cp README.md $DIR/

echo -e "[Desktop Entry]\nVersion=$VERSION\nName=DEM\nComment=3DF Development Environment Manager\nGenericName=DEM\nExec=/opt/3df-dem/dem\nPath=/opt/3df-dem\nTerminal=false\nType=Application\nIcon=/opt/3df-dem/desktop-icon.png\nStartupNotify=true\nCategories=Utility;Application;\n" > $DIR/3df-dem.desktop

tar -czvf $DIR.tar.gz $DIR
