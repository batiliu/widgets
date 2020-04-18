package framework

// TextRenderObject 文字渲染对象
type TextRenderObject struct {
	RenderObject
}

// MarkNeedsLayout 标记需要重新布局
func (*TextRenderObject) MarkNeedsLayout() {

}

// MarkNeedsCompositing 标记需要重新合成
func (*TextRenderObject) MarkNeedsCompositing() {

}

// MarkNeedsPaint 标记需要重新绘制
func (*TextRenderObject) MarkNeedsPaint() {

}

// Layout 执行布局
func (*TextRenderObject) Layout() {

}

// Paint 执行绘制
func (*TextRenderObject) Paint() {

}

// InsertChildRenderObject 插入子Render对象
func (o *TextRenderObject) InsertChildRenderObject(renderObject IRenderObject) {
	o.RenderObject.InsertChildRenderObject(renderObject)
}

// RemoveChildRenderObject 移除子Render对象
func (o *TextRenderObject) RemoveChildRenderObject(renderObject IRenderObject) {
	o.RenderObject.RemoveChildRenderObject(renderObject)
}
