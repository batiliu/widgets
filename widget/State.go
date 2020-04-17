package widget

// State 有状态声明对象
type State struct {
	widget  *StatefulWidget
	element *StatefulElement
}

// SetState 创建界面元素对象
func (o *State) SetState(f func()) {
	if f != nil {
		f()
	}
	o.element.MarkNeedsBuild()
}
