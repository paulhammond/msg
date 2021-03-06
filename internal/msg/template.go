package msg

import (
	"bytes"
	"errors"
	"html/template"
	"io/ioutil"

	"github.com/paulhammond/msg/internal/tmpl"
)

type templateSet struct {
	*template.Template
}

func newTemplateSet() *templateSet {
	set := templateSet{template.New("")}
	set.Funcs(renderContext{}.funcs())
	set.Funcs(tmpl.Map())
	return &set
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

type tmplv map[string]interface{}

func makePageVars(cfg Config, tree *Tree) (map[string]tmplv, error) {

	var pages = make(map[string]tmplv, len(tree.pages))
	for path, page := range tree.pages {

		m := make(tmplv, len(page.Metadata))
		for k, v := range page.Metadata {
			m[k] = v
		}
		m["content"] = template.HTML(page.Rendered)
		m["source"] = page.Source
		m["path"] = page.Path
		pages[path] = tmplv(m)
	}

	return pages, nil
}

func render(tree *Tree, pages map[string]tmplv, path string) ([]byte, error) {

	page, ok := pages[path]
	if !ok {
		return nil, errors.New("not found")
	}

	site := tmplv(tree.metadata)
	site["pages"] = pages

	vars := tmplv{
		"site": site,
		"page": page,
	}

	ctx := renderContext{
		Tree: tree,
		Site: site,
		Page: page,
	}

	var templateName string
	switch t := page["template"].(type) {
	case string:
		templateName = t
	default:
		return nil, errors.New("no template defined")
	}

	cloned, err := tree.templates.Clone()
	if err != nil {
		return nil, err
	}
	cloned.Funcs(ctx.funcs())
	cloned.Funcs(tmpl.Map())

	var buf bytes.Buffer
	err = cloned.ExecuteTemplate(&buf, templateName, vars)

	return buf.Bytes(), err
}
