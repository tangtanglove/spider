package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/gocolly/colly/v2"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/component/message"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource/actions"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-lite/model"
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
		baseUrl = "http://school.freekaoyan.com/hebei/hebmu/daoshi"
	)

	for i := 1; i <= 50; i++ {
		url := ""
		if i == 1 {
			url = baseUrl + "/index.shtml"
		} else {
			url = baseUrl + "/index_" + strconv.Itoa(i) + ".shtml"
		}

		fmt.Println("Get Page " + strconv.Itoa(i))
		p.GetPageData(url)
	}

	return ctx.JSON(200, message.Success("操作成功"))
}

// 获取一个页面的数据
func (p *GetDataAction) GetPageData(url string) {
	var (
		titles   = []string{}
		contents = []string{}
	)

	c := colly.NewCollector(
		colly.MaxDepth(2), // 定义访问深度
		colly.AllowedDomains("school.freekaoyan.com"), // 访问域名
	)

	// 构建访问规则
	c.OnHTML("#main > div > div > ul > li > h5 > a", func(e *colly.HTMLElement) {
		// utf8Text := p.ConvertToString(e.Text, "gbk", "utf-8")
		// fmt.Printf("%s => %s\n", utf8Text, e.Attr("href"))
		e.Request.Visit(e.Attr("href"))
	})

	// 获取详情页标题
	c.OnHTML("#main > div.container.box.content-box > h2", func(e *colly.HTMLElement) {
		utf8Text := p.ConvertToString(e.Text, "gbk", "utf-8")
		titles = append(titles, utf8Text)

		// fmt.Printf("%s\n", utf8Text)
	})

	// 获取详情页内容
	c.OnHTML("#main > div.container.box.content-box > div.content", func(e *colly.HTMLElement) {
		utf8Text := p.ConvertToString(e.Text, "gbk", "utf-8")
		contents = append(contents, utf8Text)

		// fmt.Printf("%s\n", utf8Text)
	})

	c.Visit(url)

	for k, v := range titles {
		realname := ""
		company := ""
		email := ""
		title := v
		content := contents[k]

		// 提取姓名
		getTitles := strings.Split(title, "-")
		if len(getTitles) > 1 {
			realname = getTitles[1]
		}

		// 提取工作单位
		contents1 := strings.Split(content, "工作单位：")
		if len(contents1) > 1 {
			company = contents1[1]
			contents2 := strings.Split(company, "所在科室：")
			if len(contents2) > 1 {
				company = contents2[0]
			}
		}

		// 提取email
		contents3 := strings.Split(content, "E-mail：")
		if len(contents3) > 1 {
			email = contents3[1]
			contents4 := strings.Split(email, "个人简介：")
			if len(contents4) > 1 {
				email = contents4[0]
			}
		}

		if company == "" && email == "" {
			break
		}

		(&model.Data{}).Insert(&model.Data{
			Realname: realname,
			Email:    email,
			Company:  company,
		})
	}
}

// 字符集转换
func (p *GetDataAction) ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)

	return result
}
