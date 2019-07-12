# http-client-wrapper
This package is to serve the need of cleaning up projects that tend to use a lot of the http calls. The project is to support the consolidation of useful functionality into one liners. 

I felt the state of http options are great in go, but seemed to need better constructors for reducing the size of code. In multiple project that I've worked on it felt that the code felt verbose around setting up http clients and managing hystrix streams.

While the package is meant to be a helpful feature in cleaning up code, there are a few things that maybe be over looked when it comes to requests.

## Idea behind the package
In my experience I believe the best way to use the package is like this:
* Create http clients for each specific endpoint you will hit. This will create better throughput from the client side.
* Make hystrix streams for all of these calls. We have to have circuit breakers in place to ensure that we don't propagate problems upstream.
* Find the easy way to solve the problem. Developers should always emphasize readability. Even if you're working on your own, this will help in the long run.

## How to use