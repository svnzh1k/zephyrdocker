package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"zephyr-api-mod/internal/handlers"
	"zephyr-api-mod/internal/middleware"
	"zephyr-api-mod/internal/service"
)

var Db *sql.DB

func main() {
	err := service.InitializeDatabase()
	if err != nil {
		log.Fatal("can't connect to the database", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST       /login", handlers.LoginHandler)

	// mux.Handle("POST    /admin/users", middleware.TokenValidator((middleware.AdminRoleValidator(handlers.RegisterUserHandler))))
	// mux.Handle("DELETE  /admin/users/{id}", middleware.TokenValidator(middleware.AdminRoleValidator(handlers.RemoveUserHandler)))
	// mux.Handle("GET     /admin/waiters/{id}", middleware.TokenValidator(middleware.AdminRoleValidator(handlers.GetUserHandler)))
	// mux.HandleFunc("GET /admin/waiters", handlers.GetWaitersHandler)

	// mux.Handle("POST   	/owner/categories/{parentId}/{name}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.CreateCategory)))
	// mux.Handle("PATCH  	/owner/categories/{id}/{newName}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.UpdateCategory)))
	// mux.Handle("DELETE 	/owner/categories/{id}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.RemoveCategory)))
	// mux.Handle("GET    	/owner/categories/{parentId}", middleware.TokenValidator(handlers.GetCategories))

	// mux.Handle("POST   	/owner/products/{name}/{unit}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.AddProduct)))
	// mux.Handle("PATCH   /owner/products/{id}/{newName}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.UpdateProduct)))
	// mux.Handle("DELETE  /owner/products/{id}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.RemoveProduct)))
	// mux.Handle("GET   	/owner/products/", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.GetProducts)))

	// mux.Handle("POST   	/owner/food", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.AddFood)))
	// mux.Handle("PATCH   /owner/food/{id}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.UpdateFood)))
	// mux.Handle("DELETE  /owner/food/{id}", middleware.TokenValidator(middleware.OwnerRoleValidator(handlers.RemoveFood)))
	// mux.HandleFunc("GET   	/food", handlers.GetFood)

	mux.HandleFunc("POST    /admin/users", handlers.RegisterUserHandler)
	mux.HandleFunc("DELETE  /admin/users/{id}", handlers.RemoveUserHandler)
	mux.HandleFunc("GET     /admin/waiters/{id}", handlers.GetUserHandler)
	mux.HandleFunc("GET /admin/waiters", handlers.GetWaitersHandler)

	mux.HandleFunc("POST    /owner/categories/{parentId}/{name}", handlers.CreateCategory)
	mux.HandleFunc("PATCH   /owner/categories/{id}/{newName}", handlers.UpdateCategory)
	mux.HandleFunc("DELETE  /owner/categories/{id}", handlers.RemoveCategory)
	mux.HandleFunc("GET     /owner/categories/{parentId}", handlers.GetCategories)

	mux.HandleFunc("POST    /owner/products/{name}/{unit}", handlers.AddProduct)
	mux.HandleFunc("PATCH   /owner/products/{id}/{newName}", handlers.UpdateProduct)
	mux.HandleFunc("DELETE  /owner/products/{id}", handlers.RemoveProduct)
	mux.HandleFunc("GET     /owner/products/", handlers.GetProducts)

	mux.HandleFunc("POST    /owner/food", handlers.AddFood)
	mux.HandleFunc("PATCH   /owner/food/{id}", handlers.UpdateFood)
	mux.HandleFunc("DELETE  /owner/food/{id}", handlers.RemoveFood)
	mux.HandleFunc("GET     /food", handlers.GetFood)

	wrappedMux := middleware.CORS(mux)
	fmt.Println("Application started at port 8080")
	http.ListenAndServe(":8080", middleware.RequestLogger(wrappedMux.ServeHTTP))
}
