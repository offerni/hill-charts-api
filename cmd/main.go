package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase"
	"github.com/offerni/hill-charts-api/graph"
	"github.com/offerni/hill-charts-api/graph/generated"
	"github.com/offerni/hill-charts-api/squad"
	"github.com/rs/cors"
	"google.golang.org/api/option"
)

const graphQLDefaultPort = "9092"
const httpDefaultPort = "9091"

func main() {
	_ = godotenv.Load()

	db, err := newDB()
	if err != nil {
		log.Panicln("newDB Error:", err)
	}
	defer db.Close()

	organizationRepo, err := firebase.NewOrganizationRepository(db)
	if err != nil {
		hcerrors.Wrap("firebase.NewOrganizationRepository", err)
	}

	squadRepo, err := firebase.NewSquadRepository(db)
	if err != nil {
		hcerrors.Wrap("firebase.NewSquadRepository", err)
	}

	squadSvc, err := squad.NewService(squad.NewServiceOpts{
		OrganizationRepository: organizationRepo,
		SquadRepository:        squadRepo,
	})
	if err != nil {
		hcerrors.Wrap("squad.NewService", err)
	}

	gqlResolver := graph.NewResolver(graph.NewResolverOpts{
		SquadService: squadSvc,
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		initRESTServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		initGraphQLServer(gqlResolver)
	}()

	wg.Wait()
}

func newDB() (*firestore.Client, error) {
	ctx := context.Background()
	opts := option.WithCredentialsFile("./service_account.json")
	projectID := os.Getenv("PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID, opts)
	if err != nil {
		return nil, hcerrors.Wrap("firestore.NewClient", err)
	}

	return client, nil
}

func initGraphQLServer(gqlResolver *graph.Resolver) {
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: gqlResolver}))

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
		log.Fatal(hcerrors.Wrap("initRESTServer", err))
	}

}

func setHttpRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})
}
