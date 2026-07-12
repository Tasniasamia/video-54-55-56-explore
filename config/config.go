package config;

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
)

type Config struct {
	HttpPort    int;
	Version     string;
	ServiceName string;
	JwtSecret   string;
}


var cfg Config;

func LoadConfig() {
	
	godotenv.Load();

	version :=os.Getenv("VERSION");
	if(version == ""){
    fmt.Println("Version is Required");
	os.Exit(1);
	}


	httpPortStr :=os.Getenv("HTTP_PORT");
	if(httpPortStr == ""){
    fmt.Println("HTTP_PORT is Required");
	os.Exit(1);
	}

	serviceName :=os.Getenv("SERVICE_NAME");
	if(serviceName == ""){
    fmt.Println("SERVICE_NAME is Required");
	os.Exit(1);
	}

   httpPort,err :=strconv.Atoi(httpPortStr);
   if(err != nil){
    fmt.Println("Invalid HTTP_PORT");
	os.Exit(1);
   }

   jwtSecret :=os.Getenv("JWT_SECRET");
   if(jwtSecret == ""){
    fmt.Println("JWT_SECRET is Required");
	os.Exit(1);
   }

	cfg=Config{
		HttpPort:httpPort,
        Version:version,
		ServiceName:serviceName,
		JwtSecret:jwtSecret,
    };
	
}


func GetConfig() Config {

	LoadConfig()
	return cfg;
}

