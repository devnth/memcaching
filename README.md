# todo-api-microservice

- start postgresql container
````
  docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5438:5432 \
  postgres:14-alpine
  ````

- migrate table task to containerised postresql using goose migration.
```
goose -dir db/migrations create table

goose -dir db/migrations postgres "user=user password=password dbname=dbname port=5438 sslmode=disable" up
```


- start memcached container
````
  docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine
  ````

