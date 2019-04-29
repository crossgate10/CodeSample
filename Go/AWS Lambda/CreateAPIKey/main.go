package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appsync"
	"log"
	"os"
	"time"
)

func main(){
	// On Lambda
	lambda.Start(Handler)

	// On Local
	// CreateAPIKey()
}

func Handler(ctx context.Context) {
	CreateAPIKey()
	ctx.Done()
}

func CreateAPIKey() {
	//akid := os.Getenv("AKID")
	//asak := os.Getenv("SECRET_KEY")
	apiId := os.Getenv("API_ID")
	region := os.Getenv("REGION")

	sess := session.Must(session.NewSession())

	svc := appsync.New(sess, &aws.Config{
		Region: aws.String(region),
		//// use credentials on local enviroment. In contrast, use IAM on lambda.
		//Credentials: credentials.NewStaticCredentials(akid, asak, ""),
	})

	resp, err := svc.CreateApiKey(&appsync.CreateApiKeyInput{
		ApiId: aws.String(apiId),
		Expires: aws.Int64(time.Now().Add(10*time.Second).AddDate(0,0,2).Unix()),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			log.Fatalf("canceled due to timeout, %v\n", err)
		} else {
			log.Fatalf("failed to create api key, %v\n", err)
		}
	}

	log.Printf("%+v\n", resp)
}