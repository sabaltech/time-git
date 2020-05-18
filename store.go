package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws/session"
)

func getRecords(userID string, start time.Time, end time.Time) [][]string {
	// Headers mmddyy , actual hours, invoiced hours, description
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "sabal"),
	})

	dbClient := dynamodb.New(sess)

	input := &dynamodb.QueryInput{
		TableName: aws.String("TimeSheetEntries"),
		IndexName: aws.String("UserId-WorkDate-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"UserId": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userID),
					},
				},
			},
			"WorkDate": {
				ComparisonOperator: aws.String("BETWEEN"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(start.Format("2006-01-02")),
					},
					{
						S: aws.String(end.Format("2006-01-02")),
					},
				},
			},
		},
	}

	result, err := dbClient.Query(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	records := [][]string{}

	for _, row := range result.Items {
		hours, _ := strconv.Atoi(*row["Hours"].N)
		minutes, _ := strconv.Atoi(*row["Minutes"].N)

		actualHours := float64(hours) + (float64(minutes) / 60)

		workDate, _ := time.Parse("2006-01-02", *row["WorkDate"].S)

		record := []string{workDate.Format("01/02/06"), fmt.Sprintf("%.2f", actualHours), "", *row["Description"].S}
		records = append(records, record)
	}

	return records
}

func printTimesheet(userID string, start time.Time, end time.Time) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("", "sabal"),
	})

	dbClient := dynamodb.New(sess)

	input := &dynamodb.QueryInput{
		TableName: aws.String("TimeSheetEntries"),
		IndexName: aws.String("UserId-WorkDate-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"UserId": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userID),
					},
				},
			},
			"WorkDate": {
				ComparisonOperator: aws.String("BETWEEN"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(start.Format("2006-01-02")),
					},
					{
						S: aws.String(end.Format("2006-01-02")),
					},
				},
			},
		},
	}

	result, err := dbClient.Query(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range result.Items {
		fmt.Printf("%10s %2sh %2sm %s\n", *row["WorkDate"].S, *row["Hours"].N, *row["Minutes"].N, *row["Description"].S)
	}
}

func save(desc string, hours int, minutes int, workDate time.Time) {
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
				S: aws.String(workDate.Format("2006-01-02")),
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
