package main

import (
	"bytes"
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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/nlopes/slack/slackevents"
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

type simpleBody struct {
	Token string `json:"token"`
}

func getAPIEvents(requestBody string) (slackevents.EventsAPIEvent, error) {
	log.Printf("requestBody: %+v\n", requestBody)

	var verification simpleBody
	json.Unmarshal([]byte(requestBody), &verification)
	log.Printf("verification: %+v\n", verification)

	eventsAPIEvent, err := slackevents.ParseEvent(
		json.RawMessage(requestBody),
		slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: verification.Token}),
	)
	return eventsAPIEvent, err
}

func getChallengeResponse(requestBody string) (Response, error) {
	var r *slackevents.ChallengeResponse
	err := json.Unmarshal([]byte(requestBody), &r)
	if err != nil {
		log.Print(err)
		return Response{}, err
	}
	return Response{
		StatusCode: 200,
		Body:       r.Challenge,
	}, nil
}

func getMentionEventResponce(mentionEvent *slackevents.AppMentionEvent) (Response, error) {
	log.Printf("body.event.text: %+v\n", mentionEvent.Text)
	client := comprehend.New(session.New(), aws.NewConfig().WithRegion("ap-southeast-1"))
	param := comprehend.DetectSentimentInput{}
	param.SetLanguageCode("ja")
	param.SetText(mentionEvent.Text)
	log.Printf("validate sentiment params: %+v\n", param.Validate())
	output, err := client.DetectSentiment(&param)
	if err != nil {
		log.Printf("detected sentiment failed: %+v\n", err)
	} else {
		log.Printf("sentiment: %+v\n", *(output.Sentiment))
		log.Printf("score: %+v\n", output.SentimentScore)
	}

	message := getMccallVoice(getRandomIndex())
	log.Println(message)
	sendToSlack(message)

	return Response{
		StatusCode: 200,
	}, nil
}

// HandleRequest is our lambda handler invoked by the `lambda.Start` function call
func HandleRequest(request events.APIGatewayProxyRequest) (Response, error) {
	log.Println("start")

	eventsAPIEvent, err := getAPIEvents(request.Body)
	if err != nil {
		log.Print(err)
		return Response{}, err
	}

	log.Printf("eventsAPIEvent: %+v\n", eventsAPIEvent)
	switch eventsAPIEvent.Type {
	case slackevents.URLVerification:
		return getChallengeResponse(request.Body)
	case slackevents.CallbackEvent:
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			return getMentionEventResponce(ev)
		default:
			log.Printf("unsupported event: %+v\n", ev)
		}
	default:
		log.Printf("unsupported type: %+v\n", eventsAPIEvent)
	}
	log.Println("no effect.")
	return Response{
		StatusCode: 400,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
