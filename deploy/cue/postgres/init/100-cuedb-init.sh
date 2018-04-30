#!/bin/bash
set -o errexit
readonly REQUIRED_ENVVARS=(
  "POSTGRES_USER"
  "POSTGRES_PASSWORD"
  "POSTGRES_DB"
)

main() {
  check_envvars
}

check_envvars() {
  for envvar in ${REQUIRED_ENVVARS[@]}; do
    if [[ -x "${!envvar}" ]]; then
      echo "Error:
    Environment variable '$envvar' not set.
    Ensure you have the following environment variables set:

      ${REQUIRED_ENVVARS[@]}

Aborting."
      exit
    fi
  done
}

main "@"
