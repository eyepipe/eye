package configuration

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	// JWT
	JwtAlgo      string `yaml:"jwt_algo"`
	JwtIss       string `yaml:"jwt_iss"`
	JwtSecretHex string `yaml:"jwt_secret_hex"`

	// DB
	DDShardFiles []string `yaml:"db_shard_files"`
	DbPoolSize   int      `yaml:"db_pool_size"`

	// S3
	S3ShardDSN []string `yaml:"s3_shard_dsn"`
}

func New() *Configuration {
	return &Configuration{
		DDShardFiles: make([]string, 0),
	}
}

func (c *Configuration) BindYAMLFile(filename string) error {
	err := cleanenv.ReadConfig(filename, c)
	if err != nil {
		return fmt.Errorf("failed to read configuration: %w", err)
	}

	return nil
}
