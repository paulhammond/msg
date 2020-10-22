package msg

import (
	"bytes"
	"io/ioutil"
	"strings"

	"github.com/imdario/mergo"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
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
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	context := parser.NewContext()
	var buf bytes.Buffer
	if err := md.Convert(source, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}
	goldmarkMeta := meta.Get(context)

	metadata := make(Metadata, len(goldmarkMeta))
	for k, v := range goldmarkMeta {
		metadata[k] = v
	}
	err = mergo.Merge(&metadata, cfg.Defaults)
	if err != nil {
		return nil, err
	}

	outputPath := metadata.string("file", rewritePath(cfg.FileRewrites, path))
	urlPath := metadata.string("path", rewritePath(cfg.URLRewrites, outputPath))

	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}

	return &Page{
		Metadata:   metadata,
		Rendered:   buf.String(),
		Source:     string(source),
		SourcePath: path,
		OutputPath: outputPath,
		Path:       urlPath,
	}, nil

}

// Metadata represents the YAML "Front Matter" of a page
type Metadata map[string]interface{}

// string is a convienience for fetching a string from metadata
func (m Metadata) string(key, fallback string) string {
	v := m[key]
	switch t := v.(type) {
	case string:
		return t
	default:
		return fallback
	}
}
