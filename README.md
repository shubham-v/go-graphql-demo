#### Init modules
go mod init git@github.com:shubham-v/go-graphql-demo.git

#### Init graphql
go run github.com/99designs/gqlgen init

#### Generate resolvers from graphqls
go run github.com/99designs/gqlgen generate

#### Postgres Driver
go get -u github.com/jackc/pgx/v4

#### Postgres tag for migration
go build -tags 'postgres' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/

#### Create Migration Scripts
cd internal/pkg/db/migrations/
migrate create -ext sql -dir postgres -seq create_users_table
migrate create -ext sql -dir postgres -seq create_links_table

#### Database Migration
migrate -database postgres://postgres:postgres@/news -path internal/pkg/db/migrations/postgres up




#### References
- https://www.howtographql.com/graphql-go/
- https://gowalker.org/github.com/jackc/pgx/stdlib
- https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4
- https://www.postgresql.org/docs/9.3/sql-prepare.html
- https://gqlgen.com/recipes/authentication/
