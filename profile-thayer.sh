#!/bin/bash

# Run tests with high concurrency
out=bench-trans.out
go test -v golang.org/x/net/http2 -run="^$" -cpu=10000,20000,30000,40000,50000,60000,70000,80000,90000,100000 \
   -bench=".*Transport.*Parallel.*" -count=10 >&1 | tee $out
benchstat $out
