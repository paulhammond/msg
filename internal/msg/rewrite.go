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

// parse parses the regexp in From
func (rr *Rewrite) parse() (err error) {
	rr.r, err = regexp.Compile(rr.From)
	return
}

// rewrite performs the rewrite on path
func (rr *Rewrite) rewrite(path string) string {
	return rr.r.ReplaceAllString(path, rr.To)
}

// rewritePath rewrites path using all rules in cfg
func rewritePath(cfg Config, path string) string {
	for _, r := range cfg.Rewrites {
		path = r.rewrite(path)
	}
	return path
}
