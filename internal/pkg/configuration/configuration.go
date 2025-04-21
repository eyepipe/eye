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

	// ServerBaseURL with scheme and port (if needed)
	// like "https://api.eyepipe.pw"
	// by default, eye takes the server base URL from the request headers.
	ServerBaseURL string `yaml:"server_base_url"`

	// Server limits
	ServerSingleUploadBytesLimit int64 `yaml:"server_single_upload_bytes_limit"`
	ServerShardWriteBytesLimit   int64 `yaml:"server_shard_write_bytes_limit"`
	ServerShardWriteCounterLimit int64 `yaml:"server_shard_write_counter_limit"`
	ServerShardReadBytesLimit    int64 `yaml:"server_shard_read_bytes_limit"`
	ServerShardReadCounterLimit  int64 `yaml:"server_shard_read_counter_limit"`
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
