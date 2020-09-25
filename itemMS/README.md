# item microservice

* GET or POST /helloitem
  * a hello message of item
* GET /items
  * list all the items in the mock database
* POST /items
  * insert a new item into mock database, itemID is not needed

# how to run in docker

* build the image

_$ docker build -t simple-item-ms ._

* use the image to run

_$ docker run -p 55555:55555 simple-item-ms_
