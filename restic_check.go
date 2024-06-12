package resticgo

import (
	"errors"
)

type CheckOptFunc func(*CheckOpts)

type CheckOpts struct {
	readData       bool
	readDataSubset string
}

func CheckWithReadData(o *CheckOpts) {
	o.readData = true
}

func CheckWithReadDataSubset(subset string) CheckOptFunc {
	return func(o *CheckOpts) {
		o.readDataSubset = subset
	}
}

func defaultCheckOpts() CheckOpts {
	return CheckOpts{readData: false}
}

func (r *Restic) Check(opts ...CheckOptFunc) (string, error) {
	o := defaultCheckOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "check")

	if o.readData {
		cmd = append(cmd, "--check-data")
	}

	if len(o.readDataSubset) > 0 {
		cmd = append(cmd, "--read-data-subset", o.readDataSubset)
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		var failedError *ResticFailedError
		if errors.As(err, &failedError) {
			return "", nil
		}
		return "", err
	}

	return resultRaw, nil
}
