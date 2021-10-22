#!/bin/bash
protoc -I . stream.proto --go_out=plugins=grpc:.