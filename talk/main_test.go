package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/nlopes/slack/slackevents"
)

func Test_getRandomIndex(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRandomIndex(tt.args.max); got != tt.want {
				t.Errorf("getRandomIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMccallVoice(t *testing.T) {
	type args struct {
		voices []string
		index  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMccallVoice(tt.args.voices, tt.args.index); got != tt.want {
				t.Errorf("getMccallVoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getResponses(t *testing.T) {
	type args struct {
		sentiment string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getResponses(tt.args.sentiment); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResponses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sendToSlack(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sendToSlack(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("sendToSlack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getAPIEvents(t *testing.T) {
	type args struct {
		requestBody string
	}
	tests := []struct {
		name    string
		args    args
		want    slackevents.EventsAPIEvent
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAPIEvents(tt.args.requestBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAPIEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAPIEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getChallengeResponse(t *testing.T) {
	type args struct {
		requestBody string
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getChallengeResponse(tt.args.requestBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("getChallengeResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getChallengeResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMentionEventResponce(t *testing.T) {
	type args struct {
		mentionEvent *slackevents.AppMentionEvent
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMentionEventResponce(tt.args.mentionEvent)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMentionEventResponce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMentionEventResponce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleRequest(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleRequest(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
