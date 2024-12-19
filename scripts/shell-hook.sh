#!/usr/bin/env bash

printf ">> Tooling versions\n\n"

echo "-- Golang"
go version
echo

echo "-- Mage"
mage --version
echo

echo "-- Mockery"
mockery --version
echo

echo "-- goimports-reviser"
goimports-reviser --version
echo

echo "-- Revive"
revive --version
echo

echo "-- Commitizen"
echo "cz: $(cz version)"
echo

echo ">> Installing hooks"
husky install