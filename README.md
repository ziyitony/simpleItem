Here are 2 HTTP services to build a monolith service: itemService and userService.

# itemService

* GET or POST /helloitem
  * a hello message of item
* GET /items
  * list all the items in the mock database
* POST /items
  * insert a new item into mock database, itemID is not needed
  
# userService

* GET or POST /hellouser
  * a hello message of user
* GET /users
  * list all the users in the mock database
* POST /users
  * insert a new user into mock database, userID is not needed
* GET /userid/{id}
  * get the user by id

# how to run in docker

* build the image

_$ docker build -t simple-service ._

* use the image to run

_$ docker run -d -p 12345:12345 --name simple-service simple-service_
