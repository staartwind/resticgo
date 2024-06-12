package resticgo

import (
	"encoding/json"
	"strings"
)

func (r *Restic) Find(pattern string) ([]map[string]interface{}, error) {
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "find", pattern)

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
