package resticgo

type RewriteOptFunc func(*RewriteOpts)

type RewriteOpts struct {
	exclude     []string
	excludeFile string
	forget      bool
	dryRun      bool
	snapshotId  string
}

func RewriteWithExclude(exclude ...string) RewriteOptFunc {
	return func(o *RewriteOpts) {
		o.exclude = append(o.exclude, exclude...)
	}
}

func RewriteWithExcludeFile(excludeFile string) RewriteOptFunc {
	return func(o *RewriteOpts) {
		o.excludeFile = excludeFile
	}
}

func RewriteWithForget(o *RewriteOpts) {
	o.forget = true
}

func RewriteWithDryRun(o *RewriteOpts) {
	o.dryRun = true
}

func RewriteWithSnapshotId(id string) RewriteOptFunc {
	return func(o *RewriteOpts) {
		o.snapshotId = id
	}
}

func defaultRewriteOpts() RewriteOpts {
	return RewriteOpts{forget: false, dryRun: false}
}

func (r *Restic) Rewrite(opts ...RewriteOptFunc) (string, error) {
	o := defaultRewriteOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "rewrite")

	for _, exclude := range o.exclude {
		cmd = append(cmd, "--exclude", exclude)
	}

	if len(o.excludeFile) > 0 {
		cmd = append(cmd, "--exclude-file", o.excludeFile)
	}

	if o.forget {
		cmd = append(cmd, "--forget")
	}

	if o.dryRun {
		cmd = append(cmd, "--dry-run")
	}

	if len(o.snapshotId) > 0 {
		cmd = append(cmd, o.snapshotId)
	}

	return execute(cmd)
}
