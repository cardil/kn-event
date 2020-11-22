package cli

import (
	"encoding/json"
	"time"

	"github.com/cardil/kn-event/internal/event"
	"github.com/ghodss/yaml"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

func createOptions(opts *Options) *event.Options {
	zc := zap.NewProductionConfig()
	if opts.Verbose {
		zc = zap.NewDevelopmentConfig()
	}
	zc.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	var encoder zapcore.Encoder
	switch opts.Output {
	case HumanReadable:
		encoder = zapcore.NewConsoleEncoder(zc.EncoderConfig)
	case YAML:
		encoder = &yamlEncoder{json: zapcore.NewJSONEncoder(zc.EncoderConfig)}
	case JSON:
		encoder = zapcore.NewJSONEncoder(zc.EncoderConfig)
	}
	sink := zapcore.AddSync(opts.OutWriter)
	errSink := zapcore.AddSync(opts.ErrWriter)
	zcore := zapcore.NewCore(encoder, sink, zc.Level)
	log := zap.New(
		zcore, buildOptions(zc, errSink)...,
	)

	return &event.Options{
		KnPluginOptions: opts.KnPluginOptions,
		Log:             log.Sugar(),
	}
}

func buildOptions(cfg zap.Config, errSink zapcore.WriteSyncer) []zap.Option {
	opts := []zap.Option{zap.ErrorOutput(errSink)}

	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	return opts
}

type yamlEncoder struct {
	json zapcore.Encoder
}

func (y *yamlEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	return y.json.AddArray(key, marshaler)
}

func (y *yamlEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	return y.json.AddObject(key, marshaler)
}

func (y *yamlEncoder) AddBinary(key string, value []byte) {
	y.json.AddBinary(key, value)
}

func (y *yamlEncoder) AddByteString(key string, value []byte) {
	y.json.AddByteString(key, value)
}

func (y *yamlEncoder) AddBool(key string, value bool) {
	y.json.AddBool(key, value)
}

func (y *yamlEncoder) AddComplex128(key string, value complex128) {
	y.json.AddComplex128(key, value)
}

func (y *yamlEncoder) AddComplex64(key string, value complex64) {
	y.json.AddComplex64(key, value)
}

func (y *yamlEncoder) AddDuration(key string, value time.Duration) {
	y.json.AddDuration(key, value)
}

func (y *yamlEncoder) AddFloat64(key string, value float64) {
	y.json.AddFloat64(key, value)
}

func (y *yamlEncoder) AddFloat32(key string, value float32) {
	y.json.AddFloat32(key, value)
}

func (y *yamlEncoder) AddInt(key string, value int) {
	y.json.AddInt(key, value)
}

func (y *yamlEncoder) AddInt64(key string, value int64) {
	y.json.AddInt64(key, value)
}

func (y *yamlEncoder) AddInt32(key string, value int32) {
	y.json.AddInt32(key, value)
}

func (y *yamlEncoder) AddInt16(key string, value int16) {
	y.json.AddInt16(key, value)
}

func (y *yamlEncoder) AddInt8(key string, value int8) {
	y.json.AddInt8(key, value)
}

func (y *yamlEncoder) AddString(key, value string) {
	y.json.AddString(key, value)
}

func (y *yamlEncoder) AddTime(key string, value time.Time) {
	y.json.AddTime(key, value)
}

func (y *yamlEncoder) AddUint(key string, value uint) {
	y.json.AddUint(key, value)
}

func (y *yamlEncoder) AddUint64(key string, value uint64) {
	y.json.AddUint64(key, value)
}

func (y *yamlEncoder) AddUint32(key string, value uint32) {
	y.json.AddUint32(key, value)
}

func (y *yamlEncoder) AddUint16(key string, value uint16) {
	y.json.AddUint16(key, value)
}

func (y *yamlEncoder) AddUint8(key string, value uint8) {
	y.json.AddUint8(key, value)
}

func (y *yamlEncoder) AddUintptr(key string, value uintptr) {
	y.json.AddUintptr(key, value)
}

func (y *yamlEncoder) AddReflected(key string, value interface{}) error {
	return y.json.AddReflected(key, value)
}

func (y *yamlEncoder) OpenNamespace(key string) {
	y.json.OpenNamespace(key)
}

func (y *yamlEncoder) Clone() zapcore.Encoder {
	return &yamlEncoder{
		json: y.json.Clone(),
	}
}

func (y *yamlEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := y.json.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(buf.Bytes(), &v)
	if err != nil {
		return nil, err
	}
	bytes, err := yaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	buf = buffer.NewPool().Get()
	_, _ = buf.Write([]byte("---\n"))
	_, err = buf.Write(bytes)
	return buf, err
}
