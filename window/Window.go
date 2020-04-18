package window

import (
	"tencent.com/widgets/framework"
)

// Window 应用程序窗口
type Window struct {
	rootWidget       framework.IWidget
	rootElement      framework.IElement
	rootRenderObject framework.IRenderObject
}

// AttachRootWidget 绑定根Widget
func (o *Window) AttachRootWidget(widget framework.IWidget) {
	o.rootWidget = widget
	o.rootElement = widget.CreateElement()
	o.rootElement.Mount()
}
