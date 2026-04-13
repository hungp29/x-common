package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

// ===== STRING =====

func GetString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func MustGetString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		panic(fmt.Sprintf("missing required env: %s", key))
	}
	return val
}

// ===== INT =====

func GetInt(key string, defaultVal int) int {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		panic(fmt.Sprintf("invalid int env %s=%s", key, valStr))
	}
	return val
}

func MustGetInt(key string) int {
	valStr := MustGetString(key)
	val, err := strconv.Atoi(valStr)
	if err != nil {
		panic(fmt.Sprintf("invalid int env %s=%s", key, valStr))
	}
	return val
}

// ===== BOOL =====

func GetBool(key string, defaultVal bool) bool {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}

	val, err := strconv.ParseBool(valStr)
	if err != nil {
		panic(fmt.Sprintf("invalid bool env %s=%s", key, valStr))
	}
	return val
}

// ===== DURATION =====

func GetDuration(key string, defaultVal time.Duration) time.Duration {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}

	val, err := time.ParseDuration(valStr)
	if err != nil {
		panic(fmt.Sprintf("invalid duration env %s=%s", key, valStr))
	}
	return val
}

// ===== OTHER TYPES (e.g. zapcore.Level) can be added similarly =====

func GetLogLevel(key string, def zapcore.Level) zapcore.Level {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	var lvl zapcore.Level
	if err := lvl.UnmarshalText([]byte(strings.ToUpper(val))); err != nil {
		panic(fmt.Sprintf("%s must be one of DEBUG, INFO, WARN, ERROR: %v", key, err))
	}
	return lvl
}
