package framework

// IState 状态接口
type IState interface {
	SetState(func())
	Build() IWidget
}

// State 状态类型
type State struct {
	widget  IStatefulWidget
	element IStatefulElement
}

// SetState 更新状态
func (o *State) SetState(f func()) {
	if f != nil {
		f()
	}
	o.element.MarkNeedsBuild()
}

// Build 构建Widget
func (o *State) Build() IWidget {
	return nil
}