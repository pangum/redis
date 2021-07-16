package redis

import `github.com/storezhang/pangu`

func init() {
	if err := pangu.New().Provides(newRedis); nil != err {
		panic(err)
	}
}
