package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)


type MyEvent struct{
	name string `json:"what is you name?"`
	age int `json:"how old are you"`
}

type myresponce struct{
	Message string `json:"message"`
}
func HandleLambdaEvent(event MyEvent) (myresponce, error) {
	return myresponce{Message: fmt.Sprintf("%s is %d years old!", event.name,event.age)} ,nil
}
func main() {
	lambda.Start(HandleLambdaEvent)
}