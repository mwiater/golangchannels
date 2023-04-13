#!/bin/bash
#
# From root of project, run: `bash scripts/golang_pprof.sh`

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

NUMPROCS=$(grep ^cpu\\scores /proc/cpuinfo | uniq |  awk '{print $4}')

echo -e "${CYANBOLD}Running pprofs for ${NUMPROCS} CPUs...${RESET}"

for (( workers=1; workers<=NUMPROCS; workers++ ))
do 
  echo -e "${CYANBOLD}Running pprof: ${workers} Workers${RESET}"
  go clean -testcache && go test ./dispatcher/dispatcher_test.go -bench=. -benchmem -run=^# -cpuprofile ./pprof/cpuprofile-0${workers}-workers.out -memprofile ./pprof/memprofile-0${workers}-workers.out -blockprofile ./pprof/blockprofile-0${workers}-workers.out -mutexprofile ./pprof/mutexprofile-0${workers}-workers.out -args -startingWorkerCount=${workers} -maxWorkerCount=${workers} -jobCount=64
  echo -e "    ${GREENBOLD}Generated Profile: ./pprof/cpuprofile-0${workers}-workers.out${RESET}"
  echo -e "    ${GREENBOLD}Generated Profile: ./pprof/memprofile-0${workers}-workers.out${RESET}"
  echo -e "    ${GREENBOLD}Generated Profile: ./pprof/blockprofile-0${workers}-workers.out${RESET}"
  echo -e "    ${GREENBOLD}Generated Profile: ./pprof/mutexprofile-0${workers}-workers.out${RESET}\n"
done

status=$?
if test $status -ne 0
then
  echo -e "${REDBOLD}...Error: 'go test' command failed!${RESET}"
  echo ""
  exit 1
fi
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""