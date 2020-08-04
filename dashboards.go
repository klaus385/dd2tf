//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"errors"
	"github.com/zorkian/go-datadog-api"
)

type Dashboard struct {
}

func (d Dashboard) getElement(client datadog.Client, id interface{}) (interface{}, error) {
	var idStr string
	switch v := id.(type) {
	case string:
		idStr = v
	default:
		return "", errors.New("unsupported id type, should be string or int")
	}
	dash, err := client.GetBoard(idStr)
	return dash, err
}

func (d Dashboard) getAsset() string {
	return "tmpl/dashboard.tmpl"
}

func (d Dashboard) getName() string {
	return "dashboards"
}

func (d Dashboard) String() string {
	return d.getName()
}

func (d Dashboard) getAllElements(client datadog.Client) ([]Item, error) {
	var ids []Item
	dashboards, err := client.GetBoards()
	if err != nil {
		return ids, err
	}
	for _, elem := range dashboards {
		ids = append(ids, Item{id: *elem.Id, d: Dashboard{}})
	}
	return ids, nil
}
