package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigOptions struct{
	Version string
	ServiceName string
	Port int
	PostgresUri string
} 

var config ConfigOptions

func loadConfig(){
	err := godotenv.Load()
	if err!= nil{
		log.Fatal("Error while loading env file\n",err)
	}

	version := os.Getenv("VERSION")
	if version == ""{
		log.Fatal("Can not load VERSION name from env\n")
		os.Exit(1);
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == ""{
		log.Fatal("Can not load SERVICE_NAME from env\n")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("Can not load PORT name from env\n")
		os.Exit(1);
	}

	postgresUri := os.Getenv("POSTGRES_URI")
	if postgresUri == ""{
		log.Fatal("Can not load POSTGRES_URI name from env\n")
		os.Exit(1);
	}

	portIntValue,err:= strconv.Atoi(port)
	if err!=nil{
		log.Fatal("Error converting port value")
		os.Exit(1)
	}

	config = ConfigOptions{
		Version: version,
		ServiceName: serviceName,
		Port: portIntValue,
		PostgresUri: postgresUri,
	}
}

func GetEnv()ConfigOptions{
	loadConfig()
	return config
}