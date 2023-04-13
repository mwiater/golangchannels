#!/bin/bash
#
# From root of project, run: `bash scripts/golang_benchmark.sh`

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

echo -e "${CYANBOLD}Clearing test cache...${RESET}"
go clean -testcache
status=$?
if test $status -ne 0
then
  echo -e "${REDBOLD}...Error: 'go clean' command failed!${RESET}"
  echo ""
  exit 1
fi
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Running benchmarks...${RESET}"
go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -args -startingWorkerCount=8 -maxWorkerCount=8 -jobCount=16
status=$?
if test $status -ne 0
then
  echo -e "${REDBOLD}...Error: 'go test' command failed!${RESET}"
  echo ""
  exit 1
fi
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""