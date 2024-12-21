package main

import (
	"log"
	"net/http"

	"github.com/abdisetiakawan/go-restfulapi/internal/config"
	"github.com/abdisetiakawan/go-restfulapi/internal/controller"
	"github.com/abdisetiakawan/go-restfulapi/internal/repository"
	"github.com/abdisetiakawan/go-restfulapi/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
    // Initialize database
    db := config.NewDB()
    // Initialize validation
    categoryRepository := repository.NewCategoryRepository()
    categoryService := service.NewCategoryService(categoryRepository, db)
    categoryController := controller.NewCategoryController(categoryService)

    router := httprouter.New()
    router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        w.Write([]byte("Hello World!"))
    })
    router.POST("/api/categories", categoryController.Create)
    router.GET("/api/categories/:id", categoryController.FindById)
    router.GET("/api/categories", categoryController.FindAll)
    router.PUT("/api/categories/:id", categoryController.Update)
    router.DELETE("/api/categories/:id", categoryController.Delete)


    // Start server
    server := http.Server{
        Addr:    "localhost:3000",
        Handler: router,
    }
    log.Println("Starting server on localhost:3000")
    err := server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}