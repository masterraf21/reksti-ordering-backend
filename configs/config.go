package configs

// MySQL config
var MySQL *mysql

func init() {
	MySQL = setupMySQL()
}
