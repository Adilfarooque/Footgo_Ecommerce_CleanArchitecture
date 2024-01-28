package config

type Config struct {
	BASE_URL   string `mapstructure:"BASE_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AUTHTOKEN  string `mapstructure:"TWILIO_AUTHTOKEN"`
	ACCOUNTSID string `mapstructure:"TWILIO_ACCOUNTSID"`
	SERVICESID string `mapstructure:"TWILIO_SERVIECSID"`

	KEY       string `mapstructure:"KEY"`
	KEY_ADMIN string `mapstructure:"KEY_ADMIN"`

	KEY_ID_FOR_PAY     string `mapstructure:"KEY_ID_FOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`
}

func LoadConfig()(Config,error){
	
}
