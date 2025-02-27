package resticgo

import (
	"encoding/json"
	"strings"
)

type FindOptFunc func(*FindOpts)

type FindOpts struct {
	snapshotId string
}

func FindWithSnapshotId(snapshotId string) FindOptFunc {
	return func(o *FindOpts) {
		o.snapshotId = snapshotId
	}
}

func (r *Restic) Find(pattern string, opts ...FindOptFunc) ([]map[string]interface{}, error) {
	o := FindOpts{}
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "find", pattern)

	if len(o.snapshotId) > 0 {
		cmd = append(cmd, "-s", o.snapshotId)
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		return nil, err
	}

	var out []map[string]interface{}
	if err = json.Unmarshal([]byte(strings.TrimSuffix(resultRaw, "\n")), &out); err != nil {
		return nil, err
	}

	return out, nil
}
