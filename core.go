package utils

import (
	"fmt"
	stdlog "log"

	"github.com/spf13/viper"

	"github.com/cihub/seelog"
	log "github.com/cihub/seelog"
)

type seeAsExternal struct {
	logger *seelog.LoggerInterface
}

func (l seeAsExternal) Print(m ...interface{}) {
	if l.logger == nil {
		log.Debug(m...)
	} else {
		(*l.logger).Debug(m...)
	}
}
func (l seeAsExternal) Printf(format string, m ...interface{}) {
	(*l.logger).Debugf(format, m...)
}
func (l seeAsExternal) Println(m ...interface{}) {
	(*l.logger).Info(m...)
}
func (l seeAsExternal) Error(m ...interface{}) {
	(*l.logger).Error(m...)
}
func (l seeAsExternal) Warn(m ...interface{}) {
	(*l.logger).Warn(m...)
}
func (l seeAsExternal) Info(m ...interface{}) {
	(*l.logger).Info(m...)
}
func (l seeAsExternal) Debug(m ...interface{}) {
	(*l.logger).Debug(m...)
}
func (l seeAsExternal) Write(p []byte) (n int, err error) {
	(*l.logger).Debug(string(p))
	return len(p), nil
}

// Logger main app logger object; you can use seelog.log instead
var Logger seelog.LoggerInterface

// ExternalLogger can be used were golo.ExternalLogger is required
var ExternalLogger seeAsExternal

// StandardLogger can be used where we need log.Logger
var StandardLogger *stdlog.Logger

// Init should be called on app start
func Init() {
	var err error
	Logger, err = log.LoggerFromConfigAsFile("seelog.xml")
	if err == nil {
		ExternalLogger.logger = &Logger
		log.ReplaceLogger(Logger)
	} else {
		log.Warn("problem while configuring logger: %v", err)
		fmt.Printf("problem while configuring logger: %v", err)
		ExternalLogger.logger = &log.Default
	}
	StandardLogger = stdlog.New(ExternalLogger, "std: ", 0)
}

// GetProperty returns config value with possible default
func GetProperty(name string, def ...string) string {
	ret := viper.GetString(name)
	if ret == "" && len(def) > 0 {
		ret = def[0]
	}
	return ret
}

//GetString is alias for GetProperty
func GetString(name string, def ...string) string {
	return GetProperty(name, def...)
}

//GetObject returns properties object
func GetObject(name string) interface{} {
	return viper.Get(name)
}

//UnmarshalObject tries to unmarshal object fromconfig
func UnmarshalObject(name string, obj interface{}) error {
	return viper.UnmarshalKey(name, obj)
}
