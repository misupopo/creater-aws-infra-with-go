package execute

import (
	"../service/awsService"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createTable(tableName string) {
	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	profile := "home"

	ddb := dynamodb.New(sess, &aws.Config{
		Credentials: credentials.NewSharedCredentials("", profile),
		Region: aws.String("ap-northeast-1"),
	})

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),    // プライマリキー名
				AttributeType: aws.String("S"),     // データ型(String:S, Number:N, Binary:B の三種)
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),    // インデックス名
				KeyType:       aws.String("HASH"),  // インデックスの型(HASH または RANGE)
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Int64(1),   // 読み込みキャパシティーユニット（デフォルト：１）
			WriteCapacityUnits: aws.Int64(1),   // 書き込みキャパシティーユニット（デフォルト：１）
		},
		TableName: aws.String(tableName), // テーブル名
	}

	resp, err := ddb.CreateTable(params)

	if err != nil {
		fmt.Println(err.Error())    // エラー処理
	}

	fmt.Println(resp)
}

func main() {
	dynamodbListResult := awsService.GetDynamoDBList()

	isTableFunction := func(createTableName string) bool {
		for _, tableName := range dynamodbListResult.TableNames{
			if *tableName == createTableName {
				return true
			}
		}
		return false
	}

	tableName := "goLangTable"

	if isTableFunction(tableName) {
		fmt.Println("table that you want to create is already exist.")
		return
	} else {
		createTable(tableName)
	}

	fmt.Println(dynamodbListResult.TableNames)
}
