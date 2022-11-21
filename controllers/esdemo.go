package controllers

import (
	"encoding/json"
	"fmt"
	"wsybapi/services/es"

	"github.com/astaxie/beego"
)

type EsDemoController struct {
	beego.Controller
}

// @router /es/add [*]
func (this *EsDemoController) Add() {
	body := map[string]interface{}{
		"id":    1,
		"title": "张三",
	}
	es.EsAdd("fyouku_demo", "user-1", body)
	this.Ctx.WriteString("add")
}

// @router /es/edit [*]
func (this *EsDemoController) Edit() {
	body := map[string]interface{}{
		"id":    1,
		"title": "李四",
	}
	es.EsEdit("fyouku_demo", "user-1", body)
	this.Ctx.WriteString("edit")
}

// @router /es/delete [*]
func (this *EsDemoController) Delete() {
	es.EsDelete("fyouku_demo", "user-1")
	this.Ctx.WriteString("delete")
}

// @router /es/search [*]
func (this *EsDemoController) Search() {
	sort := []map[string]string{map[string]string{"id": "desc"}}

	query := map[string]interface{}{
		"bool": map[string]interface{}{
			"must": map[string]interface{}{
				"term": map[string]interface{}{
					"id": 1,
				},
			},
		},
	}
	res := es.EsSearch("fyouku_demo", query, 0, 10, sort)
	total := res.Total
	// 返回字段根据业务转换
	var resData []ResData
	for _, v := range res.Hits {
		var data ResData
		err := json.Unmarshal([]byte(v.Source), &data)
		if err != nil {
			fmt.Println(err)
		}
		resData = append(resData, data)
	}
	fmt.Println(total)
	fmt.Println(resData)
	this.Ctx.WriteString("search")
}

type ResData struct {
	Title string
	Id    int
}
