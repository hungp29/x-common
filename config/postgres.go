package config

import "fmt"

// BuildPostgresURL assembles a PostgreSQL DSN from env vars using a shared prefix stem.
//
// Required: {prefix}_DB_HOST, {prefix}_DB_NAME, {prefix}_DB_USER, {prefix}_DB_PASSWORD.
// Optional: {prefix}_DB_PORT (default 5432), {prefix}_DB_SSLMODE (default disable), {prefix}_DB_SCHEMA (default xdata).
//
// Example: BuildPostgresURL("WISH") reads WISH_DB_HOST, WISH_DB_NAME, …
func BuildPostgresURL(prefix string) string {
	required := map[string]string{
		"DB_HOST":     MustGetString(prefix + "_DB_HOST"),
		"DB_NAME":     MustGetString(prefix + "_DB_NAME"),
		"DB_USER":     MustGetString(prefix + "_DB_USER"),
		"DB_PASSWORD": MustGetString(prefix + "_DB_PASSWORD"),
	}

	port := GetInt(prefix+"_DB_PORT", 5432)
	sslmode := GetString(prefix+"_DB_SSLMODE", "disable")
	schema := GetString(prefix+"_DB_SCHEMA", "xdata")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s",
		required["DB_USER"], required["DB_PASSWORD"],
		required["DB_HOST"], port,
		required["DB_NAME"], sslmode, schema,
	)
}
