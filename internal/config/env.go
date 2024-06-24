package config

import "os"

type Configuration struct {
	Server        ServerConfiguration
	Databases     Database
	TestDatabases Database
	App           App
	IPStack       IPStack
	Mail          Mail
	Termii        Termii
}

type BaseConfig struct {
	SERVER_PORT                      string `mapstructure:"SERVER_PORT"`
	SERVER_SECRET                    string `mapstructure:"SERVER_SECRET"`
	SERVER_ACCESSTOKENEXPIREDURATION int    `mapstructure:"SERVER_ACCESSTOKENEXPIREDURATION"`

	APP_NAME string `mapstructure:"APP_NAME"`
	APP_KEY  string `mapstructure:"APP_KEY"`

	CONNECTION_STRING string `mapstructure:"CONNECTION_STRING"`
	DB_NAME           string `mapstructure:"DB_NAME"`
	MIGRATE           bool   `mapstructure:"MIGRATE"`

	TEST_CONNECTION_STRING string `mapstructure:"TEST_CONNECTION_STRING"`
	TEST_DB_NAME           string `mapstructure:"TEST_DB_NAME"`
	TEST_MIGRATE           bool   `mapstructure:"TEST_MIGRATE"`

	IPSTACK_KEY      string `mapstructure:"IPSTACK_KEY"`
	IPSTACK_BASE_URL string `mapstructure:"IPSTACK_BASE_URL"`

	MAIL_DOMAIN          string `mapstructure:"MAIL_DOMAIN"`
	MAIL_PRIVATE_API_KEY string `mapstructure:"MAIL_PRIVATE_API_KEY"`
	MAIL_SENDER_EMAIL    string `mapstructure:"MAIL_SENDER_EMAIL"`

	TERMII_API_KEY  string `mapstructure:"TERMII_API_KEY"`
	TERMII_BASE_URL string `mapstructure:"TERMII_BASE_URL"`
}

func (config *BaseConfig) SetupConfigurationn() *Configuration {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.SERVER_PORT
	}
	return &Configuration{
		Server: ServerConfiguration{
			Port:                          port,
			Secret:                        config.SERVER_SECRET,
			AccessTokenExpirationDuration: config.SERVER_ACCESSTOKENEXPIREDURATION,
		},
		Databases: Database{
			ConnectionString: config.CONNECTION_STRING,
			DBName:           config.DB_NAME,
			Migrate:          config.MIGRATE,
		},
		TestDatabases: Database{
			ConnectionString: config.TEST_CONNECTION_STRING,
			DBName:           config.TEST_DB_NAME,
			Migrate:          config.TEST_MIGRATE,
		},
		App: App{
			Name: config.APP_NAME,
			Key:  config.APP_KEY,
		},
		IPStack: IPStack{
			Key:     config.IPSTACK_KEY,
			BaseUrl: config.IPSTACK_BASE_URL,
		},
		Mail: Mail{
			Domain:        config.MAIL_DOMAIN,
			PrivateApiKey: config.MAIL_PRIVATE_API_KEY,
			SenderEmail:   config.MAIL_SENDER_EMAIL,
		},
		Termii: Termii{
			ApiKey:  config.TERMII_API_KEY,
			BaseUrl: config.TERMII_BASE_URL,
		},
	}
}
