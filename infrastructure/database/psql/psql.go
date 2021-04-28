package psql

import (
	"fmt"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	User     string
	Password string
	DBName   string
	Hostname string
	Port     string
	Sslmode  string
	Timezone string
}

func InitDB() *gorm.DB {
	DB, err := connectDBUsingConf()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("DB connected")
	return DB
}

func connectDBUsingConf() (*gorm.DB, error) {
	PsqlConfig, err := readConf()
	if err != nil {
		return nil, err
	}

	DB, err := gorm.Open(postgres.Open(PsqlConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}

func readConf() (string, error) {
	viper.SetConfigName("psql")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./infrastructure/database/psql/")
	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	DBConfig := &config{}
	err = viper.Unmarshal(DBConfig)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DBConfig.Hostname, DBConfig.User, DBConfig.Password, DBConfig.DBName, DBConfig.Port, DBConfig.Sslmode, DBConfig.Timezone), nil
}

func Migrate(DB *gorm.DB) {
	err := DB.AutoMigrate(&domain.User{})

	if err != nil {
		fmt.Println("Migration Failed")
		fmt.Println(err)
	} else {
		fmt.Println("Migration Success")
	}
}
