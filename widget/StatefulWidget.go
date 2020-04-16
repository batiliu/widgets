package widget

// StatefulWidget 有状态声明对象
type StatefulWidget struct {
	i int
}

// CreateElement 创建界面元素对象
func (*StatefulWidget) CreateElement() Element {
	return &StatefulElement{}
}

// CreateRenderObject 创建界面渲染对象
func (*StatefulWidget) CreateRenderObject() RenderObject {
	return nil
}
