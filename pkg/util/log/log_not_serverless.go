// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !serverless

package log

// DebugServerless logs at the debug level only in a serverless context
// no-op in a non serverless context
//
//nolint:revive // TODO(ASC) Fix revive linter
func DebugServerless(v ...interface{}) {
}

// DebugfServerless logs with format at the debug level only in a serverless context
// no-op in a non serverless context
//
//nolint:revive // TODO(ASC) Fix revive linter
func DebugfServerless(format string, params ...interface{}) {
}
