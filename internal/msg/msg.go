package msg

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var version string = "dev"

// Run runs msg according to the config file at configPath and saves the output
// in outputDir
func Run(configPath, outputDir string) error {
	cfg, err := parseConfig(configPath, outputDir)
	if err != nil {
		return err
	}

	if cfg.Version != version {
		return fmt.Errorf("site requires msg version %q, have %q", cfg.Version, version)
	}

	tree, err := newTree(cfg)
	if err != nil {
		return err
	}

	return renderAll(cfg, tree)

}

// renderAll renders the site
func renderAll(cfg Config, tree *Tree) error {
	for path := range tree.pages {

		fsPath := cfg.output + "/" + rewritePath(cfg, path)
		fmt.Printf("rendering %s\n", fsPath)
		rendered, err := render(cfg, tree, path)
		if err != nil {
			return err
		}
		fmt.Println(fsPath, filepath.Dir(fsPath))
		err = os.MkdirAll(filepath.Dir(fsPath), 0755)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fsPath, rendered, 0644)
		if err != nil {
			return err
		}
	}
	for _, path := range tree.assets {
		dst := cfg.output + "/" + rewritePath(cfg, path)
		src := cfg.root() + "/" + path
		fmt.Printf("copying %s to %s\n", src, dst)
		err := copyFile(src, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// copyFile copies the file at src to dst, making any needed directories
func copyFile(src string, dst string) (err error) {
	w, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		e := w.Close()
		if err == nil {
			err = e
		}
	}()

	err = os.MkdirAll(filepath.Dir(src), 0755)
	if err != nil {
		return err
	}

	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		e := r.Close()
		if err == nil {
			err = e
		}
	}()

	_, err = io.Copy(w, r)
	return err
}
