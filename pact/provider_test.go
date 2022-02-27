package pact

import (
	"backend_todo/server"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	svr := server.NewServer()
	go svr.StartServer(port)

	pact := dsl.Pact{
		Host:     "127.0.0.1",
		Consumer: "frontend",
		Provider: "backend",

		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port),
		PactURLs: []string{
			"/Users/mehmetb.akin/Desktop/frontend/pact/pacts/frontend-backend.json",
		},
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponses), "pact tests run")
}
