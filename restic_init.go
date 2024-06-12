package resticgo

import (
	"encoding/json"
	"errors"
	"regexp"
)

type InitOptFunc func(*InitOpts)

type InitOpts struct {
	copyChunkerParams bool
	fromRepo          string
	fromPasswordFile  string
}

func InitWithCopyChunkerParams(o *InitOpts) {
	o.copyChunkerParams = true
}

func InitWithFromRepo(repo string) InitOptFunc {
	return func(o *InitOpts) {
		o.fromRepo = repo
	}
}

func InitWithFromPasswordFile(file string) InitOptFunc {
	return func(o *InitOpts) {
		o.fromPasswordFile = file
	}
}

func defaultInitOpts() InitOpts {
	return InitOpts{copyChunkerParams: false}
}

func (r *Restic) Init(opts ...InitOptFunc) (string, error) {
	o := defaultInitOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "init")

	if o.copyChunkerParams {
		cmd = append(cmd, "--copy-chunker-params")
	}

	if len(o.fromRepo) > 0 {
		cmd = append(cmd, "--from-repo", o.fromRepo)
	}

	if len(o.fromPasswordFile) > 0 {
		cmd = append(cmd, "--from-password-file", o.fromPasswordFile)
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		return "", err
	}

	return _parseInitResult(resultRaw)
}

type parseInitResult struct {
	ID string `json:"id"`
}

func _parseInitResult(result string) (string, error) {
	var jsonData parseInitResult
	if err := json.Unmarshal([]byte(result), &jsonData); err != nil {
		return _parseInitPlaintextResult(result)
	}
	return jsonData.ID, nil
}

func _parseInitPlaintextResult(result string) (string, error) {
	// Parse legacy plaintext result.
	//
	//    Prior to restic 0.15.0, restic returned a plaintext response even when the
	//    caller specified --json.
	pattern := regexp.MustCompile(`created restic repository ([a-z0-9]+) at .+`)
	matches := pattern.FindStringSubmatch(result)
	if len(matches) < 2 {
		return "", errors.New("unexpected plaintext result format")
	}
	return matches[1], nil
}
