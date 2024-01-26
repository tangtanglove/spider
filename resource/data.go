package resource

import (
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-lite/action"
	"github.com/quarkcloudio/quark-lite/model"
)

type Data struct {
	resource.Template
}

// 初始化
func (p *Data) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "数据"

	// 模型
	p.Model = &model.Data{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Data) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("realname", "姓名"),

		field.Text("email", "邮箱"),

		field.Text("company", "单位"),

		field.Datetime("created_at", "创建时间"),
	}
}

// 搜索
func (p *Data) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
	}
}

// 行为
func (p *Data) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		action.GetData(),
	}
}
