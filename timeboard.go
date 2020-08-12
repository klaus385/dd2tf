//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"github.com/zorkian/go-datadog-api"
)

type Timeboard struct {
}

func (t Timeboard) getElement(client datadog.Client, id interface{}) (interface{}, error) {
	time, err := client.GetDashboard(id)
	return time, err
}

func (t Timeboard) getAsset() string {
	return "tmpl/timeboard.tmpl"
}

func (t Timeboard) getName() string {
	return "timeboards"
}

func (t Timeboard) String() string {
	return t.getName()
}

func (t Timeboard) getAllElements(client datadog.Client) ([]Item, error) {
	var ids []Item
	timeboards, err := client.GetDashboards()
	if err != nil {
		return ids, err
	}
	for _, elem := range timeboards {
		ids = append(ids, Item{id: *elem.Id, d: Timeboard{}})
	}
	return ids, nil
}
