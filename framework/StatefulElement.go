package framework

// IStatefulElement 有状态Element接口
type IStatefulElement interface {
	IComponentElement
}

// StatefulElement 有状态Element类型
type StatefulElement struct {
	ComponentElement

	state IState
}

// Mount 挂载Element对象
func (o *StatefulElement) Mount() {
	o.ComponentElement.Mount()
}

// Unmount 卸载Element对象
func (o *StatefulElement) Unmount() {
	o.ComponentElement.Mount()
}

// Update 更新widget
func (o *StatefulElement) Update(widget IWidget) {
	o.ComponentElement.Update(widget)
}

// MarkNeedsBuild 标记需要重构
func (o *StatefulElement) MarkNeedsBuild() {
	o.ComponentElement.MarkNeedsBuild()
}

// Build 构建Widget对象
func (o *StatefulElement) Build() IWidget {
	return o.state.Build()
}
