package log

import "github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"

type LoggerMock struct{}

func (l LoggerMock) Infof(_ string, _ ...any)                 {}
func (l LoggerMock) Warnf(_ string, _ ...any)                 {}
func (l LoggerMock) Errorf(_ string, _ ...any)                {}
func (l LoggerMock) Fatalln(_ ...any)                         {}
func (l LoggerMock) WithFields(_ logger.Fields) logger.Logger { return LoggerEntryMock{} }
func (l LoggerMock) WithError(_ error) logger.Logger          { return LoggerEntryMock{} }

type LoggerEntryMock struct{}

func (l LoggerEntryMock) Infof(_ string, _ ...any)                 {}
func (l LoggerEntryMock) Warnf(_ string, _ ...any)                 {}
func (l LoggerEntryMock) Errorf(_ string, _ ...any)                {}
func (l LoggerEntryMock) Fatalln(_ ...any)                         {}
func (l LoggerEntryMock) WithFields(_ logger.Fields) logger.Logger { return LoggerEntryMock{} }
func (l LoggerEntryMock) WithError(_ error) logger.Logger          { return LoggerEntryMock{} }
