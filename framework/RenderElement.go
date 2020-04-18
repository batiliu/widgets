package framework

// IRenderElement 渲染Element接口
type IRenderElement interface {
	IElement

	InsertChildRenderObject(IRenderObject)
	RemoveChildRenderObject(IRenderObject)
}

// RenderElement 渲染Element类型
type RenderElement struct {
	Element

	renderObject IRenderObject
	renderElement IRenderElement
}

// Mount 挂载Element对象
func (o *RenderElement) Mount() {
	o.Element.Mount()

	o.renderObject = o.widget.CreateRenderObject(o)
}

// Unmount 卸载Element对象
func (o *RenderElement) Unmount() {
	o.Element.Unmount()
}

// AttachRenderObject 遍历绑定Render对象
func (o *RenderElement) AttachRenderObject() {
	ancestor := o.Parent();
    for ancestor != nil {
		if _, ok := ancestor.(IRenderElement); !ok {
			ancestor = ancestor.Parent();
		}
	}
	
	if renderElement, ok := ancestor.(IRenderElement); ok {
		o.renderElement = renderElement
		renderElement.InsertChildRenderObject(o.renderObject)
	}
}

// DetachRenderObject 遍历解绑Render对象
func (o *RenderElement) DetachRenderObject() {
	if o.renderElement != nil {
		o.renderElement.RemoveChildRenderObject(o.renderObject)
	}
}

// Update 更新widget
func (o *RenderElement) Update(widget IWidget) {
	o.Element.Update(widget)
}

// InsertChildRenderObject 插入子Render对象
func (o *RenderElement) InsertChildRenderObject(renderObject IRenderObject) {
	o.renderObject.InsertChildRenderObject(renderObject)
}

// RemoveChildRenderObject 移除子Render对象
func (o *RenderElement) RemoveChildRenderObject(renderObject IRenderObject) {
	o.renderObject.RemoveChildRenderObject(renderObject)
}