package framework

// IStatefulWidget 有状态Widget接口
type IStatefulWidget interface {
	IWidget
	CreateState() IState
}

// StatefulWidget 有状态Widget类型
type StatefulWidget struct {
	Widget
}

// CreateElement 创建Element对象
func (o *StatefulWidget) CreateElement() IElement {
	element := &StatefulElement{}
	element.state = o.CreateState()
	return element
}

// CreateRenderObject 创建界面渲染对象
func (*StatefulWidget) CreateRenderObject(element IRenderElement) IRenderObject {
	return nil
}

// CreateState 创建状态对象
func (*StatefulWidget) CreateState() IState {
	return &State{}
}
