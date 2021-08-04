# garys-tricks-go
Literally a pet project for writing a REST API in Go, which describes the tricks that my dog Gary can do or should learn :)

## Design
Root URL: `/garys-tricks/api/v1.0/`

API HTTP methods

`GET`:
* retrieve a list of all tricks Gary can already do:
  * `/garys-tricks/api/v1.0/gary-can-do/`
* retrieve a specific trick that Gary already can do:
  * `/garys-tricks/api/v1.0/gary-can-do/[trick_id]`
* retrieve a list of all tricks Gary should learn:
  * `/garys-tricks/api/v1.0/gary-should-learn/`
* retrieve a specific trick that Gary should learn do:
  * `/garys-tricks/api/v1.0/gary-should-learn/[trick_id]`

`POST`:
* add a trick to the tricks Gary already can do:
  * `garys-tricks/api/v1.0/gary-can-do/`
* add a trick to the list of all tricks Gary should learn:
  * `/garys-tricks/api/v1.0/gary-should-learn/`

`PUT`:
* update an existing trick that Gary already can do:
  * `/garys-tricks/api/v1.0/gary-can-do/[trick_id]`
* update an existing trick that Gary should learn do:
  * `/garys-tricks/api/v1.0/gary-should-learn/[trick_id]`

`DELETE`:
* delete a trick from the list of tricks that Gary already can do:
  * `/garys-tricks/api/v1.0/gary-can-do/[trick_id]`
* delete a trick from the list of tricks that Gary should learn:
  * `/garys-tricks/api/v1.0/gary-should-learn/[trick_id]`

The tricks are defined with the following fields:
* id: unique identifier for each trick, uuid type
* name: name of the trick, string type
* progress: current progress of the trick in training/learning, string type
* state: e.g. learned, in training, canceled, outstanding, string type
* added_on: date when the trick was added to the list
