package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

// AppConfig defines application specific config
type AppConfig struct {
	Port                       int    `mapstructure:"port"`
	ReadTimeout                int    `mapstructure:"read_timeout"`
	WriteTimeout               int    `mapstructure:"write_timeout"`
	IdleTimeout                int    `mapstructure:"idle_timeout"`
	WarnLimit                  int    `mapstructure:"warn_limit"`
	BulkLimit                  int    `mapstructure:"bulk_limit"`
	RequestBodyLimit           string `mapstructure:"request_body_limit"`
	DateFormat                 string
	TimestampFormat            string
	MaxPageSize                int    `mapstructure:"max_page_size"`
	DefaultPageSize            int    `mapstructure:"default_page_size"`
	FileWriteDir               string `mapstructure:"file_write_dir"`
	Disable500ErrMsgInResponse bool   `mapstructure:"disable_500_err_msg_in_response"`

	MaxFileSizeInBytes      int64 `mapstructure:"max_file_size_in_bytes"`
	MaxImageFileSizeInBytes int64 `mapstructure:"max_image_file_size_in_bytes"`
}

type DBServer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// DatabaseConfig refers database specific config
type DatabaseConfig struct {
	Primary     DBServer      `mapstructure:"primary"`
	Secondary   DBServer      `mapstructure:"secondary"`
	Name        string        `mapstructure:"name"`
	Username    string        `mapstructure:"username"`
	Password    string        `mapstructure:"password"`
	SslMode     string        `mapstructure:"ssl_mode"`
	MaxLifeTime time.Duration `mapstructure:"max_life_time"`
	MaxIdleConn int           `mapstructure:"max_idle_conn"`
	MaxOpenConn int           `mapstructure:"max_open_conn"`
	Debug       bool          `mapstructure:"debug"`
	// size of the db insert batches
	InsertBatchSize int `mapstructure:"insert_batch_size"`
	MaxBatchSize    int `mapstructure:"max_batch_size"`
}

type MinioConfig struct {
	Url       string        `mapstructure:"url"`
	Bucket    string        `mapstructure:"bucket"`
	Region    string        `mapstructure:"region"`
	AccessKey string        `mapstructure:"access_key"`
	SecretKey string        `mapstructure:"secret_key"`
	Secure    bool          `mapstructure:"secure"`
	Expires   time.Duration `mapstructure:"expires"`
}

// c is the configuration instance
var c Config

// Get returns all configurations
func Get() Config {
	return c
}

// Set is only for test purpose
func Set(cf Config) {
	c = cf
}

func Load() error {
	if err := viper.BindEnv("consul_url"); err != nil {
		return err
	}
	if err := viper.BindEnv("consul_path"); err != nil {
		return err
	}
	if err := viper.BindEnv("consul_http_token"); err != nil {
		return err
	}

	consulURL := viper.GetString("consul_url")
	consulPath := viper.GetString("consul_path")
	consulHttpToken := viper.GetString("consul_http_token")
	if consulURL == "" {
		return errors.New("CONSUL_URL is missing from ENV")
	}
	if consulPath == "" {
		return errors.New("CONSUL_PATH is missing from ENV")
	}
	if consulHttpToken == "" {
		return errors.New("CONSUL_HTTP_TOKEN is missing from ENV")
	}

	// read config from remote consul
	viper.SetConfigType("yml")
	if err := viper.AddRemoteProvider("consul", consulURL, consulPath); err != nil {
		return fmt.Errorf("failed to add remote config provider: %v", err)
	}
	if err := viper.ReadRemoteConfig(); err != nil {
		return fmt.Errorf("failed to read remote config: %v", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("failed to unmarshal consul config: %v", err)
	}
	return nil
}
