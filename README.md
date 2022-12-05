Simple REST API with two API endpoints (one that accepts GET requests and returns posts, and one which accepts POST requests and accepts posts to be added to the database and fetched later). The focus of this code was to follow the Clean Architecture principles, and some things from Domain Driven Design. NOTE: there are some things that are not the most optimal here, e.g "Accept interfaces return structs" is not used here, and the interfaces are not being defined in the client libraries, but this is due to the fact that the tutorial I was kinda following had a dude who must've come over from the Java word. This needs to be revisited in the future.

Clean architecture can mean a lot of things to a lot people (there are no "best number of layers", it really depends on the project and the dev teams), but there are some guidelines which should be followed:

## the code must be dependent on abstractions and not implementations:

this code is a great example of this.
If we look at router.go, there's a Router interface defined. This means, that anything that wants to act as a router (so wants to be able to redirect HTTP traffic from the client to the server and do something with the request) must implement three methods: GET,POST,Serve. After this, we have 3rd party library specific files (mux-router.go and chi-router.go) in which we are creating structs that implement the three functions that are defined by the Router interface in a way that is defined in the 3rd party libraries.
E.g if we were to create our own "adam-router.go" file and tried to only use the standard library, these functions (GET,POST) would only have functions that are part of the standard library, and we wouldn't need any variable that holds the connection to any third party libraries like the chiDispatcher or muxDispatcher.
Another example for this is the repository. According to our posts-repo.go, a repository (something that stores our posts) must satisfy the PostRepository interface, meaning the repository struct (regardless of type, it could be MySQL, Postgre, Firebase, in-memory...) must have a Save and a FindAll method defined. As it's with routers, we have a framework specific file defined (firestore-repo.go) where the Firestore repo struct implements the FindAll and Save methods in a way that these functions reach out to the Firestore API and are able to save data to a Firestore database.

## the code must be easy to test:

because of all the above, it's really easy to write tests (unit-tests in this case) to a code that is following clean architecture principles. When you would like to test the logic of post retrieval without having to establish connection to an actual database (which should be part of integration testing, and not unit-tests) you just create an in-memory database that implements the PostRepository interface, and you're good to go.

## your code is not dependent on any framework:

our business logic (which is really easy to understand here) is NOT dependent on any frameworks. Because of the abstractions used, e.g if we were using the FireStore database in our live code, and there would be a big change in the FireStori API, our business logic would not be affected, only the one part where we actually implement the FireStore repo -> minimalizing the blast radius

## your code is easily maintainable, extendable, easy to modify

if the clean architecture principles are being followed and basically everything is abstracted away, and the Depdendecy Inversion Principle is used, it should be easy to add new features to our code, and modify (e.g switch the db used by our code) without breaking anything.

## the dependencies in your code can only point inwards, and there must be no circular dependencies (the inner layers must never depend on the outer layers)

Imagine a world where your business logic was somehow dependent on the database logic you are using. If you were to change something in the db layer, you would have to change the business logic as well... Instead, our business logic must not know what's happening outside of its own realm. This principle can be observed perfectly in server.go, where the postService gets the postRepository as its dependency, and the PostController gets the postService as its dependency (moving from inside to out: Entity -> Service -> Controller -> Repository)
![image](https://user-images.githubusercontent.com/28739032/203933445-30557ee3-d00c-4364-a207-e902ecdb613a.png)
