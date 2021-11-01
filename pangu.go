package redis

import `github.com/pangum/pangu`

func init() {
	if err := pangu.New().Provides(newRedis); nil != err {
		panic(err)
	}
}
