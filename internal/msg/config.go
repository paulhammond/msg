package msg

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config represents a parsed config file
type Config struct {
	Root      string     `yaml:"root"`
	Templates []string   `yaml:"templates"`
	Pages     []string   `yaml:"pages"`
	Ignore    []string   `yaml:"ignore"`
	Rewrites  []*Rewrite `yaml:"rewrites"`
	Defaults  Metadata   `yaml:"defaults"`
	output    string
	cfgPath   string
}

// ParseConfig parses the config file at cfgPath into a Config struct
func parseConfig(cfgPath string, output string) (Config, error) {
	cfg := Config{
		Root:      ".",
		Templates: []string{"**/*.tmpl"},
		Pages:     []string{"**/*.md"},
		Ignore:    []string{},
		output:    output,
		cfgPath:   cfgPath,
	}

	source, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return cfg, err
	}

	err = yaml.UnmarshalStrict(source, &cfg)
	if err != nil {
		return cfg, err
	}

	// if config file is inside the directory ignore it
	if cfg.Root == "." {
		cfg.Ignore = append(cfg.Ignore, filepath.Base(cfgPath))
	}

	// if output folder is inside source folder ignore it
	rel, err := relativeOutput(cfg.root(), output)
	if err != nil {
		return cfg, err
	}
	if rel != "" {
		cfg.Ignore = append(cfg.Ignore, rel+"/**")
	}

	// lastly compile rewrites
	for _, r := range cfg.Rewrites {
		err = r.parse()
		if err != nil {
			return cfg, err
		}
	}

	return cfg, err
}

// relativeoutput gives the relative path to output from root. If the output
// directory is outside root then it returns an empty string.
func relativeOutput(root, output string) (string, error) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}

	absOutput, err := filepath.Abs(output)
	if err != nil {
		return "", err
	}

	rel, err := filepath.Rel(absRoot, absOutput)
	if err != nil || strings.HasPrefix(rel, "../") {
		return "", nil
	}
	return rel, nil
}

// root returns the root input directory for the site
func (c Config) root() string {
	cfgDir := filepath.Dir(c.cfgPath)
	root := filepath.Clean(filepath.Join(cfgDir, c.Root))
	return strings.TrimSuffix(root, "/")
}
