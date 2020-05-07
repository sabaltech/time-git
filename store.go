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
			"Date": {
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
