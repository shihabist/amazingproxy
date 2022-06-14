package repository

import "proxy/model"

type ProxyRepository interface {
	GetAllowedUriList() ([]model.AllowedUri, error)
	GetRulesList() (map[string]string, error)
}
