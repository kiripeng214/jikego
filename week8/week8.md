## 1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
### redis-benchmark -d 10 -t set,get -q
SET: 52056.22 requests per second, p50=0.335 msec \
GET: 54229.93 requests per second, p50=0.327 msec 
### redis-benchmark -d 20 -t set,get -q
SET: 50607.29 requests per second, p50=0.343 msec \
GET: 50075.11 requests per second, p50=0.343 msec 
### redis-benchmark -d 50 -t set,get -q
SET: 51334.70 requests per second, p50=0.343 msec \
GET: 51466.80 requests per second, p50=0.343 msec
### redis-benchmark -d 100 -t set,get -q
SET: 51466.80 requests per second, p50=0.343 msec \
GET: 51072.52 requests per second, p50=0.343 msec
### redis-benchmark -d 200 -t set,get -q
SET: 49261.09 requests per second, p50=0.359 msec \
GET: 50025.02 requests per second, p50=0.351 msec 
### redis-benchmark -d 1000 -t set,get -q
SET: 50813.01 requests per second, p50=0.351 msec \
GET: 50454.09 requests per second, p50=0.343 msec
### redis-benchmark -d 5000 -t set,get -q
SET: 46511.62 requests per second, p50=0.367 msec \
GET: 46533.27 requests per second, p50=0.367 msec
### 前面10到1k基本上都在每秒52左右，到5k的时候直线下降10%左右

## 2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。