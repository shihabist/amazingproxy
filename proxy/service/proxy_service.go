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

var (
	proxyRepo repository.ProxyRepository
)

func NewProxyService(repo repository.ProxyRepository) ProxyService {
	proxyRepo = repo
	return &service{}
}

func (*service) Validate(uri *model.AllowedUri) bool {
	allowedUriList, err := proxyRepo.GetAllowedUriList()
	if err == nil {
		logger.Print("Allowed Uri loaded", "info")
	}

	rules, err := proxyRepo.GetRulesList()
	if err == nil {
		logger.Print("Allowed Uri rules loaded", "info")
	}
	allowedList := make([]*regexp.Regexp, len(allowedUriList))

	for key, uri := range allowedUriList {
		for rule, value := range rules {
			uri.Uri = strings.ReplaceAll(uri.Uri, "{"+rule+"}", value)
		}
		allowedList[key] = regexp.MustCompile(`^` + uri.Uri + `$`)
	}
	for _, u := range allowedList {
		if u.MatchString(strings.Trim(strings.ToLower(strings.Split(uri.Uri, "?")[0]), "/")) {
			return true
		}
	}
	logger.Print("Uri not allowed", "warning")
	return false
}
