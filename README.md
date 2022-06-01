# Welcome to Tom's Pizza Delivery Service
Delivering pizzas to anyone in our path since 2022
## Rules
The neightborhood we deliver to lies within a massive 2^64 x 2^64 grid with the first delivery always starting in the very center<br />
The deliverer moves always exactly one house to the north (`^`), south (`v`), east (`>`), or west (`<`)<br />
After each move, exactly one pizza is delivered to that address<br />
If there are multiple deliverers, they will read the same route but take turns on which moves to take<br />
The first deliverer will complete their route before the second one starts theirs and so on<br />
A Route of `^v^v^v` with 2 deliverers is the same as 1 deliverer with route `^^^` and another with route `vvv`<br />
If a deliverer can't interpret their next move (invalid input), or the route takes the deliverer past the edge of the neighborhood, all deliveries will stop immediately

## Installation
## Usage
To use Tom's Pizza Delivery Service, you need to provide the number of deliverers and the route you wish them to take
The number of deliverers is passed through the command line and the route can be passed from the command line or as a file


To pass a route as a file, add it to the routes directory first, then call the program using the filename as a parameter<br />
In this example `PizzaDeliveryInput.txt` is the filename
* `go run main.go PizzaDeliveryInput.txt`<br />
  Welcome to Thomas's Pizza Delivery Company!<br />
  Please enter the number of deliverers<br />
  `2`<br />
  Houses Visited = 2639
  
  
To pass the route through the command line, simply run the program and follow the prompts
* `go run main.go`<br />
  Welcome to Thomas's Pizza Delivery Company!<br />
  Please enter the number of deliverers<br />
  `3`<br />
  Please enter the desired route<br />
  `^^^vvv`<br />
  Houses Visited = 2
  
### Enjoy your pizza!

