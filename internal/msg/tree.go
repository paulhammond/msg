package msg

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v2"
)

// A Tree represents the parsed contents of the input directory
type Tree struct {
	pages     map[string]*Page
	templates templateSet
	assets    map[string]string
}

// newTree parses all files referenced in cfg and creates a Tree
func newTree(cfg Config) (*Tree, error) {
	root := cfg.root()
	tree := Tree{
		pages:     map[string]*Page{},
		templates: newTemplateSet(),
		assets:    map[string]string{},
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		trimmed := strings.TrimPrefix(path, root+"/")

		isIgnored, err := matches(cfg.Ignore, trimmed)
		if err != nil {
			return err
		}
		if isIgnored {
			return nil
		}

		isPage, err := matches(cfg.Pages, trimmed)
		if err != nil {
			return err
		}
		if isPage {
			tree.pages[trimmed], err = parsePage(trimmed, path, cfg)
			if err != nil {
				return err
			}
		}

		isTemplate, err := matches(cfg.Templates, trimmed)
		if err != nil {
			return err
		}
		if isTemplate {
			err := tree.templates.Parse(trimmed, path)
			if err != nil {
				return err
			}
		}

		if !isTemplate && !isPage {
			tree.assets[trimmed] = rewritePath(cfg.FileRewrites, trimmed)
		}

		return nil
	})

	return &tree, err
}

// matches returns true if path matches any of the doublestar patterns
func matches(patterns []string, path string) (matched bool, err error) {
	for _, pattern := range patterns {
		matched, err = doublestar.Match(pattern, path)
		if matched || err != nil {
			return
		}
	}
	return false, nil
}
