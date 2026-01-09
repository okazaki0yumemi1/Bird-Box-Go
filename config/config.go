package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Conf struct {
	Server ConfServer
	DB     ConfDB
}

type ConfServer struct {
	Port         int `env:"SERVER_PORT,required"`
	TimeoutIdle  time.Duration
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
}

type ConfDB struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
}

func New() *Conf {
	var c Conf
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.DB = *NewDB()
	c.Server = *NewServer()
	return &c
}

func NewDB() *ConfDB {
	var c ConfDB
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.Host = os.Getenv("DB_HOST")
	c.Port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Unable to parse string to int (DB_PORT)")
	}
	c.Username = os.Getenv("DB_USER")
	c.Password = os.Getenv("DB_PASS")
	c.DBName = os.Getenv("DB_NAME")
	return &c
}

func NewServer() *ConfServer {
	var c ConfServer
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c.Port, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Unable to parse string to int (DB_PORT)")
	}
	c.TimeoutIdle = 10 * time.Second
	c.TimeoutRead = 10 * time.Second
	c.TimeoutWrite = 10 * time.Second
	return &c
}
