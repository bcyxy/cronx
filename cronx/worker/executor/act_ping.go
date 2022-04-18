package executor

import "context"

func init() {
	act := actPing{}
	actMap[act.GetName()] = &act
}

type actPing struct{}

func (sf *actPing) GetName() string {
	return "PING"
}

func (sf *actPing) Do(ctx context.Context, obj, param string) (ret map[string]string, err error) {
	return
}
