package client_test

import (
	"bytes"
	"github.com/mk29142/pooled-reverse-geocode/client"
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"net/http"

	"github.com/mk29142/pooled-reverse-geocode/client/internal"
)

var _ = Describe("Client", func() {

	var (
		apiToken   string
		clientFake *internal.FakeClient
		service    client.Client
	)

	BeforeEach(func() {
		clientFake = new(internal.FakeClient)
		clientFake.DoReturns(&http.Response{
			Body:       ioutil.NopCloser(new(bytes.Buffer)),
			StatusCode: http.StatusNoContent,
		}, nil)
		apiToken = "super-secret-header-value"
		service = client.New(apiToken, clientFake)
	})


})
