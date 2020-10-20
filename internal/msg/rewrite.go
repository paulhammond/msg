package msg

import (
	"regexp"
)

// Rewrite represents a rewrite config
type Rewrite struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
	r    *regexp.Regexp
}

// compiles compiles the regexp in From
func (rr *Rewrite) compile() (err error) {
	rr.r, err = regexp.Compile(rr.From)
	return
}

// rewrite performs the rewrite on path
func (rr *Rewrite) rewrite(path string) string {
	return rr.r.ReplaceAllString(path, rr.To)
}

// rewritePath rewrites path using all rules in cfg
func rewritePath(list []*Rewrite, path string) string {
	for _, r := range list {
		path = r.rewrite(path)
	}
	return path
}

func compileRewrites(list []*Rewrite) error {
	var err error
	for _, r := range list {
		err = r.compile()
		if err != nil {
			return err
		}
	}
	return nil
}
