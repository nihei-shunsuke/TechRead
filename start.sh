#!/bin/bash -eu
go mod tidy
air
tail -f /dev/null