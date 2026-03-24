package config

import {
	"errors"
	"fmt"
	"os"
	"strings"
}

type Config struct {
	Port string
	DatabaseURL string
	SuppabaseURL string
	SupabaseAnonKey	string
	SuppabaseServiceRoleKey string
}

func Load() (Config, error) {
	cfg := Config{
		Port: getEnv("PORT", "8080"),
		DatabaseURL: strings.TrimSpace(os.Getenv("DATABASE_URL")),
		SuppabaseURL: strings.TrimSpace(os.Getenv("SUPABASE_URL")),
		SupabaseAnonKey: strings.TrimSpace(os.Getenv("SUPABASE_ANON_KEY")),
		SuppabaseServiceRoleKey: strings.TrimSpace(os.Getenv("SUPABASE_SERVICE_ROLE_KEY")),
	}
	var missing []string
	if cfg.DatabaseURL == "" {
		missing = append(missing, "DATABASE_URL")
	} if cfg.SuppabaseURL == "" {
		missing = append(missing, "SUPABASE_URL")
	} if cfg.SuppabaseAnonKey == "" {
		missing = append(missing, "SUPABASE_ANON_KEY")
	} if cfg.SuppabaseServiceRoleKey == "" {
		missing = append(missing, "SUPABASE_SERVICE_ROLE_KEY")
	}

	if len(missing) > 0 {
		return cfg, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	return v
}

func MustLoad() Config {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	} 
	return cfg
}