// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022-present Datadog, Inc.

//go:build test

package retry

//nolint:revive // TODO(ASC) Fix revive linter
func NewPointCountTelemetryMock() *PointCountTelemetry {
	return NewPointCountTelemetry("domain")
}
