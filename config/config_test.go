package config

import (
	"os"
	"reflect"
	"testing"
)

var testingEnvVaribles = map[string]string{
	"DATABASE_DIALECT":  "dialect-mysql",
	"DATABASE_HOST":     "10.13.37.1",
	"DATABASE_PORT":     "1337",
	"DATABASE_USER":     "database-root",
	"DATABASE_PASSWORD": "database-password",
	"DATABASE_NAME":     "database-name",
	"WEBSERVER_ADDRESS": "10.13.37.1",
	"WEBSERVER_PORT":    "1234",
}

func setTestEnvVariables() {
	for key, value := range testingEnvVaribles {
		os.Setenv(key, value)
	}
}

func unsetTestEnvVariables() {
	for key, _ := range testingEnvVaribles {
		os.Unsetenv(key)
	}
}

func TestGetFromEnvAsString(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestDefault", args{"wasd", "default"}, "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFromEnvAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFromEnvAsInt(t *testing.T) {
	type args struct {
		key          string
		defaultValue int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"TestDefault", args{"wasd", 1}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFromEnvAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfigWithDefaultValues(t *testing.T) {
	unsetTestEnvVariables()
	config := GetConfig()

	if config.DB.Dialect != "mysql" {
		t.Errorf("config.DB.Dialect = %s; want mysql", config.DB.Dialect)
	}
	if config.DB.Host != "127.0.0.1" {
		t.Errorf("config.DB.Host = %s; want 127.0.0.1", config.DB.Host)
	}
	if config.DB.Port != 3306 {
		t.Errorf("config.DB.Port = %d; want 3306", config.DB.Port)
	}
	if config.DB.Username != "root" {
		t.Errorf("config.DB.Username = %s; want root", config.DB.Username)
	}
	if config.DB.Password != "password" {
		t.Errorf("config.DB.Password = %s; want password", config.DB.Password)
	}
	if config.DB.Name != "test" {
		t.Errorf("config.DB.Name = %s; want test", config.DB.Name)
	}
	if config.Web.Address != "127.0.0.1" {
		t.Errorf("config.Web.Address = %s; want 127.0.0.1", config.DB.Name)
	}
	if config.Web.Port != "1337" {
		t.Errorf("config.Web.Port = %s; want 1337", config.DB.Name)
	}
}

func TestGetConfigWithValuesFromEnvironment(t *testing.T) {
	setTestEnvVariables()
	config := GetConfig()

	if config.DB.Dialect != "dialect-mysql" {
		t.Errorf("config.DB.Dialect = %s; want dialect-mysql", config.DB.Dialect)
	}
	if config.DB.Host != "10.13.37.1" {
		t.Errorf("config.DB.Host = %s; want 10.13.37.1", config.DB.Host)
	}
	if config.DB.Port != 1337 {
		t.Errorf("config.DB.Port = %d; want 1337", config.DB.Port)
	}
	if config.DB.Username != "database-root" {
		t.Errorf("config.DB.Username = %s; want database-root", config.DB.Username)
	}
	if config.DB.Password != "database-password" {
		t.Errorf("config.DB.Password = %s; want database-password", config.DB.Password)
	}
	if config.DB.Name != "database-name" {
		t.Errorf("config.DB.Name = %s; want database-name", config.DB.Name)
	}
	if config.Web.Address != "10.13.37.1" {
		t.Errorf("config.Web.Address = %s; want 10.13.37.1", config.DB.Name)
	}
	if config.Web.Port != "1234" {
		t.Errorf("config.Web.Port = %s; want 1234", config.DB.Name)
	}

	unsetTestEnvVariables()
}
