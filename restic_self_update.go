package resticgo

func (r *Restic) SelfUpdate() (string, error) {
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "self-update")

	return execute(cmd)
}
