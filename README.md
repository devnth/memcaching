# todo-api-microservice

- create container for postgresq
````bash
    docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=dbname \
  -p 5438:5432 \
  postgres:14-alpine
  ````


- make memecached container
````
    docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine
  ````

