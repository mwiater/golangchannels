#!/bin/bash
#
# From root of project, run: `bash scripts/golang_godoc.sh`

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

echo -e "${CYANBOLD}Starting godoc server on port: 6060...${RESET}"
godoc -http=:6060
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""