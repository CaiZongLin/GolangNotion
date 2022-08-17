package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jomei/notionapi"
)

const (
	apiKey = "abc123456" // Notion_API_KEY
)

func NotionInit() *notionapi.Client {
	client := notionapi.NewClient(apiKey)
	return client
}

func Update(client *notionapi.Client) {

	Time := time.Now().Format("2006-01-02T15:04:05Z07:00")

	NewTime, _ := time.Parse(time.RFC3339, Time)
	dataObj := notionapi.Date(NewTime)
	updateParams := &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"項目": notionapi.TitleProperty{
				//ID:   "title",
				//Type: "title",
				Title: []notionapi.RichText{
					{
						Type: "text",
						Text: notionapi.Text{
							Content: "測試用",
						},
						Annotations: &notionapi.Annotations{
							Color: "blue",
						},
					},
				},
			},
			"金額": notionapi.NumberProperty{
				//ID:     "number",
				//Type:   "number",
				Number: 99,
			},
			"分類": notionapi.SelectProperty{
				//ID:   "select",
				//Type: "select",
				Select: notionapi.Option{
					Name: "⛽️",
				},
			},
			"收入/支出": notionapi.SelectProperty{
				//ID:   "select",
				//Type: "select",
				Select: notionapi.Option{
					Name: "Spend",
				},
			},
			"個人/公司": notionapi.SelectProperty{
				//ID:   "select",
				//Type: "select",
				Select: notionapi.Option{
					Name: "Personal",
				},
			},
			"日期": notionapi.DateProperty{
				//ID:   "date",
				//Type: "date",
				Date: &notionapi.DateObject{
					Start: &dataObj,
				},
			},
			"支付方式": notionapi.SelectProperty{
				//ID:   "select",
				//Type: "select",
				Select: notionapi.Option{
					Name: "現金",
				},
			},
		},
	}

	update, err := client.Page.Update(context.Background(), "1359a759-2780-474c-ae19-c9326525cdd1", updateParams)
	if err != nil {
		log.Println(err)
	}
	log.Println("https://charming-punishment-e15.notion.site/" + update.URL[22:])
}

func get(client *notionapi.Client) {
	page, err := client.Page.Get(context.Background(), "927bfec8-0962-4663-9f8b-a594a303c798") //"4b960810-1a20-4686-a50e-51610cf1c32b")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(page.Properties["金額"])
}

func Create(client *notionapi.Client) {

	Time := time.Now().Format("2006-01-02T15:04:05Z07:00")

	NewTime, _ := time.Parse(time.RFC3339, Time)
	dataObj := notionapi.Date(NewTime)

	createParams := notionapi.PageCreateRequest{
		Properties: notionapi.Properties{
			"項目": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Type: "text",
						Text: notionapi.Text{
							Content: "測試新增",
						},
						Annotations: &notionapi.Annotations{
							Color: "blue",
						},
					},
				},
			},
			"金額": notionapi.NumberProperty{
				Number: 789,
			},
			"分類": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "早餐",
				},
			},
			"收入/支出": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "Spend",
				},
			},
			"個人/公司": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "Personal",
				},
			},
			"日期": notionapi.DateProperty{
				Date: &notionapi.DateObject{
					Start: &dataObj,
				},
			},
			"支付方式": notionapi.SelectProperty{
				Select: notionapi.Option{
					Name: "現金",
				},
			},
		},
		Parent: notionapi.Parent{ // create_database ID
			DatabaseID: "99d247ad-6401-4442-9856-87ae3c74df74",
		},
	}
	create, err := client.Page.Create(context.Background(), &createParams)
	if err != nil {
		log.Println(err)
	}
	log.Println("https://charming-punishment-e15.notion.site/" + create.URL[22:])
}
