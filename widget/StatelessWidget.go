package widget

// StatelessWidget 无状态声明对象
type StatelessWidget struct {
	s string
}

// CreateElement 创建界面元素对象
func (*StatelessWidget) CreateElement() Element {
	return &StatelessElement{}
}

// CreateRenderObject 创建界面渲染对象
func (*StatelessWidget) CreateRenderObject() RenderObject {
	return nil
}
