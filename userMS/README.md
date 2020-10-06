# user microservice

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

_$ docker build -t simple-user-ms ._

* use the image to run

_$ docker run -d -p 44444:44444 --name simple-user-ms simple-user-ms_
