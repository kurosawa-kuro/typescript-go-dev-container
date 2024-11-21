package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PingDBHandler struct {
	db *gorm.DB
}

func NewPingDBHandler(db *gorm.DB) *PingDBHandler {
	return &PingDBHandler{db: db}
}

type dbCheckResult struct {
	expectedDB string
	actualDB   string
	err        error
}

func (h *PingDBHandler) verifyDatabaseName(db *gorm.DB, expectedDB string) dbCheckResult {
	var dbName string
	err := db.Raw("SELECT current_database()").Scan(&dbName).Error
	return dbCheckResult{
		expectedDB: expectedDB,
		actualDB:   dbName,
		err:        err,
	}
}

func (h *PingDBHandler) sendDatabaseResponse(c *gin.Context, result dbCheckResult) {
	if result.err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to get database name",
			"details": result.err.Error(),
		})
		return
	}

	if result.actualDB == result.expectedDB {
		c.JSON(200, gin.H{
			"message":  "Connected to " + result.expectedDB + " database",
			"database": result.actualDB,
		})
	} else {
		c.JSON(404, gin.H{
			"error":            "Not connected to " + result.expectedDB + " database",
			"current_database": result.actualDB,
		})
	}
}

func (h *PingDBHandler) CheckConnection(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "DB connected"})
}

func (h *PingDBHandler) CheckDevDatabase(c *gin.Context) {
	result := h.verifyDatabaseName(h.db, "dev_db")
	h.sendDatabaseResponse(c, result)
}

func (h *PingDBHandler) CheckTestDatabase(c *gin.Context) {
	testDBConfig := "host=test-db user=postgres password=postgres dbname=test_db port=5432 sslmode=disable"
	testDB, err := gorm.Open(postgres.Open(testDBConfig), &gorm.Config{})
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to connect to test database",
			"details": err.Error(),
		})
		return
	}

	result := h.verifyDatabaseName(testDB, "test_db")
	h.sendDatabaseResponse(c, result)
}
