package widget

// ComponentElement 组合元素对象
type ComponentElement interface {
	Child() Element
	Build() Widget
}
