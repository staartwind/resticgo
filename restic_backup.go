package resticgo

import (
	"encoding/json"
	"fmt"
	"strings"
)

type BackupOptFunc func(*BackupOpts)

type BackupOpts struct {
	excludePatterns []string
	excludeFiles    []string
	tags            []string
	dryRun          bool
	host            string
	scan            bool
}

func BackupWithExcludePatterns(patterns ...string) BackupOptFunc {
	return func(o *BackupOpts) {
		o.excludePatterns = append(o.excludePatterns, patterns...)
	}
}

func BackupWithExcludeFiles(files ...string) BackupOptFunc {
	return func(o *BackupOpts) {
		o.excludeFiles = append(o.excludeFiles, files...)
	}
}

func BackupWithTags(tags ...string) BackupOptFunc {
	return func(o *BackupOpts) {
		o.tags = append(o.tags, tags...)
	}
}

func BackupWithDryRun(o *BackupOpts) {
	o.dryRun = true
}

func BackupWithHost(host string) BackupOptFunc {
	return func(o *BackupOpts) {
		o.host = host
	}
}

func BackupWithoutScan(o *BackupOpts) {
	o.scan = false
}

func defaultBackupOpts() BackupOpts {
	return BackupOpts{
		dryRun: false,
		scan:   true,
	}
}

func (r *Restic) Backup(paths []string, opts ...BackupOptFunc) (map[string]interface{}, error) {
	o := defaultBackupOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "backup")
	cmd = append(cmd, paths...)

	for _, pattern := range o.excludePatterns {
		cmd = append(cmd, "--exclude", pattern)
	}

	for _, file := range o.excludeFiles {
		cmd = append(cmd, "--exclude-file", file)
	}

	if len(o.tags) > 0 {
		cmd = append(cmd, "--tag", strings.Join(o.tags, ","))
	}

	if o.dryRun {
		cmd = append(cmd, "--dry-run")
	}

	if len(o.host) > 0 {
		cmd = append(cmd, "--host", o.host)
	}

	if !o.scan {
		cmd = append(cmd, "--no-scan")
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		return nil, err
	}

	return _parseBackupResult(resultRaw)
}

func _parseBackupResult(result string) (map[string]interface{}, error) {
	terminalMarkers := "\x1b[2K"
	var lines []string

	for _, line := range strings.Split(result, "\n") {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			lines = append(lines, strings.TrimPrefix(trimmedLine, terminalMarkers))
		}
	}

	if len(lines) == 0 {
		return nil, &UnexpectedResticResult{"No valid lines found in the result"}
	}

	lastLine := lines[len(lines)-1]
	var parsedResult map[string]interface{}
	err := json.Unmarshal([]byte(lastLine), &parsedResult)
	if err != nil {
		return nil, &UnexpectedResticResult{fmt.Sprintf("Expected valid JSON response from restic, got %s", result)}
	}

	return parsedResult, nil
}
