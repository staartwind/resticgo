package resticgo

import "strings"

func (r *Restic) ListLocks() ([]string, error) {
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "list", "locks")

	resultRaw, err := execute(cmd)
	if err != nil {
		return nil, err
	}

	return strings.Split(resultRaw, "\n"), nil
}
