#!/usr/bin/env bash
# bash

echo "go clean testcache: `pwd`..."
echo "-------------------------------------------------------------------------------"
echo ""

go clean -testcache ./...

echo ""
echo "-------------------------------------------------------------------------------"
echo "go clean testcache: DONE "
echo ""