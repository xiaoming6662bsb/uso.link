package log2

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().CallerWithSkipFrameCount(3).Logger()

func Der(v error) {
	if v != nil {
		logger.Fatal().Err(v).Msg("Error occurred")
	}
}

func Dd(v ...interface{}) {
	logger.Debug().Msgf(formatArgs(v), v...)
}

func Debugf(f string, v ...interface{}) {
	logger.Debug().Msgf(f, v...)
}

func Errorf(f string, v ...interface{}) {
	logger.Error().Msgf(f, v...)
}

func IfErrPrt(v error) bool {
	if v == nil {
		return false
	}
	logger.Error().Err(v).Msg("Error occurred")
	return true
}

func Logf(f string, v ...interface{}) {
	logger.Info().Msgf(f, v...)
}

type logHelper struct {
	prefix string
}

func (l *logHelper) Write(p []byte) (n int, err error) {
	logger.Info().Str("prefix", l.prefix).Msg(string(p))
	return len(p), nil
}

func newLogHelper(prefix string) *logHelper {
	return &logHelper{prefix: prefix}
}

func formatArgs(v ...interface{}) string {
	ft := ""
	for range v {
		ft += "%v,"
	}
	return ft[:len(ft)-1] // Remove the trailing comma
}
