go test -v golang.org/x/net/http2 -run "TestTransport.*" >&1 | tee bench-trans.out


go test -v golang.org/x/net/http2 -run ".*Transport.*" -bench ".*Transport.*" >&1 | tee bench-trans.out


go test -v golang.org/x/net/http2 -run="^$" -bench=".*Transport.*" -count=10 >&1 | tee bench-trans2.out

go test -v golang.org/x/net/http2 -run="^$" -cpu=10000,20000,30000,40000,50000,60000,70000,80000,90000,100000 \
   -bench=".*Transport.*Parallel.*" -count=3 >&1 | tee bench-trans.out
