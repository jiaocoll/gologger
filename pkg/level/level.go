package level

/*
创建人员：云深不知处
创建时间：2022/2/3
程序功能：日志等级
*/

type Thelevel int

const (
	Fatal Thelevel = iota
	Info
	Warning
	Error
	Debug
	Verbose
)

func (L Thelevel) Level() string {
	return [...]string{"FATA", "INFO", "WARN", "ERRO", "DEBG", "VEBO"}[L]
}
