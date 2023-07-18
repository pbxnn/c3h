package log

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/durationpb"
)

var _ log.Logger = (*LogrusLogger)(nil)

var defaultRotateTime = time.Hour * 24
var defaultMaxAge = time.Hour * 24 * 7

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(options ...Option) log.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	for _, option := range options {
		option(logger)
	}

	return &LogrusLogger{
		log: logger,
	}
}

func (l *LogrusLogger) Log(level log.Level, keyvals ...interface{}) (err error) {
	var (
		logrusLevel logrus.Level
		fields      logrus.Fields = make(map[string]interface{})
		msg         string
	)

	switch level {
	case log.LevelDebug:
		logrusLevel = logrus.DebugLevel
	case log.LevelInfo:
		logrusLevel = logrus.InfoLevel
	case log.LevelWarn:
		logrusLevel = logrus.WarnLevel
	case log.LevelError:
		logrusLevel = logrus.ErrorLevel
	default:
		logrusLevel = logrus.DebugLevel
	}

	if logrusLevel > l.log.Level {
		return
	}

	if len(keyvals) == 0 {
		return nil
	}
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "")
	}
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue
		}
		if key == logrus.FieldKeyMsg {
			msg, _ = keyvals[i+1].(string)
			continue
		}
		fields[key] = keyvals[i+1]
	}

	if len(fields) > 0 {
		l.log.WithFields(fields).Log(logrusLevel, msg)
	} else {
		l.log.Log(logrusLevel, msg)
	}

	return
}

type Option func(log *logrus.Logger)

func Level(level string) Option {
	return func(log *logrus.Logger) {
		switch strings.ToLower(level) {
		case "debug":
			log.Level = logrus.DebugLevel
		case "info":
			log.Level = logrus.InfoLevel
		case "warning":
			log.Level = logrus.WarnLevel
		case "error":
			log.Level = logrus.ErrorLevel
		default:
			log.Level = logrus.InfoLevel
		}
	}
}

func Output(w io.Writer) Option {
	return func(log *logrus.Logger) {
		log.Out = w
	}
}

func Formatter(formatter logrus.Formatter) Option {
	return func(log *logrus.Logger) {
		log.Formatter = formatter
	}
}

func FileOutput(filename string, rotateTime, maxAge *durationpb.Duration) Option {
	return func(log *logrus.Logger) {
		rt := defaultRotateTime
		if rotateTime != nil {
			rt = rotateTime.AsDuration()
		}

		ma := defaultMaxAge
		if maxAge != nil {
			ma = maxAge.AsDuration()
		}

		logf, err := rotatelogs.New(
			filename,
			rotatelogs.WithLinkName(filename+".%Y&m%d%H"),
			rotatelogs.WithRotationTime(rt),
			rotatelogs.WithMaxAge(ma),
		)

		if err != nil {
			panic(err)
		}

		log.SetOutput(io.MultiWriter(logf, os.Stdout))
	}
}
