package config_test

import (
	"endeus/wawan/internal/config"
	"endeus/wawan/pkg/endeus"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type envValue struct {
	value        string
	defaultValue string
}

var _ = Describe("Config", func() {
	var (
		cfg    endeus.Config
		envVar = map[string]envValue{
			"ENVIRONMENT": {
				value: "dev", defaultValue: "dev",
			},
			"SERVER_PORT": {
				value: "3000", defaultValue: "5000",
			},
			"LOG_LEVEL": {
				value: "DEBUG", defaultValue: "DEBUG",
			},

			"DB_USERNAME": {
				value: "user", defaultValue: "",
			},
			"DB_PASSWORD": {
				value: "pass", defaultValue: "",
			},
			"DB_DATABASE": {
				value: "app_test", defaultValue: "",
			},
			"DB_HOST": {
				value: "127.0.0.1", defaultValue: "127.0.0.1",
			},
			"DB_PORT": {
				value: "3306", defaultValue: "3306",
			},
		}
	)

	Context("When environment variables are set", func() {
		It("Fills up from environment variable", func() {
			for name, env := range envVar {
				os.Setenv(name, env.value)
			}

			defer func() {
				for name := range envVar {
					os.Unsetenv(name)
				}
			}()

			cfg = config.New()

			// server env
			Expect(cfg.Environment()).To(Equal(envVar["ENVIRONMENT"].value))
			Expect(cfg.LogLevel()).To(Equal(envVar["LOG_LEVEL"].value))
			Expect(cfg.ServerPort()).To(Equal(envVar["SERVER_PORT"].value))

			// db
			Expect(cfg.DBUsername()).To(Equal(envVar["DB_USERNAME"].value))
			Expect(cfg.DBPassword()).To(Equal(envVar["DB_PASSWORD"].value))
			Expect(cfg.DBName()).To(Equal(envVar["DB_DATABASE"].value))
			Expect(cfg.DBHost()).To(Equal(envVar["DB_HOST"].value))
			Expect(cfg.DBPort()).To(Equal(envVar["DB_PORT"].value))
			addressDb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				envVar["DB_USERNAME"].value,
				envVar["DB_PASSWORD"].value,
				envVar["DB_HOST"].value,
				envVar["DB_PORT"].value,
				envVar["DB_DATABASE"].value,
			)
			Expect(cfg.DBAddress()).To(Equal(addressDb))
		})
	})

	Context("When environment variables are NOT set", func() {
		It("Uses default value", func() {
			cfg = config.New()

			// server
			Expect(cfg.LogLevel()).To(Equal(envVar["LOG_LEVEL"].defaultValue))
			Expect(cfg.ServerPort()).To(Equal(envVar["SERVER_PORT"].defaultValue))

			// db
			Expect(cfg.DBHost()).To(Equal(envVar["DB_HOST"].defaultValue))
			Expect(cfg.DBPort()).To(Equal(envVar["DB_PORT"].defaultValue))
		})
	})

})
