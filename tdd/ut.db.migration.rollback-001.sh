#!/usr/bin/env bash
# bash

echo "go test: github.com/d3ta-go/ddd-mod-account/modules/account/infrastructure/migration -run ^TestRDBMSMigration_RollBack$... "
echo "-------------------------------------------------------------------------------"
echo ""

go test -timeout 120s github.com/d3ta-go/ddd-mod-account/modules/account/infrastructure/migration -run ^TestRDBMSMigration_RollBack$ -v -cover

echo ""
echo "-------------------------------------------------------------------------------"
echo "go test: DONE "
echo ""