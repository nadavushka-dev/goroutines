# Burger shop

A little burger shop that can make orders in a asynchronic way

### Why?
- learn go routins 

### Explenation:
We have 3 functionalities out burger shop.
- makeBuger
- makeChips
- makeDrink

Each function takes some time (differ between functions).

We want to make the order simultaniuosly (in an async way) 
in order to wait the least amount of time (i.e finish the order with the longest time taking function)

To do that we utilize the sync.WaitGroup and go keyword (which allows us to run async with goroutines)
