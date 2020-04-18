package framework

// IRenderWidget 渲染Widget接口
type IRenderWidget interface {
	IWidget
}

// RenderWidget 渲染Widget类型
type RenderWidget struct {
	Widget
}

// CreateElement 创建界面元素对象
func (o *RenderWidget) CreateElement() IElement {
	element := &RenderElement{}
	element.widget = o
	return element
}

// CreateRenderObject 创建界面渲染对象
func (*RenderWidget) CreateRenderObject(element IRenderElement) IRenderObject {
	return &TextRenderObject{}
}
