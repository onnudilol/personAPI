# Deploying

## Install

`go get -u github.com/gorilla/mux`

`go get github.com/lib/pq`

`go build`

`./personAPI`

## Setting up the DB (with Docker)

`docker run --name personAPI_postgres -d -p 32768:5432 postgres`

`docker cp ./person.sql personAPI_postgres:/var/person.sql`

`docker exec -u postgres personAPI_postgres psql postgres postgres -f /var/person.sql
`

