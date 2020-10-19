package msg

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"

	"github.com/imdario/mergo"
	"github.com/paulhammond/msg/internal/tmpl"
)

type templateSet struct {
	*template.Template
}

func newTemplateSet() templateSet {
	set := templateSet{template.New("")}
	set.Funcs(tmpl.Map())
	return set
}

func (t *templateSet) Parse(path, fsPath string) error {

	source, err := ioutil.ReadFile(fsPath)
	if err != nil {
		return err
	}

	t2 := t.New(path)
	_, err = t2.Parse(string(source))
	if err != nil {
		return err
	}
	return nil
}

// Variables represents the data available to templates
type Variables struct {
	Site *Tree
	Page *Page
}

func render(cfg Config, tree *Tree, path string) ([]byte, error) {

	p, ok := tree.pages[path]

	m := make(Metadata, len(p.Metadata))
	for k, v := range p.Metadata {
		m[k] = v
	}
	err := mergo.Merge(&m, cfg.Defaults)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("not found")
	}

	vars := Variables{
		Site: tree,
		Page: &Page{
			Metadata: m,
			Rendered: p.Rendered,
			Source:   p.Source,
		},
	}

	var buf bytes.Buffer
	err = tree.templates.ExecuteTemplate(&buf, m.string("template"), vars)

	return buf.Bytes(), err
}
