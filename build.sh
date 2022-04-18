VERSION="v0.01"

POSITIONAL=()

if [ "$#" != 2 ] && [ "$1" != "-h" ]; then
    echo -e '\nSomething is missing... Type "./build.sh -h" without the quotes to find out more...\n'
    exit 0
fi

while [[ $# -gt 0 ]]
do
key="$1"

case $key in
    -o|--os)
    OS="$2"
    shift # past argument
    shift # past value
    ;;
    -h|--help)
    echo -e "\nbuild.sh $VERSION\n\nOPTIONS:\n\n-o: <l/m/w>\n    l = Linux   | ./build.sh -o l\n    m = Mac     | ./build.sh -o m\n    w = Windows | ./build.sh -o w\n\n-h, --help: Help\n\n"
    exit 0
    ;;
esac
done
set -- "${POSITIONAL[@]}" # restore positional parameters

if [ $OS == "l" ]; then
   wails build
fi

if [ $OS == "m" ]; then
   wails build -platform darwin/amd64
fi

if [ $OS == "w" ]; then
   wails build -platform windows/amd64
fi
