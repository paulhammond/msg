package msg

import (
	"bytes"
	"io/ioutil"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

// Page represents a single page
type Page struct {
	Metadata   Metadata
	Rendered   string
	Source     string
	SourcePath string
	OutputPath string
	Path       string
}

// parsePage parses the file at path into a page struct
func parsePage(path, fspath string, cfg Config) (*Page, error) {
	source, err := ioutil.ReadFile(fspath)
	if err != nil {
		return nil, err
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	context := parser.NewContext()
	var buf bytes.Buffer
	if err := md.Convert(source, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}
	metadata := meta.Get(context)

	outputPath := rewritePath(cfg.FileRewrites, path)
	return &Page{
		Metadata:   metadata,
		Rendered:   buf.String(),
		Source:     string(source),
		SourcePath: path,
		OutputPath: outputPath,
		Path:       rewritePath(cfg.URLRewrites, outputPath),
	}, nil

}

// Metadata represents the YAML "Front Matter" of a page
type Metadata map[string]interface{}

// string is a convienience for fetching a string from metadata
func (m Metadata) string(key string) string {
	v := m[key]
	switch t := v.(type) {
	case string:
		return t
	default:
		return ""
	}
}
