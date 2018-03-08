package manager

import "sync"

type Manage struct {
}

var manage *Manage
var once sync.Once

func GetManage() *Manage {
	once.Do(func() {
		manage = &Manage{}
	})
	return manage
}
