package config;

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
	"strconv"
)

type Dbconfig struct{
Dbuser string `db:"user"`
Password string `db:"password"`
Host string `db:"host"`
Port int `db:"port"`
Dbname string `db:"dbname"`
EnableSSLMode bool `db:"enablesslmode"`
}

type Config struct {
	HttpPort    int;
	Version     string;
	ServiceName string;
	JwtSecret   string;
	Dbconfig *Dbconfig;
}
// DBNAME=ecommerce
// ENABLESSLMODE=disable
// }

// type Config struct {
// 	HttpPort    int;
// 	Version     string;
// 	ServiceName string;
// 	JwtSecret   string;
// }

// =postgres
// PASSWORD=123456
// HOST=localhost 
// PORT=5432
// DBNAME=ecommerce
// ENABLESSLMODE=disable


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

   httpPort,err :=strconv.Atoi(httpPortStr);
   if(err != nil){
    fmt.Println("HTTP_PORT must be number");
	os.Exit(1);
   }


	serviceName :=os.Getenv("SERVICE_NAME");
	if(serviceName == ""){
    fmt.Println("SERVICE_NAME is Required");
	os.Exit(1);
	}

 

   jwtSecret :=os.Getenv("JWT_SECRET");
   if(jwtSecret == ""){
    fmt.Println("JWT_SECRET is Required");
	os.Exit(1);
   }

  dbUser :=os.Getenv("DB_USER");
   if(dbUser == ""){
    fmt.Println("DB_USER is Required");
	os.Exit(1);
   }

   dbPassword :=os.Getenv("PASSWORD");
   if(dbPassword == ""){
    fmt.Println("PASSWORD is Required");
	os.Exit(1);
   }

   dbHost :=os.Getenv("HOST");
   if(dbHost == ""){
    fmt.Println("HOST is Required");
	os.Exit(1);
   }

   dbPort :=os.Getenv("PORT");
   if(dbPort == ""){
    fmt.Println("PORT is Required");
	os.Exit(1);
   }
   
   dbPortInt,err :=strconv.Atoi(dbPort);
   if(err != nil){
    fmt.Println("PORT must be number");
	os.Exit(1);
   }

   dbName :=os.Getenv("DBNAME");
   if(dbName == ""){
    fmt.Println("DBNAME is Required");
	os.Exit(1);
   }

   enableSSLMode :=os.Getenv("ENABLESSLMODE");
   if(enableSSLMode == ""){
    fmt.Println("ENABLESSLMODE is Required");
	os.Exit(1);
   }
   enableSSLModeBool,err :=strconv.ParseBool(enableSSLMode);
   if(err != nil){
    fmt.Println("ENABLESSLMODE must be a boolean value");
	os.Exit(1);
   }

	cfg=Config{
		HttpPort:httpPort,
        Version:version,
		ServiceName:serviceName,
		JwtSecret:jwtSecret,
		Dbconfig:&Dbconfig{
			Dbuser:dbUser,
			Password:dbPassword,
			Host:dbHost,
			Port:dbPortInt,
			Dbname:dbName,
			EnableSSLMode:enableSSLModeBool,
		},
    };
	
}


func GetConfig() Config {

	LoadConfig()
	return cfg;
}

