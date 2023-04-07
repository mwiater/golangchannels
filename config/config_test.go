package config_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAppConfig(t *testing.T) {
	cfg, err := godotenv.Read("../.env")
	if err != nil {
		t.Error("Error loading .env file.", err)
	}
	var testMap = make(map[string]string)
	assert.IsType(t, testMap, cfg)
}
