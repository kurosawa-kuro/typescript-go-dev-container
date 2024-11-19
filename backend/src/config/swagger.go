package config

import "backend/src/docs"

// SetupSwagger initializes swagger documentation settings
func SetupSwagger() {
	docs.SwaggerInfo.Title = "Your API Title"
	docs.SwaggerInfo.Description = "Your API Description"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
