package framework

// IRenderObject Render接口
type IRenderObject interface {
	MarkNeedsLayout()
	MarkNeedsCompositing()
	MarkNeedsPaint()
	Layout()
	Paint()

	InsertChildRenderObject(IRenderObject)
	RemoveChildRenderObject(IRenderObject)
}

// RenderObject Render类型
type RenderObject struct {
	child IRenderObject
}

// MarkNeedsLayout 标记需要重新布局
func (*RenderObject) MarkNeedsLayout() {

}

// MarkNeedsCompositing 标记需要重新合成
func (*RenderObject) MarkNeedsCompositing() {

}

// MarkNeedsPaint 标记需要重新绘制
func (*RenderObject) MarkNeedsPaint() {

}

// Layout 执行布局
func (*RenderObject) Layout() {

}

// Paint 执行绘制
func (*RenderObject) Paint() {

}

// InsertChildRenderObject 插入子Render对象
func (o *RenderObject) InsertChildRenderObject(renderObject IRenderObject) {
	o.child = renderObject
}

// RemoveChildRenderObject 移除子Render对象
func (o *RenderObject) RemoveChildRenderObject(renderObject IRenderObject) {
	o.child = nil
}
