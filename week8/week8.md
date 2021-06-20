## 1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
### redis-benchmark -d 10 -t set,get -q
SET: 52056.22 requests per second, p50=0.335 msec \
GET: 54229.93 requests per second, p50=0.327 msec 
### redis-benchmark -d 20 -t set,get -q
SET: 50607.29 requests per second, p50=0.343 msec \
GET: 50075.11 requests per second, p50=0.343 msec 
### redis-benchmark -d 50 -t set,get -q
SET: 51334.70 requests per second, p50=0.343 msec
GET: 51466.80 requests per second, p50=0.343 msec