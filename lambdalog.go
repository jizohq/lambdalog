// MIT License
//
// Copyright (c) 2020 Jizo, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package lambdalog

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"go.uber.org/zap"
)

const (
	ServiceNameKey     string = "service.name"
	AWSRequestID       string = "aws.request_id"
	AWSFunctionName    string = "aws.function_name"
	AWSFunctionVersion string = "aws.function_version"
	AWSLogGroupName    string = "aws.log_group_name"
	AWSLogStreamName   string = "aws.log_stream_name"
)

type Logger struct {
	SugaredLogger *zap.SugaredLogger
	ZapLogger     *zap.Logger
}

// New returns new `lambdalog.Logger` that sends log events to a zap.Logger.
func New(name string, ctx lambdacontext.LambdaContext) (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{
		SugaredLogger: logger.WithOptions(zap.AddCallerSkip(1)).Sugar().With(
			ServiceNameKey, name,
			AWSRequestID, ctx.AwsRequestID,
			AWSFunctionName, lambdacontext.FunctionName,
			AWSFunctionVersion, lambdacontext.FunctionVersion,
			AWSLogGroupName, lambdacontext.LogGroupName,
			AWSLogStreamName, lambdacontext.LogStreamName,
		),
		ZapLogger: logger.With(
			zap.String(ServiceNameKey, name),
			zap.String(AWSRequestID, ctx.AwsRequestID),
			zap.String(AWSFunctionName, lambdacontext.FunctionName),
			zap.String(AWSFunctionVersion, lambdacontext.FunctionVersion),
			zap.String(AWSLogGroupName, lambdacontext.LogGroupName),
			zap.String(AWSLogStreamName, lambdacontext.LogStreamName),
		),
	}, nil
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() error {
	return l.SugaredLogger.Sync()
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(i ...interface{}) {
	l.SugaredLogger.Debug(i...)
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(i ...interface{}) {
	l.SugaredLogger.Info(i...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(i ...interface{}) {
	l.SugaredLogger.Warn(i...)
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(i ...interface{}) {
	l.SugaredLogger.Error(i...)
}

// Info uses fmt.Sprint to construct and log a message,
// as an alternative to Info.
func (l *Logger) Log(i ...interface{}) error {
	l.SugaredLogger.Info(i...)
	return nil
}
