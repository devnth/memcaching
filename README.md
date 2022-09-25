# todo-api-microservice

- create container for postgresq
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


- create container for memecaching
````
    docker run \
  -d \
  -p 11211:11211 \
  memcached:1.6.9-alpine
  ````

