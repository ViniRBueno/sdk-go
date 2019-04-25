package template

// TemplateType is the structure of the TemplateType
type TemplateType struct {
	ID   int64
	Name string
}

type TemplateTypesList struct {
	Collection []TemplateType
}
