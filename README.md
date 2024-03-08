# Book my Event
event booking REST API using
- Gin Framework
- PgSql (Sqlx + pgx)
- JWT

## Current endpoints

- `GET /events` get available events
- `GET /events/{id}` get specific event
- `POST /events` create new event <auth>
- `PUT /events/{id}` update an event <auth>
- `DELETE /events/{id}` delete an event <auth>
- `POST /signup` create user
- `POST /login` authenticate
- `POST /events/{id}/register` register for an event <auth>
- `DELETE /events/{id}/register` cancel registration <auth>