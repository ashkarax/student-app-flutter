package config

import "github.com/spf13/viper"

type Keys struct {
	ApiKey string `mapstructure:"API_KEY"`
}

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBHost     string `mapstructure:"DBHOST"`
	DBPort     string `mapstructure:"DBPORT"`
}

type AWS struct {
	Region     string `mapstructure:"AWS_REGION"`
	AccessKey  string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecrectKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	Endpoint   string `mapstructure:"AWS_ENDPOINT"`
}

type Config struct {
	DB     DataBase
	AwsS3  AWS
	ApiKey Keys
}

func LoadConfig() (*Config, error) {
	var db DataBase
	var awsS3 AWS
	var keys Keys

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&awsS3)
	if err != nil {
		return nil, err
	}

	config := Config{DB: db, AwsS3: awsS3, ApiKey: keys}
	return &config, nil
}
