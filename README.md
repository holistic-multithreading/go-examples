# run benchmarks

```bash
go test -bench=.
```

## benchmark results

MacOS Ventura Intel i9 2.9 GHz, 6 physical cores (12 logical) 32Gb memory

### MacOS i9

| range | 1 thread | 12 threads |
| ----- | ------------- | --------- |
| thousand | 257 ns     | 5.3 µs    |
| million  | 239.5 µs   | 380.8 µs  |
| billion  | 238.8 ms   | 221.0 ms  |

### MacOS i9 - Docker

```bash
docker run -it -v "$PWD":/code golang sh
cd /code
go test -bench=.
```

| range | 1 thread | 12 threads |
| ----- | --------------- | --------- |
| thousand | 265.9 ns     | 5.9 µs    |
| million  | 514.7 µs     | 616.8 µs  |
| billion  | 501.3 ms     | 275.4 ms  |
