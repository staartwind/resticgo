package resticgo

func (r *Restic) Unlock() (string, error) {
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "unlock")

	return execute(cmd)
}
