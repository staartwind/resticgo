package resticgo

import "encoding/json"

type StatsOptFunc func(*StatsOpts)

type StatsOpts struct {
	mode string
	tags []string
	host string
}

func StatsWithMode(mode string) StatsOptFunc {
	return func(o *StatsOpts) {
		o.mode = mode
	}
}

func StatsWithTags(tags ...string) StatsOptFunc {
	return func(o *StatsOpts) {
		o.tags = append(o.tags, tags...)
	}
}

func StatsWithHost(host string) StatsOptFunc {
	return func(o *StatsOpts) {
		o.host = host
	}
}

func (r *Restic) Stats(opts ...StatsOptFunc) (map[string]interface{}, error) {
	o := StatsOpts{}
	for _, opt := range opts {
		opt(&o)
	}

	cmd := []string{r.command}
	cmd = append(cmd, r.flags...)
	cmd = append(cmd, "stats")

	if len(o.mode) > 0 {
		cmd = append(cmd, "--mode", o.mode)
	}

	for _, tag := range o.tags {
		cmd = append(cmd, "--tag", tag)
	}

	if len(o.host) > 0 {
		cmd = append(cmd, "--host", o.host)
	}

	resultRaw, err := execute(cmd)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	if err = json.Unmarshal([]byte(resultRaw), &out); err != nil {
		return nil, err
	}

	return out, nil
}
