package routes

import (
	"github-watcher/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Router struct {
	app *fiber.App
	db  *config.Database
}

func NewRouter(db *config.Database) *Router {
	return &Router{
		app: fiber.New(),
		db:  db,
	}
}

func (r *Router) SetupRoutes() {
	// Swagger documentation
	r.app.Get("/swagger/*", swagger.HandlerDefault)

	// User routes
	users := r.app.Group("/api/users")
	users.Post("/register", r.handleUserRegister)
	users.Post("/login", r.handleUserLogin)

	// Project routes
	projects := r.app.Group("/api/projects")
	projects.Get("/", r.handleListProjects)
	projects.Post("/", r.handleCreateProject)
	projects.Get("/:id", r.handleGetProject)
	projects.Put("/:id", r.handleUpdateProject)
	projects.Delete("/:id", r.handleDeleteProject)

	// Repository routes
	repos := r.app.Group("/api/repositories")
	repos.Get("/", r.handleListRepositories)
	repos.Post("/", r.handleCreateRepository)
	repos.Get("/:id", r.handleGetRepository)
	repos.Put("/:id", r.handleUpdateRepository)
	repos.Delete("/:id", r.handleDeleteRepository)
}

func (r *Router) GetApp() *fiber.App {
	return r.app
}
