package gonew

type Template interface {
	Name() string
	Description() string
	Make(args []string) error
}

type TemplateContainer map[string]Template

func (this TemplateContainer) RegistTemplate(name string, tmp Template) *TemplateContainer {
	this[name] = tmp
	return &this
}

func (this TemplateContainer) GetTemplate(name string) Template {
	t, ok := this[name]
	if ok {
		return t
	}
	return nil
}

var c = &TemplateContainer{}

func GetTemplateContainer() *TemplateContainer {
	return c
}
