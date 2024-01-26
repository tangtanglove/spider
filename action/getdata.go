package action

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"gorm.io/gorm"
)

type GetDataAction struct {
	actions.Action
}

// 获取数据
func GetData(options ...interface{}) *GetDataAction {
	action := &GetDataAction{}

	action.Name = "获取数据"
	if len(options) == 1 {
		action.Name = options[0].(string)
	}

	return action
}

// 初始化
func (p *GetDataAction) Init(ctx *builder.Context) interface{} {

	// 设置按钮类型,primary | ghost | dashed | link | text | default
	p.Type = "primary"

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 在表格行内展示
	p.SetOnlyOnIndex(true)

	return p
}

// 执行行为句柄
func (p *GetDataAction) Handle(ctx *builder.Context, query *gorm.DB) error {

	var (
		url = "http://school.freekaoyan.com/hebei/hebmu/daoshi/index.shtml"
	)

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("#main > div > div > ul > li > h5 > a", func(e *colly.HTMLElement) {
		// e.ForEach("a", func(i int, element *colly.HTMLElement) {
		// 	fmt.Printf("第%d超链接\t %s :: %s\n", i, element.Text, element.Attr("href"))
		// })

		fmt.Println(e.Attr("href"))

		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println(string(r.Body))
	})

	c.Visit(url)

	return ctx.JSON(200, message.Success("操作成功"))
}
