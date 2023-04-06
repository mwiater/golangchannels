#!/bin/bash
#
# From root of project, run: `bash scripts/golang_run.sh`

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

echo -e "${CYANBOLD}Formatting *.go files...${RESET}"
for i in *.go **/*.go ; do
  gofmt -w "$i"
  status=$?
  if test $status -ne 0
  then
    echo -e "${REDBOLD}...Error: 'gofmt' command failed!${RESET}"
    echo ""
    exit 1
  fi
  echo "Formatted: $i"
done;
echo -e "${GREENBOLD}...Complete${RESET}"
echo ""

echo -e "${CYANBOLD}Building Go app:${RESET} ${GREENBOLD}go build -o bin/golangchannels .${RESET}"
go build -o bin/golangchannels .
status=$?
if test $status -ne 0
then
	echo -e "${REDBOLD}...Error: 'go build' command failed!${RESET}"
  echo ""
  exit 1
fi

echo -e "${CYANBOLD}Running Go app:${RESET} ${GREENBOLD}./bin/golangchannels${RESET}"
./bin/golangchannels