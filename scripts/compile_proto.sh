#!/bin/bash
set -e

cd proto
protoc --go_out=plugins=grpc:. *.proto 
