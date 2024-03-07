# Book my Event
event booking REST API using Gin framework

## Current endpoints

- `GET /events` get available events
- `GET /events/{id}` get specific event
- `POST /events` create new event <auth>
- `PUT /events` update an event <auth>
- `DELETE /events` delete an event <auth>
- `POST /events` create user
- `POST /login` authenticate
- `POST /events/{id}/register` register for an event <auth>
- `DELETE /events/{id}/register` cancel registration <auth>