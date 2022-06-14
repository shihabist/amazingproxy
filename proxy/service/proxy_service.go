package service

import (
	"regexp"
	"strings"

	logger "proxy/log"
	"proxy/model"
	"proxy/repository"
)

type ProxyService interface {
	Validate(proxy *model.AllowedUri) bool
}

type service struct{}
type AllowedListStruct struct {
	uri    *regexp.Regexp
	method string
}

var (
	proxyRepo repository.ProxyRepository
)

func NewProxyService(repo repository.ProxyRepository) ProxyService {
	proxyRepo = repo
	return &service{}
}

func (*service) Validate(incomingUri *model.AllowedUri) bool {
	allowedUriList, err := proxyRepo.GetAllowedUriList()
	if err == nil {
		logger.Print("Allowed Uri loaded", "info")
	}

	rules, err := proxyRepo.GetRulesList()
	if err == nil {
		logger.Print("Allowed Uri rules loaded", "info")
	}
	allowedList := make([]AllowedListStruct, len(allowedUriList))

	for key, uri := range allowedUriList {
		for rule, value := range rules {
			uri.Uri = strings.ReplaceAll(uri.Uri, "{"+rule+"}", value)
		}
		allowedList[key].uri = regexp.MustCompile(`^` + uri.Uri + `$`)
		allowedList[key].method = uri.Method
	}
	for _, u := range allowedList {
		if u.uri.MatchString(strings.Trim(strings.ToLower(strings.Split(incomingUri.Uri, "?")[0]), "/")) && u.method == incomingUri.Method {
			return true
		}
	}
	logger.Print("Uri not allowed", "warning")
	return false
}
