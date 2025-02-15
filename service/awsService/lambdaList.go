package awsService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"os"
)

// Lists all of your Lambda functions in us-west-2
func GetLambdaList() *lambda.ListFunctionsOutput {
	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	profile := "home"

	// Create Lambda service client
	svc := lambda.New(sess, &aws.Config{
		Credentials: credentials.NewSharedCredentials("", profile),
		Region: aws.String("ap-northeast-1"),
	})

	result, err := svc.ListFunctions(nil)

	if err != nil {
		fmt.Println("Cannot list functions")
		os.Exit(0)
	}

	//fmt.Println("Functions:")

	//for _, f := range result.Functions {
	//	fmt.Println("Name:        " + aws.StringValue(f.FunctionName))
	//	fmt.Println("Description: " + aws.StringValue(f.Description))
	//	fmt.Println("")
	//}

	return result
}