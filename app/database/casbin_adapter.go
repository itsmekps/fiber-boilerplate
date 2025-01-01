package database

import (
	"fiber-boilerplate/config"
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
)

func InitCasbinEnforcer() (*casbin.Enforcer, error) {

	// Load configuration values (e.g., server port) from the configuration file
	v, err := config.InitConfig()
	if err != nil {
		log.Fatal(err) // Exit the application if the configuration fails to load
	}

	// Initialize MongoDB adapter
	dbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=CLG",
		v.GetString("Mongodb_user"),
		v.GetString("Mongodb_password"),
		v.GetString("Mongodb_host"),
	)

	adapter, err := mongodbadapter.NewAdapter(dbURI)
	if err != nil {
		log.Fatalf("Failed to initialize Casbin adapter: %v", err)
		return nil, err
	}

	// Initialize Casbin enforcer
	enforcer, err := casbin.NewEnforcer("config/casbin_model.conf", adapter)
	if err != nil {
		log.Fatalf("Failed to initialize Casbin enforcer: %v", err)
		return nil, err
	}

	// Load policies from the database
	if err := enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Failed to load policies: %v", err)
		return nil, err
	}

	// Debug: Print loaded policies
	policies := enforcer.GetPolicy()
	fmt.Printf("Loaded Policies: %+v\n", policies)

	// Add default admin policies if no policies exist
	if len(enforcer.GetPolicy()) == 0 {
		log.Println("No policies found. Adding default admin policies...")
		enforcer.AddPolicy("admin", "/api/admin/add-policy", "POST")
		enforcer.AddPolicy("admin", "/api/admin/remove-policy", "POST")
		enforcer.AddPolicy("admin", "/api/admin/get-policies", "GET")
		log.Println("Default admin policies added.")
	}

	return enforcer, nil
}
