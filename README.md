# ratelimiter


An atomic counter wrapped in Take/Release API. Usage 
```Go
limiter := ratelimiter.New(100)    // Maximum 100 tokens
limiter.Take(10)                   // Take() will block excecution until 10 tokens are available
limiter.Release(10)    
```
What about https://codereview.stackexchange.com/questions/15954/concurrency-limit-map-in-go
