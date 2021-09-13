#!/bin/bash

cd test/rpcdemo/proto/

protoc --go_out=plugins=grpc:. *.proto

cd .. && cd ..
