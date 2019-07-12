# http-client-wrapper
This package is to serve the need of cleaning up projects that tend to use a lot of the http calls. The project is to support the consolidation of useful functionality into one liners. 

I felt the state of http options are great in go, but seemed to need better constructors for reducing the size of code. In multiple project that I've worked on it felt that the code was verbose around setting up http clients and managing hystrix streams.

While the package is meant to be a helpful feature in cleaning up code, there are a few things that may be overlooked when it comes to requests.

## Idea behind the package
In my experience the package is built around these ideals:
* Create http clients for each specific endpoint you will hit. This will create better throughput from the client side.
* Make hystrix streams for all of these calls. We have to have circuit breakers in place to ensure that we don't propagate problems upstream.
* Find the easy way to solve the problem. Developers should always emphasize readability. Even if you're working on your own, this will help in the long run.

## How to use
1) Initialize a client of the https-client-wrapper with one of the functions:
* InitializeClientWithTimeout(): Client doesn't have a timeout
* InitializeClientWithoutTimeout(): Client has a timeout

It is important to note that the client will be associated by the "api" parameter because of the ideas of this package: 1 api to 1 client.

2) Initialize the Hystrix Route for a specific api. This creates the hystrix stream with the name "{api}_{route}. The route parameter doesn't have to be the whole route, but should be an alias to what the route is accomplishing.

3) Call a method function. The client can do both asynchronous and synchronous calls. Asynchronous will return an error to the error channel or a byte array to the byte array channel that is specified to the method. Synchronous will either return a byte array or an error.

For any problems or if anything is unclear, feel free to message me.