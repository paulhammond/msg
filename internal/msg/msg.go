package msg

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var version string = "dev"

func Version() string {
	return version
}

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
	pages, err := makePageVars(cfg, tree)
	if err != nil {
		return err
	}
	for path, page := range tree.pages {

		fsPath := cfg.output + "/" + page.OutputPath
		fmt.Printf("rendering %s\n", fsPath)
		rendered, err := render(tree, pages, path)
		if err != nil {
			return err
		}
		err = os.MkdirAll(filepath.Dir(fsPath), 0755)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fsPath, rendered, 0644)
		if err != nil {
			return err
		}
	}
	for path, asset := range tree.assets {
		dst := cfg.output + "/" + asset.outputPath
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
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		return err
	}

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
