package jobconf

type confDb interface {
	Init() error
	Load() (bool, error)
}

var cfgDb confDb = &mysqlDb{}

func Init() error {
	return cfgDb.Init()
}

func LoadConf() (bool, error) {
	return cfgDb.Load()
}

func GenJobs() (interface{}, error) {
	return nil, nil
}
