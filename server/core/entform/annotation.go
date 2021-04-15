package entform

// Annotation 表单schema的描述
type Annotation struct {
}

// Name 实现 ent.Annotation 接口
func (Annotation) Name() string {
	return "EntForm"
}
