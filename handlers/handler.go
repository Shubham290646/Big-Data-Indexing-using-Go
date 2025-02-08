package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"rest-api-json-handler/models"
	"rest-api-json-handler/storage"

	"github.com/labstack/echo/v4"
	"github.com/xeipuuv/gojsonschema"
)

// JSON schema validation
const schemaFile = "schema/schema.json"

// Validate JSON structure using JSON Schema
func validateJSON(data interface{}) bool {
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaFile)
	documentLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false
	}
	return result.Valid()
}

// Generate ETag based on content hash
func generateETag(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// Create a new plan (POST) with ETag generation
func CreatePlan(c echo.Context) error {
	var plan models.Plan
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid content"})
	}

	// Validate against JSON schema
	if !validateJSON(plan) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid content"})
	}

	// Check if plan already exists
	existingData, _ := storage.RedisClient.Get(context.Background(), plan.ObjectID).Result()
	if existingData != "" {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Plan already exists"})
	}

	// Convert to JSON
	data, err := json.Marshal(plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not serialize data"})
	}

	// Generate ETag for the resource
	etag := generateETag(data)

	// Store data and ETag in Redis
	err = storage.RedisClient.Set(context.Background(), plan.ObjectID, data, 0).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to store in Redis"})
	}

	// Store ETag in Redis
	err = storage.RedisClient.Set(context.Background(), plan.ObjectID+":etag", etag, 0).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to store ETag in Redis"})
	}

	// Return response with ETag in header
	c.Response().Header().Set("ETag", etag)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Plan created", "objectId": plan.ObjectID})
}

// Retrieve a plan (GET) with Conditional Read using ETag
func GetPlan(c echo.Context) error {
	objectID := c.Param("id")

	// Fetch plan data from Redis
	data, err := storage.RedisClient.Get(context.Background(), objectID).Result()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Plan not found"})
	}

	// Fetch ETag from Redis
	etag, err := storage.RedisClient.Get(context.Background(), objectID+":etag").Result()
	if err != nil {
		etag = generateETag([]byte(data)) // Generate ETag if not found
	}

	// Check If-None-Match header for conditional read
	if match := c.Request().Header.Get("If-None-Match"); match == etag {
		c.Response().Header().Set("ETag", etag)
		return c.NoContent(http.StatusNotModified) // 304 Not Modified
	}

	// Set ETag in response header
	c.Response().Header().Set("ETag", etag)
	return c.JSON(http.StatusOK, json.RawMessage(data))
}

// Delete a plan (DELETE)
func DeletePlan(c echo.Context) error {
	objectID := c.Param("id")

	// Check if plan exists before deleting
	existingData, err := storage.RedisClient.Get(context.Background(), objectID).Result()
	if err != nil || existingData == "" {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Plan not found"})
	}

	// Delete data and ETag from Redis
	err = storage.RedisClient.Del(context.Background(), objectID, objectID+":etag").Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete"})
	}

	return c.JSON(http.StatusNoContent, nil) // 204 No Content
}
