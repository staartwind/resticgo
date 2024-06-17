package resticgo

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"strings"
)

type LsOptFunc func(*LsOpts)

type LsOpts struct {
	host          string
	humanReadable bool
	long          bool
	paths         []string
	recursive     bool
	tags          []string
}

func LsWithHost(host string) LsOptFunc {
	return func(o *LsOpts) {
		o.host = host
	}
}

func LsWithHumanReadable(o *LsOpts) {
	o.humanReadable = true
}

func LsWithLong(o *LsOpts) {
	o.long = true
}

func LsWithPaths(paths []string) LsOptFunc {
	return func(o *LsOpts) {
		o.paths = append(o.paths, paths...)
	}
}

func LsWithRecursive(o *LsOpts) {
	o.recursive = true
}

func LsWithTags(tags []string) LsOptFunc {
	return func(o *LsOpts) {
		o.tags = append(o.tags, tags...)
	}
}

func defaultLsOpts() LsOpts {
	return LsOpts{humanReadable: false, long: false, recursive: false}
}

func (r *Restic) Ls(snapshotId string, opts ...LsOptFunc) ([]map[string]interface{}, error) {
	o := defaultLsOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "ls")
	cmd = append(cmd, snapshotId)

	if o.host != "" {
		cmd = append(cmd, "--host", o.host)
	}

	if o.humanReadable {
		cmd = append(cmd, "--human-readable")
	}

	if o.long {
		cmd = append(cmd, "--long")
	}

	for _, path := range o.paths {
		cmd = append(cmd, "--paths", path)
	}

	if o.recursive {
		cmd = append(cmd, "--recursive")
	}

	for _, tag := range o.tags {
		cmd = append(cmd, "--tag", tag)
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		return nil, err
	}

	var out []map[string]interface{}

	scanner := bufio.NewScanner(strings.NewReader(resultRaw))
	for scanner.Scan() {
		var m map[string]interface{}
		if err = json.Unmarshal([]byte(scanner.Text()), &m); err != nil {
			slog.Info("unable to unmarshal json result", "err", err)
			continue
		}
		out = append(out, m)
	}

	return out, nil
}
