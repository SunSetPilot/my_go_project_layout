package log

//
//import (
//	"bytes"
//	"fmt"
//	"io"
//	"os"
//	"path/filepath"
//	"strconv"
//	"strings"
//	"time"
//
//	"github.com/gin-gonic/gin"
//	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
//	"github.com/sirupsen/logrus"
//
//)
//
//var logLevelMapping = map[string]logrus.Level{
//	"panic": logrus.PanicLevel,
//	"fatal": logrus.FatalLevel,
//	"error": logrus.ErrorLevel,
//	"warn":  logrus.WarnLevel,
//	"info":  logrus.InfoLevel,
//	"debug": logrus.DebugLevel,
//	"trace": logrus.TraceLevel,
//}
//
//const (
//	red    = 31
//	yellow = 33
//	blue   = 36
//	gray   = 37
//)
//
//type MyFormatter struct {
//	levelTextMaxLength int
//}
//
//func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
//	var (
//		levelColor int
//		buffer     *bytes.Buffer
//	)
//	switch entry.Level {
//	case logrus.DebugLevel, logrus.TraceLevel:
//		levelColor = gray
//	case logrus.WarnLevel:
//		levelColor = yellow
//	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
//		levelColor = red
//	case logrus.InfoLevel:
//		levelColor = blue
//	default:
//		levelColor = blue
//	}
//
//	levelText := strings.ToUpper(entry.Level.String())
//	formatString := "[%-" + strconv.Itoa(7) + "s]"
//	levelText = fmt.Sprintf(formatString, levelText)
//
//	if entry.Buffer != nil {
//		buffer = entry.Buffer
//	} else {
//		buffer = &bytes.Buffer{}
//	}
//
//	timestamp := entry.Time.Format("2006-01-02 15:04:05")
//	if entry.HasCaller() {
//		fName := filepath.Base(entry.Caller.File)
//		fmt.Fprintf(buffer, "\x1b[0m[%s] \x1b[%dm%s [%s:%d] %-44s \x1b[0m \n",
//			timestamp, levelColor, levelText, fName, entry.Caller.Line, entry.Message)
//	} else {
//		fmt.Fprintf(buffer, "\x1b[0m[%s] \x1b[%dm%s %-44s \u001B[0m \n",
//			timestamp, levelColor, levelText, entry.Message)
//	}
//
//	return buffer.Bytes(), nil
//}
//
//var Runtime RuntimeLogger
//
//type RuntimeLogger struct {
//	logrus.Logger
//	logWriter *rotatelogs.RotateLogs
//}
//
//func (logger *RuntimeLogger) Init(logFile string) error {
//	var (
//		writers []io.Writer
//		err     error
//	)
//
//	logger.SetLevel(logLevelMapping[global.ProjectConfig.CommonConf.LogLevel])
//	logger.SetReportCaller(true)
//	logger.SetFormatter(&MyFormatter{})
//
//	logger.logWriter, err = rotatelogs.New(
//		logFile+".%Y%m%d%H",
//		rotatelogs.WithLinkName(logFile),
//		rotatelogs.WithMaxAge(7*24*time.Hour),
//		rotatelogs.WithRotationTime(time.Hour),
//	)
//	if err != nil {
//		return err
//	}
//	writers = append(writers, logger.logWriter)
//
//	if logLevelMapping[global.ProjectConfig.CommonConf.LogLevel] >= logrus.InfoLevel {
//		writers = append(writers, os.Stdout)
//	}
//	logWriter := io.MultiWriter(writers...)
//	logger.SetOutput(logWriter)
//	gin.DefaultWriter = logWriter
//	return nil
//}
//
//func (logger *RuntimeLogger) CloseLog() {
//	logger.Exit(0)
//	_ = logger.logWriter.Close()
//}
