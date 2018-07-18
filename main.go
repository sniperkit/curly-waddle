package main

import (
  "database/sql"
  "log"
  "net/http"

  _ "github.com/go-sql-driver/mysql"
  "github.com/tinrab/curly-waddle/graphql"
  "github.com/vektah/gqlgen/handler"
)

func main() {
  db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?parseTime=true")
  if err != nil {
    log.Fatal(err)
  }
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }

  s := graphql.NewGraphQLServer(db)
  http.Handle("/graphql", graphql.PostLoaderMiddleware(db, handler.GraphQL(graphql.NewExecutableSchema(s))))

  log.Println("Running on port 8080...")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
