#!/usr/bin/env bash

set -eE -o pipefail

# Get environment variables, depending on shell should be added already
. ./.envrc

export DB_URL=postgres://localhost/world_test?sslmode=disable

ginkgo -r -cover ./world