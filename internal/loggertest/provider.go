// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package loggertest

import (
	"context"
	"io"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-log/internal/logging"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
)

func ProviderRoot(ctx context.Context, output io.Writer) context.Context {
	return tfsdklog.NewRootProviderLogger(
		ctx,
		logging.WithoutLocation(),
		logging.WithoutTimestamp(),
		logging.WithOutput(output),
	)
}

// ProviderRootWithLevel creates a provider root logger at the given level.
// Use this in tests that need to verify behaviour at a specific log level.
func ProviderRootWithLevel(ctx context.Context, output io.Writer, level hclog.Level) context.Context {
	return tfsdklog.NewRootProviderLogger(
		ctx,
		logging.WithoutLocation(),
		logging.WithoutTimestamp(),
		logging.WithOutput(output),
		tfsdklog.WithLevel(level),
	)
}

// ProviderRootWithLocation is for testing code that affects go-hclog's caller
// information (location offset). Most testing code should avoid this, since
// correctly checking differences including the location is extra effort
// with little benefit.
func ProviderRootWithLocation(ctx context.Context, output io.Writer) context.Context {
	return tfsdklog.NewRootProviderLogger(
		ctx,
		logging.WithoutTimestamp(),
		logging.WithOutput(output),
	)
}
