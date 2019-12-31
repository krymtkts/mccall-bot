package main

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/json"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

var requestURI = os.Getenv("incoming_webhook_uri")
var mccallVoices = []string{
	"力になりたくても、やる気がなければ無理だ。",
	"完璧さよりも進歩だ。",
	"体と知力と心だ。",
	"直感だよ。",
	"老人は老人で魚は魚だ。自分以外のものにはなれない、何があっても。",
	"自分を疑えば失敗する。",
	"なれるよ。そうなりたいと、望むなら、何にでもなれる。",
	"世界を変えろ。",
	"ああ、ちょっとドジを踏んでね。",
	"ベストを尽くせ。",
	"警官は倫理の象徴だ、クズが。保護と奉仕、法の遵守、正義、忘れたか。",
	"ある日誰かがひどいことをする。被害者とは他人だが、見過ごせない。なぜなら、力になってやれるからだ。",
	"正義を行うんだ。正義を行え。正義を行え。いい警官たちのために。",
	"組織とビジネスは潰す、ひとつずつ、一ドルずつ、一人ずつ。",
	"雨乞いをするならぬかるみも覚悟しろ。",
	"さっきその目に何が見えるか聞いたな？私の目には何が見える。",
	"自分で紬げる。",
}

func getRandomIndex() int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return rand.Intn(len(mccallVoices))
}

func getMccallVoice(index int) string {
	return mccallVoices[index]
}

func sendToSlack(message string) error {
	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"text": message,
	})
	json.HTMLEscape(&buf, body)

	req, err := http.NewRequest(http.MethodPost, requestURI, &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}

type MccallEventBody struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type MccallEvent struct {
	Body string `json:"body"`
	User string `json:"user"`
	Text string `json:"text"`
}

// HandleRequest is our lambda handler invoked by the `lambda.Start` function call
func HandleRequest(ctx context.Context, event MccallEvent) (Response, error) {
	log.Println("start")

	var challenge = ""
	log.Printf("event: %+v\n", event)
	if event.Body != "" {
		var eventBody MccallEventBody
		json.Unmarshal([]byte(event.Body), &eventBody)
		log.Printf("eventBody: %+v\n", eventBody)
		if &eventBody != nil {
			challenge = eventBody.Challenge
			log.Println(challenge)
		}
	}

	var message = getMccallVoice(getRandomIndex())
	log.Println(message)
	sendToSlack(message)

	var body, _ = json.Marshal(map[string]interface{}{
		"challenge": challenge,
	})
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "world-handler",
		},
	}
	log.Printf("resp: %+v\n", resp)

	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
