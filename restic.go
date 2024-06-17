package resticgo

import (
	"log/slog"
	"os"
)

func init() {
	var level slog.Level
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})))
}

type ResticOptFunc func(*ResticOpts)

type ResticOpts struct {
	command string
	flags   []string
}

func WithRepository(repository string) ResticOptFunc {
	return func(o *ResticOpts) {
		o.flags = append(o.flags, []string{"--repository", repository}...)
	}
}

func WithPasswordFile(passwordFile string) ResticOptFunc {
	return func(o *ResticOpts) {
		o.flags = append(o.flags, []string{"--password-file", passwordFile}...)
	}
}

func WithoutCache(o *ResticOpts) {
	o.flags = append(o.flags, "--no-cache")
}

func WithCustomArgs(args ...string) ResticOptFunc {
	return func(o *ResticOpts) {
		o.flags = append(o.flags, args...)
	}
}

func OverwriteCommand(command string) ResticOptFunc {
	return func(o *ResticOpts) {
		o.command = command
	}
}

func defaultOpts() ResticOpts {
	return ResticOpts{
		command: "restic",
		flags:   []string{"--json"},
	}
}

// Restic provides an instance to call Restic command line actions like backups, snapshots etc...
type Restic struct {
	command string
	flags   []string
}

// NewRestic returns an instance of Restic
// With the Restic instance you can call Restic actions like Backup
// Example:
// resticClient := resticgo.NewRestic()
// resticClient.Backup([]string{"/tmp/test.txt"})
func NewRestic(opts ...ResticOptFunc) *Restic {
	o := defaultOpts()
	for _, opt := range opts {
		opt(&o)
	}
	return &Restic{
		command: o.command,
		flags:   o.flags,
	}
}
