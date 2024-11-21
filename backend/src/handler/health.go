package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	statusOK            = 200
	statusNotFound      = 404
	statusInternalError = 500
)

const (
	msgHealthy          = "Service is healthy"
	msgDBConnected      = "DB connected"
	msgFailedDBName     = "Failed to get database name"
	msgFailedConnection = "Failed to connect to test database"
)

type HealthHandler struct {
	db *gorm.DB
}

type healthDBCheckResult struct {
	expectedDB string
	actualDB   string
	err        error
}

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
}

func (c dbConfig) toString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.host, c.user, c.password, c.dbname, c.port, c.sslmode,
	)
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) CheckHealth(c *gin.Context) {
	c.JSON(statusOK, gin.H{"message": msgHealthy})
}

func (h *HealthHandler) verifyDatabaseName(db *gorm.DB, expectedDB string) healthDBCheckResult {
	var dbName string
	err := db.Raw("SELECT current_database()").Scan(&dbName).Error
	return healthDBCheckResult{
		expectedDB: expectedDB,
		actualDB:   dbName,
		err:        err,
	}
}

func (h *HealthHandler) sendDatabaseResponse(c *gin.Context, result healthDBCheckResult) {
	if result.err != nil {
		c.JSON(statusInternalError, gin.H{
			"error":   msgFailedDBName,
			"details": result.err.Error(),
		})
		return
	}

	if result.actualDB == result.expectedDB {
		c.JSON(statusOK, gin.H{
			"message":  fmt.Sprintf("Connected to %s database", result.expectedDB),
			"database": result.actualDB,
		})
		return
	}

	c.JSON(statusNotFound, gin.H{
		"error":            fmt.Sprintf("Not connected to %s database", result.expectedDB),
		"current_database": result.actualDB,
	})
}

func (h *HealthHandler) CheckDBConnection(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(statusInternalError, gin.H{"error": err.Error()})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(statusInternalError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusOK, gin.H{"message": msgDBConnected})
}

func (h *HealthHandler) CheckDevDatabase(c *gin.Context) {
	result := h.verifyDatabaseName(h.db, "dev_db")
	h.sendDatabaseResponse(c, result)
}

func (h *HealthHandler) CheckTestDatabase(c *gin.Context) {
	config := dbConfig{
		host:     "test-db",
		user:     "postgres",
		password: "postgres",
		dbname:   "test_db",
		port:     "5432",
		sslmode:  "disable",
	}

	testDB, err := gorm.Open(postgres.Open(config.toString()), &gorm.Config{})
	if err != nil {
		c.JSON(statusInternalError, gin.H{
			"error":   msgFailedConnection,
			"details": err.Error(),
		})
		return
	}

	result := h.verifyDatabaseName(testDB, config.dbname)
	h.sendDatabaseResponse(c, result)
}
