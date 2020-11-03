package msg

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// An asset represents an asset file
type asset struct {
	fspath     string
	sourcePath string
	outputPath string
	path       string
	hash       []byte
}

func newAsset(path, fspath string, cfg Config) *asset {
	return &asset{
		fspath:     fspath,
		sourcePath: path,
		outputPath: rewritePath(cfg.FileRewrites, path),
		path:       "/" + rewritePath(cfg.URLRewrites, path),
	}
}

type tmplAsset map[string]interface{}

func (a tmplAsset) String() string {
	if s, ok := a["path"].(string); ok {
		return s
	}
	return ""
}

func (a *asset) tmplv() tmplAsset {
	return tmplAsset{
		"path":      a.path,
		"integrity": a.integrity(),
	}
}

func (a *asset) integrity() string {
	if a.hash == nil {
		return ""
	}
	return "sha384-" + base64.StdEncoding.EncodeToString(a.hash)
}

func (a *asset) hashify() error {
	if a.hash != nil {
		return nil
	}

	hash, err := a.sha()
	if err != nil {
		return err
	}

	a.hash = hash
	hex := fmt.Sprintf("%.5x", hash)
	a.outputPath = insertHash(a.outputPath, hex)
	a.path = insertHash(a.path, hex)
	return nil
}

func insertHash(filepath string, hash string) string {
	ext := path.Ext(filepath)
	return strings.TrimSuffix(filepath, ext) + "-" + hash + ext
}

func (a *asset) sha() (b []byte, err error) {
	fd, err := os.Open(a.fspath)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := fd.Close()
		if err == nil {
			err = e
		}
	}()

	h := sha512.New384()
	if _, err := io.Copy(h, fd); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
