package framework

// IStatelessElement 无状态Element接口
type IStatelessElement interface {
	IComponentElement
}

// StatelessElement 无状态Element类型
type StatelessElement struct {
	ComponentElement
}

// Mount 挂载Element对象
func (o *StatelessElement) Mount() {
	o.ComponentElement.Mount()
}

// Unmount 卸载Element对象
func (o *StatelessElement) Unmount() {
	o.ComponentElement.Mount()
}

// Update 更新widget
func (o *StatelessElement) Update(widget IWidget) {
	o.ComponentElement.Update(widget)
}

// MarkNeedsBuild 标记需要重构
func (o *StatelessElement) MarkNeedsBuild() {
	o.ComponentElement.MarkNeedsBuild()
}

// Build 构建Widget对象
func (*StatelessElement) Build() IWidget {
	return nil
}
