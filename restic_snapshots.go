package resticgo

import (
	"encoding/json"
	"strings"
)

type SnapshotOptFunc func(*SnapshotOpts)

type SnapshotOpts struct {
	snapshotId string
	groupBy    string
	tags       []string
}

func SnapshotWithSnapshotId(snapshotId string) SnapshotOptFunc {
	return func(o *SnapshotOpts) {
		o.snapshotId = snapshotId
	}
}

func SnapshotWithGroupBy(groupBy string) SnapshotOptFunc {
	return func(o *SnapshotOpts) {
		o.groupBy = groupBy
	}
}

func SnapshotWithTags(tags ...string) SnapshotOptFunc {
	return func(o *SnapshotOpts) {
		o.tags = tags
	}
}

func (r *Restic) Snapshots(opts ...SnapshotOptFunc) ([]map[string]interface{}, error) {
	o := SnapshotOpts{}
	for _, opt := range opts {
		opt(&o)
	}
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "snapshots")

	if len(o.snapshotId) > 0 {
		cmd = append(cmd, o.snapshotId)
	}

	if len(o.groupBy) > 0 {
		cmd = append(cmd, "--group-by", o.groupBy)
	}

	if len(o.tags) > 0 {
		cmd = append(cmd, "--tag", strings.Join(o.tags, ","))
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
