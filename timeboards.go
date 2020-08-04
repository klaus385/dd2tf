//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"github.com/zorkian/go-datadog-api"
)

type Timeboard struct {
}

func (d Timeboard) getElement(client datadog.Client, id interface{}) (interface{}, error) {
	dash, err := client.GetDashboard(id)
	return dash, err
}

func (d Timeboard) getAsset() string {
	return "tmpl/timeboard.tmpl"
}

func (d Timeboard) getName() string {
	return "timeboards"
}

func (d Timeboard) String() string {
	return d.getName()
}

func (d Timeboard) getAllElements(client datadog.Client) ([]Item, error) {
	var ids []Item
	dashboards, err := client.GetDashboards()
	if err != nil {
		return ids, err
	}
	for _, elem := range dashboards {
		ids = append(ids, Item{id: *elem.Id, d: Timeboard{}})
	}
	return ids, nil
}
