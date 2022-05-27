package database

import (
	"github.com/joho/godotenv"
)

type DriverSupabase struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func readEnvSupabase() (DriverSupabase, error) {
	envSupabase, err := godotenv.Read()
	if err != nil {
		return DriverSupabase{}, err
	}
	return DriverSupabase{
		User:     envSupabase["SUPABASE_USER"],
		Password: envSupabase["SUPABASE_PASSWORD"],
		Host:     envSupabase["SUPABASE_HOST"],
		Port:     envSupabase["SUPABASE_PORT"],
		DbName:   envSupabase["SUPABASE_DB_NAME"],
	}, nil
}

func NewDriverSupabase() (DriverSupabase, error) {
	env, err := readEnvSupabase()
	if err != nil {
		return DriverSupabase{}, err
	}
	return env, nil
}
