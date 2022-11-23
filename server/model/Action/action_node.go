// 自动生成模板Trigger
package Action

// Trigger 结构体
type ActionNode struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	Data       interface{} `json:"data"`
	Label      string      `json:"label"`
	ActionName string      `json:"actionName"`
}
