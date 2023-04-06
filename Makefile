SHELL=/bin/bash

.DEFAULT_GOAL := list

# Color Console Output
RESET=\033[0m
REDBOLD=\033[1;31m
GREENBOLD=\033[1;32m
YELLOWBOLD=\033[1;33m
CYANBOLD=\033[1;36m

list:
	@echo ""
	@echo -e "${GREENBOLD}Targets in this Makefile:${RESET}"
	@echo ""
	@LC_ALL=C $(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]" && $$1 !~ "^[list.]" && $$1 !~ "^[always.]") {print "make "$$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
	@echo ""
	@echo "For details on these commands, see the bash scripts in the 'scripts/' directory."
	@echo ""

golang-run:
	scripts/golang_run.sh

golang-build:
	scripts/golang_build.sh

golang-test:
	scripts/golang_test.sh

golang-lint:
	scripts/golang_lint.sh