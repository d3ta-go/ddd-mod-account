#!/usr/bin/env bash
# bash

echo "go test: github.com/d3ta-go/ddd-mod-account/modules/account/infrastructure/migration -run ^TestRDBMSMigration_Run$... "
echo "-------------------------------------------------------------------------------"
echo ""

go test -timeout 120s github.com/d3ta-go/ddd-mod-account/modules/account/infrastructure/migration -run ^TestRDBMSMigration_Run$ -v -cover

echo ""
echo "-------------------------------------------------------------------------------"
echo "go test: DONE "
echo ""