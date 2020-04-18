package framework

// IElement Element接口
type IElement interface {
	Mount()
	Unmount()

	AttachRenderObject()
	DetachRenderObject()

	Update(IWidget)

	Parent() IElement
}

// Element Element类型
type Element struct {
	widget IWidget
	parent IElement
}

// Parent 返回父对象
func (o *Element) Parent() IElement {
	return o.parent
}

// Mount 挂载Element对象
func (o *Element) Mount() {

}

// Unmount 卸载Element对象
func (*Element) Unmount() {

}

// AttachRenderObject 挂载Element对象
func (o *Element) AttachRenderObject() {

}

// DetachRenderObject 卸载Element对象
func (*Element) DetachRenderObject() {

}

// UpdateChild 更新子element
func (o *Element) UpdateChild(element IElement, widget IWidget) IElement {
	if widget == nil {
		if element != nil {
			element.DetachRenderObject();
        }
        return nil;
	}
			
	if element != nil {
		element.Update(widget)
		return element
	}

	return o.InflateWidget(widget)
}

// InflateWidget 展开widget
func (*Element) InflateWidget(widget IWidget) IElement {
	element := widget.CreateElement();
	element.Mount()
	return element
}

// Update 更新widget
func (o *Element) Update(widget IWidget) {
	o.widget = widget
}
