# GolangGraphQL
This project shows how to implement GraphQL service API by Golang

### Tools using
- API server: https://github.com/gin-gonic/gin
- Graph: https://github.com/99designs/gqlgen
- Database: [GORM](https://gorm.io/), [Migration](https://github.com/golang-migrate/migrate)
- Dependent Injection: Wire

### Run Demo
First of all you need setup Mysql in your local environment.
##### Migrate Database by run
```sh
export MIGRATION_DIR=your_SQL_scriptions
export DATABASE_URL=your_mysql_URI
make migrate 
```
##### Generate graphql resolvers in Golang
Default directory is **graph**
```sh
make gqlgen
```
##### Generate indepent injection code
```sh
make di
```

##### Start API
```sh
# cd to you project folder 
go run ./apps/server
```
After run abover command your service will be started on default port 8080.
You can use POSTMAN or another Graphql request [tool](https://github.com/graphql/graphql-playground) to try example API
###### create a new product
```sh
curl --location 'localhost:8080/v1/mutation' \
--header 'Content-Type: application/json' \
--data '{"query":"mutation {\n createProduct(\n input:{\n name: \"Iphone 14 pro max\",\n sku: \"4234265565\",\n properties: {\n price: 999999999,\n color: \"PINK\",\n size:\"6.5'\''0\"\n}\n}){\n id,\n name,\n sku,\n properties {\n price,\n color,\n size\n}\n}\n}","variables":{}}'
```
###### filter products created
```sh
curl --location 'localhost:8080/v1/query' \
--header 'Content-Type: application/json' \
--data '{"query":"query {\n filter(page:1, limit:1){\n id,\n name,\n sku,\n properties{\n price,\n color,\n size\n}\n}\n}","variables":{}}'
```
### Document References
If you have been not aware about GraphQL API yet, you can learn [here](https://graphql.org/learn/)
Or find a tutorial suitable with your programing language [here] (https://www.howtographql.com/basics/1-graphql-is-the-better-rest/)
