package execute

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"io/ioutil"
	"os"
	"../service/awsService"
	"fmt"
)

func createFunction(zipFileName string, functionName string, handler string, resourceArn string, runtime string) {
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

	contents, err := ioutil.ReadFile(zipFileName + ".zip")

	if err != nil {
		fmt.Println("Could not read " + zipFileName + ".zip")
		os.Exit(0)
	}

	// localのzipファイルをソースコードにする場合はS3はコメントアウト
	createCode := &lambda.FunctionCode{
		//S3Bucket:        aws.String(bucketName),
		//S3Key:           aws.String(zipFileName),
		//S3ObjectVersion: aws.String("1"),
		ZipFile:         contents,
	}

	createArgs := &lambda.CreateFunctionInput{
		Code:         createCode,
		FunctionName: aws.String(functionName),
		Handler:      aws.String(handler),
		Role:         aws.String(resourceArn),
		Runtime:      aws.String(runtime),
	}

	result, err := svc.CreateFunction(createArgs)
	if err != nil {
		fmt.Println("Cannot create function: " + err.Error())
	} else {
		fmt.Println(result)
	}
}

func main() {
	// Lambda FunctionのListと取得する
	lambdaListResult := awsService.GetLambdaList()

	// 作成するLambda Functionが存在しているかチェックする
	isLambdaFunction := func(createLambdaFunctionName string) bool {
		for _, lambdaFunctionName := range lambdaListResult.Functions{
			if *lambdaFunctionName.FunctionName == createLambdaFunctionName {
				return true
			}
		}
		return false
	}

	zipFilePath := "./test.py"
	functionName := "goLangLambda"
	handler := "lambda_function.lambda_handler"
	resourceArn := "arn:aws:iam::932446063073:role/service-role/executeSlackLambda"
	runtime := "python3.6"

	if isLambdaFunction(functionName) {
		fmt.Println("lambda function you want to create is already exist.")
		return
	}

	createFunction(zipFilePath, functionName, handler, resourceArn, runtime)
}

