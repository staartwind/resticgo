package resticgo

func (r *Restic) Cat(catType, id string) (string, error) {
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "cat", catType, id)

	return execute(cmd)
}
