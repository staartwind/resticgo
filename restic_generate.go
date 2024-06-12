package resticgo

type GenerateOptFunc func(*GenerateOpts)

type GenerateOpts struct {
	bashCompletionPath string
	manDirectory       string
	zshCompletionPath  string
}

func GenerateWithBashCompletionPath(path string) GenerateOptFunc {
	return func(o *GenerateOpts) {
		o.bashCompletionPath = path
	}
}

func GenerateWithManDirectory(dir string) GenerateOptFunc {
	return func(o *GenerateOpts) {
		o.manDirectory = dir
	}
}

func GenerateWithZshCompletionPath(path string) GenerateOptFunc {
	return func(o *GenerateOpts) {
		o.zshCompletionPath = path
	}
}

func (r *Restic) Generate(opts ...GenerateOptFunc) (string, error) {
	o := GenerateOpts{}
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "generate")

	if len(o.bashCompletionPath) > 0 {
		cmd = append(cmd, "--bash-completion", o.bashCompletionPath)
	}

	if len(o.manDirectory) > 0 {
		cmd = append(cmd, "--man", o.manDirectory)
	}

	if len(o.zshCompletionPath) > 0 {
		cmd = append(cmd, "--zsh-completion", o.zshCompletionPath)
	}

	return execute(cmd)
}
