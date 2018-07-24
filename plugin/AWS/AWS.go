package AWS
/*
A Class to encapsulate the lambda SDK.
 */

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)

 type LambdaClient struct{
	Svc *lambda.Lambda
 }

 func NewLambdaClient() LambdaClient{
	var client LambdaClient

 	ses := session.Must(session.NewSessionWithOptions(session.Options{
		 SharedConfigState: session.SharedConfigEnable,
	 }))

 	client.Svc = lambda.New(ses, &aws.Config{Region: aws.String("us-east-1")})
	return client
 }

func (l *LambdaClient) Invoke(resource string, input []byte) ([]byte, string, error){
	result, err :=  l.Svc.Invoke(&lambda.InvokeInput{FunctionName:aws.String(resource),Payload:input})
	return result.Payload, *result.LogResult, err
}