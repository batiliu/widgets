package framework

// IWidget Widget接口
type IWidget interface {	
	CreateElement() IElement
	CreateRenderObject(IRenderElement) IRenderObject
}

// Widget Widget类型
type Widget struct {
}
