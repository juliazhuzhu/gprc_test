#!/bin/bash
protoc -I . hello.proto --go_out=plugins=grpc:.
