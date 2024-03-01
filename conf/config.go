package conf

type Conf struct {
	Port  string `yaml:"port"`
	Env   string `yaml:"env"`
	MySql mysql  `yaml:"mysql"`
	JWT   jwt    `yaml:"jwt"`
}

type mysql struct {
	DSN string `yaml:"dsn"`
}

type jwt struct {
	Secret string `yaml:"secret"`
}
