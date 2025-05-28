package model

type Parcel struct {
	Receiver  string `json:"receiver" dynamodbav:"receiver"` // DynamoDB와 JSON 직렬화에 모두 적용
	Sender    string `json:"sender" dynamodbav:"sender"`
	Address   string `json:"address" dynamodbav:"address"`
	Status    string `json:"status" dynamodbav:"status"`
	CreatedAt int    `json:"created_at" dynamodbav:"created_at"`
}
