# Go Interfaces

Interfaces in Go.

Resources:
https://gobyexample.com/interfaces  
https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go


So, I'm not sure why, but when I do `civic.turnOn()` it does not update the engine field. I think I'm going to see if I can't pick up a book.

The answer was that I was passing by value, when I needed to actually change the instance! So passing in a pointer works as expected.

Car engine example I came up with after reading through Go by example.

Car & velocity/vehicle examples via: https://softchris.github.io/golang-book/