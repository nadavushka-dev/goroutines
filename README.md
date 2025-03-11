# Burger shop

A little burger shop that can make orders in a asynchronic way

### Why?
- learn go routins and cocurrency

### Explenation:
Our burger shop can do only 3 things (but we do it well(-done?))
- make a burger
- make chips
- make a drink

Each function takes some time (differ between functions).

We want to make the order simultaniuosly (in an async way) 
in order to wait the least amount of time (i.e finish the order with the longest time taking function)

To do that we utilize the sync.WaitGroup and go keyword (which allows us to run async with goroutines)
