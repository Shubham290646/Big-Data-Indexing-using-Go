package handlers

import (
	"context"
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

// Create a new plan (POST)
func CreatePlan(c echo.Context) error {
	var plan models.Plan
	if err := c.Bind(&plan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	// Validate against JSON schema
	if !validateJSON(plan) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "JSON does not match schema"})
	}

	// Convert to JSON and store in Redis
	data, err := json.Marshal(plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not serialize data"})
	}

	err = storage.RedisClient.Set(context.Background(), plan.ObjectID, data, 0).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to store in Redis"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Plan created", "objectId": plan.ObjectID})
}

// Retrieve a plan (GET) with Conditional Read
func GetPlan(c echo.Context) error {
	objectID := c.Param("id")

	data, err := storage.RedisClient.Get(context.Background(), objectID).Result()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Plan not found"})
	}

	var plan models.Plan
	err = json.Unmarshal([]byte(data), &plan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not deserialize data"})
	}

	return c.JSON(http.StatusOK, plan)
}

// Delete a plan (DELETE)
func DeletePlan(c echo.Context) error {
	objectID := c.Param("id")

	err := storage.RedisClient.Del(context.Background(), objectID).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Plan deleted"})
}
