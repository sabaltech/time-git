package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/aws/aws-sdk-go/aws/session"
)

func scanToday() {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "sabal"),
	})

	svc := dynamodb.New(sess)

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames: map[string]*string{
			"#DS": aws.String("Description"),
			"#HR": aws.String("Hours"),
			"#MN": aws.String("Minutes"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				S: aws.String(time.Now().Format("2006-01-02")),
			},
		},
		FilterExpression:     aws.String("WorkDate = :a"),
		ProjectionExpression: aws.String("#DS, #HR, #MN"),
		TableName:            aws.String("TimeSheetEntries"),
	}

	result, err := svc.Scan(input)

	if err != nil {
		fmt.Println(err)
	}

	for _, row := range result.Items {
		fmt.Printf("%2sh %2sm %s\n", *row["Hours"].N, *row["Minutes"].N, *row["Description"].S)
	}
}

func save(desc string, hours int, minutes int) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "sabal"),
	})

	svc := dynamodb.New(sess)
	uuid, _ := uuid.NewRandom()

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Uuid": {
				S: aws.String(uuid.String()),
			},
			"UserId": {
				S: aws.String("jllombart"),
			},
			"WorkDate": {
				S: aws.String(time.Now().Format("2006-01-02")),
			},
			"Description": {
				S: aws.String(desc),
			},
			"Hours": {
				N: aws.String(strconv.Itoa(hours)),
			},
			"Minutes": {
				N: aws.String(strconv.Itoa(minutes)),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("TimeSheetEntries"),
	}

	result, _ := svc.PutItem(input)

	fmt.Println(result)
}
