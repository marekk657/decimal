package decimal

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// MarshalDynamoDBAttributeValue is implementation of dynamodbattribute.Marshaler
func (d Decimal) MarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	n := d.String()
	av.N = &n
	return nil
}

// UnmarshalDynamoDBAttributeValue is implementation of dynamodbattribute.Unmarshaler
func (d *Decimal) UnmarshalDynamoDBAttributeValue(av *dynamodb.AttributeValue) error {
	if av.N == nil {
		return nil
	}

	dec, err := NewFromString(*av.N)
	if err != nil {
		return err
	}
	*d = dec
	return nil
}
