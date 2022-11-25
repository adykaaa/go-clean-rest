Simple REST API with two API endpoints (one that accepts GET requests and returns posts, and one which accepts POST requests and accepts posts to be added to the database and fetched later). The focus of this code was to follow the Clean Architecture principles. NOTE: there are some things that are not the most optimal here, e.g "Accept interfaces return structs" is not used here, and the interfaces are not being defined in the client libraries, but this is due to the fact that the tutorial I was kinda following had a dude who must've come over from the Java word. This needs to be revisited in the future.

Clean architecture can mean a lot of things to a lot people (there are no "best number of layers", it really depends on the project and the dev teams), but there are some guidelines which should be followed:

× the code must be dependent on abstractions and not implementations:
-> this code is a great example of this.
If we look at router.go, there's a Router interface defined. This means, that anything that wants to act as a router (so wants to be able to redirect HTTP traffic from the client to the server and do something with the request) must implement three methods: GET,POST,Serve. After this, we have 3rd party library specific files (mux-router.go and chi-router.go) in which we are creating structs that implement the three functions that are defined by the Router interface in a way that is defined in the 3rd party libraries.
E.g if we were to create our own "adam-router.go" file and tried to only use the standard library, these functions (GET,POST) would only have functions that are part of the standard library, and we wouldn't need any variable that holds the connection to any third party libraries like the chiDispatcher or muxDispatcher.

Another example for this is the repository. According to our posts-repo.go, a repository (something that stores our posts) must satisfy the PostRepository interface, meaning the repository struct (regardless of type, it could be MySQL, Postgre, Firebase, in-memory...) must have a Save and a FindAll method defined.
