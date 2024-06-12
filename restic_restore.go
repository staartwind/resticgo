package resticgo

type RestoreOptFunc func(*RestoreOpts)

type RestoreOpts struct {
	snapshotId string
	include    string
	exclude    string
	targetDir  string
}

func RestoreWithSnapshotId(id string) RestoreOptFunc {
	return func(o *RestoreOpts) {
		o.snapshotId = id
	}
}

func RestoreWIthInclude(include string) RestoreOptFunc {
	return func(o *RestoreOpts) {
		o.include = include
	}
}

func RestoreWithExclude(exclude string) RestoreOptFunc {
	return func(o *RestoreOpts) {
		o.exclude = exclude
	}
}

func RestoreWithTargetDir(dir string) RestoreOptFunc {
	return func(o *RestoreOpts) {
		o.targetDir = dir
	}
}

func defaultRestoreOpts() RestoreOpts {
	return RestoreOpts{snapshotId: "latest"}
}

func (r *Restic) Restore(opts ...RestoreOptFunc) (string, error) {
	o := defaultRestoreOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "restore", o.snapshotId)

	if len(o.include) > 0 {
		cmd = append(cmd, "--include", o.include)
	}

	if len(o.exclude) > 0 {
		cmd = append(cmd, "--exclude", o.exclude)
	}

	if len(o.targetDir) > 0 {
		cmd = append(cmd, "--target-dir", o.targetDir)
	}

	return execute(cmd)
}
