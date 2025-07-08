package runner

import (
	"os"
	"testing"

	"github.com/ducksify/cloudlist/pkg/schema"
)

func TestEnumerateResourcesWithEnvConfig(t *testing.T) {
	awsAccessKey := os.Getenv("TEST_AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("TEST_AWS_SECRET_KEY")
	cloudflareEmail := os.Getenv("TEST_CLOUDFLARE_EMAIL")
	cloudflareToken := os.Getenv("TEST_CLOUDFLARE_API_TOKEN")

	if awsAccessKey == "" || awsSecretKey == "" || cloudflareEmail == "" || cloudflareToken == "" {
		t.Skip("Environment variables for provider config not set")
	}

	providerConfig := schema.Options{
		{
			"provider":       "aws",
			"id":             "testaws",
			"aws_access_key": awsAccessKey,
			"aws_secret_key": awsSecretKey,
		},
		{
			"provider":  "cloudflare",
			"id":        "testcloudflare",
			"email":     cloudflareEmail,
			"api_token": cloudflareToken,
		},
	}

	r, err := NewWithConfig(providerConfig)
	if err != nil {
		t.Fatalf("Could not create runner: %v", err)
	}
	resources := r.EnumerateResources()
	if len(resources) == 0 {
		t.Errorf("Expected resources, got none")
	}
	for _, resource := range resources {
		t.Logf("%+v", resource)
	}
}
