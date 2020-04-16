package widget

// RenderWidget 界面声明渲染对象
type RenderWidget struct {
}

// CreateElement 创建界面元素对象
func (*RenderWidget) CreateElement() Element {
	return &RenderElement{}
}

// CreateRenderObject 创建界面渲染对象
func (*RenderWidget) CreateRenderObject() RenderObject {
	return &TextRenderObject{}
}
