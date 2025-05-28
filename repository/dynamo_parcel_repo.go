package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"golang-coupang-backend.com/m/model"
)

const (
	tableName = "parcels"
)

type DynamoParcelRepository struct {
	Client    *dynamodb.Client
	TableName string
}

func NewDynamoParcelRepository(c *dynamodb.Client) *DynamoParcelRepository {
	return &DynamoParcelRepository{
		Client:    c,
		TableName: tableName,
	}
}

func (r *DynamoParcelRepository) Create(ctx context.Context, parcel model.Parcel) error {
	log.Printf("model.Parcel: %v\n", parcel)
	item, _ := attributevalue.MarshalMap(parcel)

	log.Printf("item receiver: %v", item["receiver"])
	log.Printf("item Receiver", item["Receiver"])
	log.Printf("item: %v", item)
	_, err := r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      item,
	})

	return err
}

//func (r *DynamoParcelRepository) Update(ctx context.Context, parcel model.Parcel) error {
//	// UpdateItem 입력 데이터 구성
//	updateExpression := "set sender = :sender, receiver = :receiver, address = :address, status = :status, created_at = :created_at"
//	expressionValues := map[string]types.AttributeValue{
//		":sender":     &types.AttributeValueMemberS{Value: parcel.Sender},
//		":receiver":   &types.AttributeValueMemberS{Value: parcel.Receiver},
//		":address":    &types.AttributeValueMemberS{Value: parcel.Address},
//		":status":     &types.AttributeValueMemberS{Value: parcel.Status},
//		":created_at": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", parcel.CreatedAt)},
//	}
//
//	// UpdateItem 실행
//	_, err := r.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
//		TableName:                 aws.String(r.TableName),
//		Key:                       map[string]types.AttributeValue{"id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", parcel.ID)}},
//		UpdateExpression:          aws.String(updateExpression),
//		ExpressionAttributeValues: expressionValues,
//	})
//	return err
//}

func (r *DynamoParcelRepository) Delete(ctx context.Context, id int) error {
	_, err := r.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", id)},
		},
	})
	return err
}

func (r *DynamoParcelRepository) GetAll(ctx context.Context) ([]model.Parcel, error) {
	output, err := r.Client.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String(r.TableName),
	})
	if err != nil {
		return nil, err
	}
	var parcels []model.Parcel
	for _, item := range output.Items {
		var parcel model.Parcel
		err := attributevalue.UnmarshalMap(item, &parcel)
		if err != nil {
			return nil, err
		}
		parcels = append(parcels, parcel)
	}
	return parcels, nil
}

func (r *DynamoParcelRepository) GetByID(ctx context.Context, id int) (model.Parcel, error) {
	output, err := r.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", id)},
		},
	})
	if err != nil {
		return model.Parcel{}, err
	}
	if output.Item == nil {
		return model.Parcel{}, fmt.Errorf("parcel not found with id %d", id)
	}
	var parcel model.Parcel
	err = attributevalue.UnmarshalMap(output.Item, &parcel)
	if err != nil {
		return model.Parcel{}, err
	}
	return parcel, nil
}
