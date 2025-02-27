package resticgo

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	godotenv.Load()

	m.Run()
}

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

func TestRestic_Cat(t *testing.T) {
	restic := NewRestic(WithoutCache)
	res, err := restic.Cat("snapshot", "9feda324")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_Ls(t *testing.T) {
	restic := NewRestic(WithoutCache)
	res, err := restic.Ls("9feda324")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestRestic_ListLocks(t *testing.T) {
	restic := NewRestic(WithoutCache)
	res, err := restic.ListLocks()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
