package resticgo

import (
	"encoding/json"
	"strings"
)

type ForgetOptFunc func(*ForgetOpts)

type ForgetOpts struct {
	dryRun      bool
	groupBy     string
	tags        []string
	host        string
	keepLast    string
	keepHourly  string
	keepDaily   string
	keepWeekly  string
	keepMonthly string
	keepYearly  string
	keepWithin  string
	prune       bool
}

func ForgetWithDryRun(o *ForgetOpts) {
	o.dryRun = true
}

func ForgetWithGroupBy(groupBy string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.groupBy = groupBy
	}
}

func ForgetWithTags(tags ...string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.tags = append(o.tags, tags...)
	}
}

func ForgetWithHost(host string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.host = host
	}
}

func ForgetWithKeepLast(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepLast = keep
	}
}

func ForgetWithKeepHourly(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepHourly = keep
	}
}

func ForgetWithKeepDaily(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepDaily = keep
	}
}

func ForgetWithKeepWeekly(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepWeekly = keep
	}
}

func ForgetWithKeepMonthly(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepMonthly = keep
	}
}

func ForgetWithKeepYearly(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepYearly = keep
	}
}

func ForgetWithKeepWithin(keep string) ForgetOptFunc {
	return func(o *ForgetOpts) {
		o.keepWithin = keep
	}
}

func ForgetWithPrune(o *ForgetOpts) {
	o.prune = true
}

func defaultForgetOpts() ForgetOpts {
	return ForgetOpts{
		dryRun: false,
		prune:  false,
	}
}

func (r *Restic) Forget(opts ...ForgetOptFunc) ([]map[string]interface{}, error) {
	o := defaultForgetOpts()
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "forget")

	if o.dryRun {
		cmd = append(cmd, "--dry-run")
	}

	if len(o.groupBy) > 0 {
		cmd = append(cmd, "--group-by", o.groupBy)
	}

	if len(o.tags) > 0 {
		cmd = append(cmd, "--tag", strings.Join(o.tags, ","))
	}

	if len(o.host) > 0 {
		cmd = append(cmd, "--host", o.host)
	}

	if len(o.keepLast) > 0 {
		cmd = append(cmd, "--keep-last", o.keepLast)
	}
	if len(o.keepHourly) > 0 {
		cmd = append(cmd, "--keep-hourly", o.keepHourly)
	}
	if len(o.keepDaily) > 0 {
		cmd = append(cmd, "--keep-daily", o.keepDaily)
	}
	if len(o.keepWeekly) > 0 {
		cmd = append(cmd, "--keep-weekly", o.keepWeekly)
	}
	if len(o.keepMonthly) > 0 {
		cmd = append(cmd, "--keep-monthly", o.keepMonthly)
	}
	if len(o.keepYearly) > 0 {
		cmd = append(cmd, "--keep-yearly", o.keepYearly)
	}
	if len(o.keepWithin) > 0 {
		cmd = append(cmd, "--keep-within", o.keepWithin)
	}

	if o.prune {
		cmd = append(cmd, "--prune")
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		return nil, err
	}

	return _parseForgetResult(resultRaw)
}

func _parseForgetResult(result string) ([]map[string]interface{}, error) {
	resultLines := strings.Split(result, "\n")
	if len(resultLines) > 0 && resultLines[0] != "" {
		var jsonData []map[string]interface{}
		if err := json.Unmarshal([]byte(resultLines[0]), &jsonData); err != nil {
			return nil, &UnexpectedResticResult{
				"Unexpected result from restic. Expected JSON, got: " + resultLines[0],
			}
		}
		return jsonData, nil
	}
	return nil, nil
}
