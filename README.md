# ratelimiter


An atomic counter wrapped in Take/Release API. Usage 
    limiter := ratelimiter.New(100)  // Maximum 100 tokens
    limiter.Take(10)  // Take() will block excecution until there are 10 tokens available
    limiter.Release(10)    
    
What about https://codereview.stackexchange.com/questions/15954/concurrency-limit-map-in-go
