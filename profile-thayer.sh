#!/bin/bash

# Run tests with high concurrency
out=bench-trans.out
thread_list=10000,20000,30000,40000,50000,60000,70000,80000,90000,100000
go test -v ./http2 -run="^$" -cpu=$thread_list \
   -bench=".*Transport.*Parallel.*" -count=10 >&1 | tee ${out}-lock
benchstat $out

go test -tags smart_rwlock -v ./http2 -run="^$" -cpu=$thread_list \
   -bench=".*Transport.*Parallel.*" -count=10 >&1 | tee ${out}-rwlock
benchstat $out

benchstat ${out}-*
