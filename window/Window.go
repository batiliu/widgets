package window

import (
	"tencent.com/widgets/widget"
)

// Window 应用程序窗口
type Window struct {
	RootWidget       widget.Widget
	RootElement      widget.Element
	RootRenderObject widget.RenderObject
}
