package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: "Alejandro",
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello Alejandro",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				request: events.APIGatewayProxyRequest{
					Body: "John Truman",
				},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello John Truman",
			},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				request: events.APIGatewayProxyRequest{},
			},
			want:    events.APIGatewayProxyResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
