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

type tmplv map[string]interface{}

func makePageVars(cfg Config, tree *Tree) (map[string]tmplv, error) {

	var pages = make(map[string]tmplv, len(tree.pages))
	for path, page := range tree.pages {

		m := make(Metadata, len(page.Metadata))
		for k, v := range page.Metadata {
			m[k] = v
		}
		err := mergo.Merge(&m, cfg.Defaults)
		if err != nil {
			return nil, err
		}

		m["contents"] = template.HTML(page.Rendered)
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

	vars := tmplv{
		"site": tmplv{
			"pages": pages,
		},
		"page": page,
	}

	var templateName string
	switch t := page["template"].(type) {
	case string:
		templateName = t
	default:
		return nil, errors.New("no template defined")
	}

	var buf bytes.Buffer
	err := tree.templates.ExecuteTemplate(&buf, templateName, vars)

	return buf.Bytes(), err
}
