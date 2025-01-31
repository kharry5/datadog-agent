// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !ec2

package ec2

import "context"

// GetTags grabs the host tags from the EC2 api
//
//nolint:revive // TODO(ASC) Fix revive linter
func GetTags(ctx context.Context) ([]string, error) {
	return []string{}, nil
}

//nolint:revive // TODO(ASC) Fix revive linter
func fetchTagsFromCache(ctx context.Context) ([]string, error) {
	return []string{}, nil
}
