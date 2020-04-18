package framework

// IComponentElement 组合Element接口
type IComponentElement interface {
	IElement

	MarkNeedsBuild()
	Build() IWidget
}

// ComponentElement 组合Element类型
type ComponentElement struct {
	Element

	child IElement
	dirty bool
}

// Mount 挂载Element对象
func (o *ComponentElement) Mount() {
	o.Element.Mount()
	o.ReBuild()
}

// Unmount 卸载Element对象
func (o *ComponentElement) Unmount() {
	o.Element.Unmount()
}

// AttachRenderObject 遍历绑定Render对象
func (o *ComponentElement) AttachRenderObject() {
    o.VisitChildren(func(child IElement) {
		child.AttachRenderObject()
	})
}

// DetachRenderObject 遍历解绑Render对象
func (o *ComponentElement) DetachRenderObject() {
    o.VisitChildren(func(child IElement) {
		child.DetachRenderObject()
	})
}

// VisitChildren 访问子对象
func (o *ComponentElement) VisitChildren(visitor func(IElement)) {
    if o.child != nil {
        visitor(o.child);
    }
}

// Update 更新widget
func (o *ComponentElement) Update(widget IWidget) {
	o.Element.Update(widget)
}

// MarkNeedsBuild 标记需要重构
func (o *ComponentElement) MarkNeedsBuild() {
	o.dirty = true
}

// ReBuild 标记需要重构
func (o *ComponentElement) ReBuild() {
	widget := o.Build()
	o.child = o.UpdateChild(o.child, widget)
}

// Build 构建Widget对象
func (*ComponentElement) Build() IWidget {
	return nil
}
