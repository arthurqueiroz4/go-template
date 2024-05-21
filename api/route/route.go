package route

import (
	"crud-golang/api/controller"
	"crud-golang/config"
	"crud-golang/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Setup(DB *gorm.DB, cfg *config.Cfg) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cdb := database.NewCategoryDB(DB)

	cc := controller.NewCategoryController(cdb)

	r.Route("/categories", func(r chi.Router) {
		r.Post("/", cc.CreateCategory)
		r.Get("/{id}", cc.GetCategory)
		r.Delete("/{id}", cc.DeleteCategory)
		r.Put("/{id}", cc.UpdateCategory)
		r.Get("/", cc.GetAllCategory)
	})

	log.Println("Server running on port", cfg.WebServerPort)
	log.Fatal(http.ListenAndServe(cfg.WebServerPort, r))
}
