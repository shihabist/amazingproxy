package service_test

import (
	"proxy/model"
	"proxy/repository"
	"proxy/service"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockAllowedUri struct {
	uri     string
	method  string
	expects bool
}

func TestValidator(t *testing.T) {
	var cases = []MockAllowedUri{
		{"company", "GET", true},
		{"company/", "GET", true},
		{"company", "POST", true},
		{"company/", "POST", true},

		{"company/abc78dsds", "GET", true},
		{"company/account", "GET", true},
		{"company/account/", "GET", true},
		{"company/account", "POST", false},
		{"company/accounts", "GET", false},
		{"company/randomText?foo=bar", "GET", false},
		{"company?", "GET", true},
		{"company?key=value&more=value", "GET", true},

		{"account/blocked", "GET", false},
		{"account/pewew56ewew", "GET", true},
		{"account/kdjff235fss/user", "GET", true},

		{"tenant/aaaa8989adf", "GET", false},
		{"tenant/aaaa8989adf?first=value&sec=value", "GET", false},
		{"tenant/account/vfdfd4554fdfd", "GET", false},
		{"tenant/account/blocked", "GET", true},
		{"tenant/account/blocked?first=value&sec=value&third=value", "GET", true},

		{"rayhan10101mahmud", "GET", true},
		{"rayhan10101mahmud?param", "GET", true},
		{"rayhan10101mahmud", "POST", false},
		{"rayhan10101mahmud?param", "POST", false},
	}

	var (
		fileRepo     = repository.NewFileRepository()
		proxyService = service.NewProxyService(fileRepo)
	)

	for _, tc := range cases {
		allowedUriTemp := &model.AllowedUri{
			Uri:    tc.uri,
			Method: tc.method,
		}
		require.Equal(t, tc.expects, proxyService.Validate(allowedUriTemp), "Test failed for URI: %s", tc.uri)
	}
}
