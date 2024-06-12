package resticgo

import (
	"errors"
	"regexp"
)

type Version struct {
	ResticVersion   string
	GoVersion       string
	PlatformVersion string
	Architecture    string
}

func (r *Restic) Version() (Version, error) {
	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "version")

	resultRaw, err := execute(cmd)
	if err != nil {
		return Version{}, err
	}

	pattern := regexp.MustCompile(`restic ([0-9\.]+) compiled with go([0-9\.]+) on ([a-zA-Z0-9]+)/([a-zA-Z0-9]+)`)
	matches := pattern.FindStringSubmatch(resultRaw)
	if len(matches) < 5 {
		return Version{}, errors.New("unexpected output format")
	}
	return Version{
		ResticVersion:   matches[1],
		GoVersion:       matches[2],
		PlatformVersion: matches[3],
		Architecture:    matches[4],
	}, nil
}
