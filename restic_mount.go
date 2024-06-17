package resticgo

type MountOptFunc func(*MountOpts)

type MountOpts struct {
	allowOther           bool
	host                 string
	noDefaultPermissions bool
	ownerRoot            bool
	paths                []string
	pathTemplates        []string
	tags                 []string
	timeTemplate         string
}

func MountWithAllowOther(o *MountOpts) {
	o.allowOther = true
}

func MountWithHost(host string) MountOptFunc {
	return func(o *MountOpts) {
		o.host = host
	}
}

func MountWithNoDefaultPermissions(o *MountOpts) {
	o.noDefaultPermissions = true
}

func MountWithOwnerRoot(o *MountOpts) {
	o.ownerRoot = true
}

func MountWithPaths(paths []string) MountOptFunc {
	return func(o *MountOpts) {
		o.paths = append(o.paths, paths...)
	}
}

func MountWithPathTemplates(paths []string) MountOptFunc {
	return func(o *MountOpts) {
		o.pathTemplates = append(o.pathTemplates, paths...)
	}
}

func MountWithTags(tags []string) MountOptFunc {
	return func(o *MountOpts) {
		o.tags = append(o.tags, tags...)
	}
}

func MountWithTimeTemplate(timeTemplate string) MountOptFunc {
	return func(o *MountOpts) {
		o.timeTemplate = timeTemplate
	}
}

func defaultMountOpts() MountOpts {
	return MountOpts{allowOther: false, ownerRoot: false}
}

func (r *Restic) Mount(mountPoint string, opts ...MountOptFunc) (string, error) {
	o := defaultMountOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "mount", mountPoint)

	if o.allowOther {
		cmd = append(cmd, "--allow-other")
	}

	if o.host != "" {
		cmd = append(cmd, "--host", o.host)
	}

	if o.noDefaultPermissions {
		cmd = append(cmd, "--no-default-permissions")
	}

	if o.ownerRoot {
		cmd = append(cmd, "--owner-root")
	}

	for _, path := range o.paths {
		cmd = append(cmd, "--paths", path)
	}

	for _, pathTemplate := range o.pathTemplates {
		cmd = append(cmd, "--paths-template", pathTemplate)
	}

	for _, tag := range o.tags {
		cmd = append(cmd, "--tag", tag)
	}

	if o.timeTemplate != "" {
		cmd = append(cmd, "--time-template", o.timeTemplate)
	}

	return execute(cmd)
}
