package gonew

type Template interface {
	Name() string
	Description() string
	Usage() string
	Make(args []string) error
}

type TemplateContainer map[string]Template

func (this TemplateContainer) RegistTemplate(name string, tmp Template) *TemplateContainer {
	this[name] = tmp
	return &this
}

func (this TemplateContainer) GetTemplate(name string) Template {
	return this[name]
}

func (this TemplateContainer) GetTemplates() []Template {
	tmps := make([]Template, len(this))
	var i = 0
	for _, v := range this {
		tmps[i] = v
		i = i + 1
	}
	return tmps
}

var c = &TemplateContainer{}

func GetTemplateContainer() *TemplateContainer {
	return c
}
