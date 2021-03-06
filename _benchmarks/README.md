# Benchmarks

Higher is better.

## Results

- install bombardier: 

```bash
go get -u github.com/codesenberg/bombardier
```

### Static Path

```bash
bombardier -c 125 -n 1000000 http://localhost:3000
bombardier -c 200 -n 1000000 http://localhost:3000
```

### Parameterized (dynamic) Path

```bash
bombardier -c 125 -n 1000000 http://localhost:3000/user/42
bombardier -c 200 -n 1000000 http://localhost:3000/user/42
```

### Details

#### echo

> 2020.04.18

- static path

```text
Bombarding http://localhost:3000 with 1000000 request(s) using 125 connection(s)
 1000000 / 1000000 [================================================================================================================================================================] 100.00% 49599/s 20s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     49896.58    4647.15   91974.75
  Latency        2.51ms   682.91us    65.84ms
  HTTP codes:
    1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     8.89MB/s

```

### fasthttp

> 2020.04.18

- static path

```text
Bombarding http://localhost:3000 with 1000000 request(s) using 125 connection(s)
 1000000 / 1000000 [================================================================================================================================================================] 100.00% 62769/s 15s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     63391.52    5806.82   76139.45
  Latency        1.97ms   679.23us    57.51ms
  HTTP codes:
    1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    12.38MB/s

```

### gin

> 2020.04.18

- static path

```text
Bombarding http://localhost:3000 with 1000000 request(s) using 125 connection(s)
 1000000 / 1000000 [================================================================================================================================================================] 100.00% 51607/s 19s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     51712.17    4323.32   59599.19
  Latency        2.41ms   574.26us    49.16ms
  HTTP codes:
    1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     9.22MB/s

```

#### gorilla-mux

> 2020.04.18

- static path

```text
Bombarding http://localhost:3000 with 1000000 request(s) using 125 connection(s)
 1000000 / 1000000 [================================================================================================================================================================] 100.00% 49093/s 20s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     49596.82    4578.00   57997.43
  Latency        2.52ms   609.98us    49.80ms
  HTTP codes:
    1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     8.84MB/s

```

#### rux

> 2020.04.18

- static path

```text
Bombarding http://localhost:3000 with 1000000 request(s) using 125 connection(s)
 1000000 / 1000000 [================================================================================================================================================================] 100.00% 49088/s 20s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     49338.86    4102.08   58959.67
  Latency        2.53ms   532.33us    31.46ms
  HTTP codes:
    1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:     8.80MB/s

```