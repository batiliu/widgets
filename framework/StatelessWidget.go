package framework

// IStatelessWidget 无状态Widget接口
type IStatelessWidget interface {
	IWidget
	Build() IWidget
}

// StatelessWidget 无状态Widget类型
type StatelessWidget struct {
	Widget
}

// CreateElement 创建Element对象
func (*StatelessWidget) CreateElement() IElement {
	return &StatelessElement{}
}

// CreateRenderObject 创建界面渲染对象
func (*StatelessWidget) CreateRenderObject(element IRenderElement) IRenderObject {
	return nil
}

// Build 创建Widget对象
func (*StatelessWidget) Build() IWidget {
	return nil
}
