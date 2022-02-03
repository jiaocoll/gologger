package format

import (
	"fmt"
	"github.com/fatih/color"
	"gologger/pkg/level"
	"time"
)

/*
创建人员：云深不知处
创建时间：2022/2/3
程序功能：日志格式
*/

type Logger struct {
	Level level.Thelevel
}

func (L *Logger) INFO(args ...string) {
	_, _ = fmt.Fprintln(color.Output, color.HiCyanString("[INFO]"), "["+time.Now().Format("2006-01-02 15:04:05")+"]", SlinceToStr(args))
}

func (L *Logger) WARN(args ...string) {
	_, _ = fmt.Fprintln(color.Output, color.HiYellowString("[WARN]"), "["+time.Now().Format("2006-01-02 15:04:05")+"]", SlinceToStr(args))
}

func (L *Logger) ERRO(args ...string) {
	_, _ = fmt.Fprintln(color.Output, color.HiRedString("[ERRO]"), "["+time.Now().Format("2006-01-02 15:04:05")+"]", SlinceToStr(args))
}

func (L *Logger) DEBG(args ...string) {
	_, _ = fmt.Fprintln(color.Output, color.HiGreenString("[DEBG]"), "["+time.Now().Format("2006-01-02 15:04:05")+"]", SlinceToStr(args))
}

func (L *Logger) VEBO(args ...string) {
	_, _ = fmt.Fprintln(color.Output, color.HiMagentaString("[VEBO]"), "["+time.Now().Format("2006-01-02 15:04:05")+"]", SlinceToStr(args))
}

func (L *Logger) FATA(args ...string) {
	_, _ = fmt.Fprintln(color.Output, color.HiRedString("[FATA]"), "["+time.Now().Format("2006-01-02 15:04:05")+"]", SlinceToStr(args))
}

func SlinceToStr(SourceData []string) string {
	var res string
	for k, v := range SourceData {
		if k == 0 {
			res += fmt.Sprintf("%s ", v)
		} else {
			res += fmt.Sprintf("%s ", v)
		}
	}
	return res
}
