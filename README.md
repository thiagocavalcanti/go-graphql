
# go-graphql
An implementation of graphql using go based in howtographql tutorial

## Getting started
This project is based on https://www.howtographql.com/graphql-go, but i found many problems during the development in the tutorial
and I'm addressing it here.

### Tip 1 - Regenerate go core logic
To regenerate the core business along with models, remove `graph/generated/*` and `graph/model/*` and run `go run github.com/99designs/gqlgen`

### Tip 2 - Updated gqlgen generated files
Everytime the tutorial says to change in `resolver.go` , you should change `graph/schema.resolvers.go` instead

### Tip 3 - Configuring Mysql docker
To configure a mysql database use:
```bash
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -d mysql:latest
```

If you try to connect with Dbeaver, you'll get an **caching_sha2_password** error. To solve this problem, open a terminal and run the following commands:
```bash
**docker exec -it mysql bash**
mysql -u root -p
use mysql  
select host, user from user;
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'Your New Strong Password';
FLUSH PRIVILEGES;
```
Reference: [https://medium.com/@rauf.rahmancz/connect-docker-and-mysql-in-right-way-95602f833cb0](https://medium.com/@rauf.rahmancz/connect-docker-and-mysql-in-right-way-95602f833cb0)

Also, you need to create the table for this project:
```bash
CREATE DATABASE hackernews;
```

### Tip 4 - New version of golang-migrate
The version of golang-migrate used in the tutorial is not working anymore (v4 is the only version not depreceated)
```bash
go get -u github.com/go-sql-driver/mysql
go build -tags 'mysql' -ldflags="-X main.Version=$(git describe --tags)" -o $GOPATH/bin/migrate github.com/golang-migrate/v4/cmd/migrate
cd internal/pkg/db/migrations/ #Create this directory
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table
```
And then to run the migration:
```bash
migrate -database mysql://root:dbpass@localhost:3306/hackernews -path internal/pkg/db/migrations/mysql up
```

## How to run

```bash
git clone https://github.com/thiagocavalcanti/go-graphql.git
go run github.com/99designs/gqlgen #This will generate the models and core graphql logic (generated.go)
go run server.go
```


