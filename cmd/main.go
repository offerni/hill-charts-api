package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/labstack/echo"
	"github.com/offerni/hill-charts-api/graph"
	"github.com/offerni/hill-charts-api/graph/generated"
	"github.com/rs/cors"
)

const graphQLDefaultPort = "9092"
const httpDefaultPort = "9091"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		initRESTServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		initGraphQLServer()
	}()

	wg.Wait()
}

func initGraphQLServer() {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	port := os.Getenv("GRAPHQL_PORT")
	if port == "" {
		port = graphQLDefaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func initRESTServer() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = httpDefaultPort
	}

	e := echo.New()
	setHttpRoutes(e)

	log.Printf("connect to http://localhost:%s/ for the REST API Server", port)
	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}

func setHttpRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})
}
