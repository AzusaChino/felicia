package main

import (
	"log"

	"github.com/azusachino/felicia/aws"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("fail to read environment variables")
	}

	aws.InitS3()
}

func main() {
	aws.ListS3Files()
}
