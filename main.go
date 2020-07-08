package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	l "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type requestParam struct {
	instrument string `json:"instrument"`
	musicName  string `json:"musicName"`
}

type MusicScore struct {
	Id			int    `json:"id"`
	ServiceName string `json:"serviceName"`
	MusicName   string `json:"musicName"`
	Composer    string `json:"composer"`
	Price       int    `json:"price"`
	Url         string `json:"url"`
	Instrument  string `json:"instrument"`
	ServiceId   string `json:"serviceId"`
	Difficulty  string `json:"difficulty"`
	CreatedAt   string `json:"createdAt"`

}

func main() {
	l.Start(output)
}


func output() (MusicScore, error) {

	fmt.Println("input 開始")
	svc := lambda.New(session.New())
	abc := requestParam{"altoSaxophone", "紅蓮"}
	jsonBytes, _ := json.Marshal(abc)
	input := &lambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("TARGET_ARN")),
		Payload:        jsonBytes,
	}
	resp, err := svc.Invoke(input)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(resp)
	fmt.Println("↓ExecutedVersion~~~~~~~~~~~~~~~~~~~")
	fmt.Println(resp.ExecutedVersion)
	fmt.Println("↓FunctionError~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(resp.FunctionError)
	fmt.Println("↓StatusCode~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(resp.StatusCode)
	fmt.Println("↓LogResult~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(resp.LogResult)
	fmt.Println("↓Payload~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(string(resp.Payload))
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~end")

	var mc MusicScore
	if err := json.Unmarshal(resp.Payload, &mc); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(mc.MusicName)
	return mc, err

	//x, n := binary.Varint(resp.Payload)
	//if n != len(resp.Payload) {
	//	fmt.Println("Varint did not consume all of in")
	//}
	//fmt.Println(len(resp.Payload))

	//for _, mc := range resp.Payload {
	//	fmt.Println("#############")
	//	fmt.Println(mc)
	//
	//}
	//var respvalue returns
	//err = json.Unmarshal(resp.Payload, &respvalue)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(respvalue.Str)
	//fmt.Println(respvalue.Slc)
	//fmt.Println(respvalue.Int)
	//


	//
	//fmt.Println("input 終了")
	//fmt.Println(resp)
	//fmt.Println(string(resp.Payload))
	//fmt.Println(resp.FunctionError)
	//fmt.Println(resp.LogResult)
	//fmt.Println(resp.ExecutedVersion)
	//fmt.Println(resp.StatusCode)
}
