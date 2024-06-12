package resticgo

type CopyOptFunc func(*CopyOpts)

type CopyOpts struct {
	fromRepo         string
	fromPasswordFile string
}

func CopyWithFromRepo(repo string) CopyOptFunc {
	return func(o *CopyOpts) {
		o.fromRepo = repo
	}
}

func CopyWithFromPasswordFile(file string) CopyOptFunc {
	return func(o *CopyOpts) {
		o.fromPasswordFile = file
	}
}

func (r *Restic) Copy(opts ...CopyOptFunc) (string, error) {
	o := CopyOpts{}
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "copy")

	if len(o.fromRepo) > 0 {
		cmd = append(cmd, "--from-repo", o.fromRepo)
	}

	if len(o.fromPasswordFile) > 0 {
		cmd = append(cmd, "--from-password-file", o.fromPasswordFile)
	}

	return execute(cmd)
}
