package main

import (
	"context"
	"os"
	"strings"
	"testing"

	"manochy-api/apps"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var testLambda *ginadapter.GinLambda

// setup ก่อน test
func init() {
	// set env สำหรับ test
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("MONGO_URI", "mongodb://localhost:27017") // mock/local

	r := apps.SetupRouter()
	testLambda = ginadapter.New(r)
}

func TestLogin(t *testing.T) {
	body := `{"username":"admin"}`

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/login",
		Body:       body,
	}

	resp, err := testLambda.ProxyWithContext(context.Background(), req)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}

	if !strings.Contains(resp.Body, "token") {
		t.Fatalf("expected token in response")
	}
}

func TestUnauthorized(t *testing.T) {
	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/members",
	}

	resp, _ := testLambda.ProxyWithContext(context.Background(), req)

	if resp.StatusCode != 401 {
		t.Fatalf("expected 401 got %d", resp.StatusCode)
	}
}

func TestChatbot(t *testing.T) {
	body := `{"question":"bank"}`

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Path:       "/chatbot",
		Headers: map[string]string{
			"Authorization": "Bearer test", // bypass (token fake)
		},
		Body: body,
	}

	resp, _ := testLambda.ProxyWithContext(context.Background(), req)

	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
}
