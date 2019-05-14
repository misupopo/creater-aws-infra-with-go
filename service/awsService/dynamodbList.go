package awsService

import (
	//"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func GetDynamoDBList() *dynamodb.ListTablesOutput {
	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	profile := "home"

	svc := dynamodb.New(sess, &aws.Config{
		Credentials: credentials.NewSharedCredentials("", profile),
		Region:      aws.String("ap-northeast-1"),
	})

	input := &dynamodb.ListTablesInput{}

	result, err := svc.ListTables(input)

	if err != nil {
		if awsError, ok := err.(awserr.Error); ok {
			switch awsError.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, awsError.Error())
			default:
				fmt.Println(awsError.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	//for {
	//	// Get the list of tables
	//	result, err := svc.ListTables(input)
	//	if err != nil {
	//		if aerr, ok := err.(awserr.Error); ok {
	//			switch aerr.Code() {
	//			case dynamodb.ErrCodeInternalServerError:
	//				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
	//			default:
	//				fmt.Println(aerr.Error())
	//			}
	//		} else {
	//			// Print the error, cast err to awserr.Error to get the Code and
	//			// Message from an error.
	//			fmt.Println(err.Error())
	//		}
	//		return nil
	//	}

	return result
}

