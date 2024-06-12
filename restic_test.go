package resticgo

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestNewRestic(t *testing.T) {
	restic := NewRestic(WithRepository("kaas"), WithoutCache)
	fmt.Println(restic.command)
	fmt.Println(restic.flags)
}

func TestRestic_Snapshots(t *testing.T) {
	restic := NewRestic(WithoutCache)
	res, err := restic.Snapshots()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Backup(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}

	restic := NewRestic(WithoutCache)
	res, err := restic.Backup([]string{filepath.Join(homedir, "Downloads/api")}, BackupWithDryRun)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Check(t *testing.T) {
	restic := NewRestic()
	res, err := restic.Check()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Find(t *testing.T) {
	restic := NewRestic()
	res, err := restic.Find("errors.go")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Forget(t *testing.T) {
	restic := NewRestic()
	res, err := restic.Forget(ForgetWithKeepHourly("1"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Init(t *testing.T) {
	restic := NewRestic()
	res, err := restic.Init()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Stats(t *testing.T) {
	restic := NewRestic()
	res, err := restic.Stats()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Version(t *testing.T) {
	restic := NewRestic()
	res, err := restic.Version()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
