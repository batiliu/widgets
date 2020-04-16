package widget

// Widget 界面声明对象
type Widget interface {
	CreateElement() Element
	CreateRenderObject() RenderObject
}
