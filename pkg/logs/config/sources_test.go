// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package config

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSource(t *testing.T) {
	sources := NewLogSources([]*LogSource{})
	assert.Equal(t, 0, len(sources.GetSources()))
	sources.AddSource(NewLogSource("foo", nil))
	assert.Equal(t, 1, len(sources.GetSources()))
	sources.AddSource(NewLogSource("foo", nil))
	assert.Equal(t, 2, len(sources.GetSources()))
}

func TestRemoveSource(t *testing.T) {
	source1 := NewLogSource("", nil)
	source2 := NewLogSource("", nil)
	sources := NewLogSources([]*LogSource{source1, source2})
	assert.Equal(t, 2, len(sources.GetSources()))
	sources.RemoveSource(source1)
	assert.Equal(t, 1, len(sources.GetSources()))
	assert.Equal(t, source2, sources.GetSources()[0])
	sources.RemoveSource(source2)
	assert.Equal(t, 0, len(sources.GetSources()))
}

func TestGetSources(t *testing.T) {
	sources := NewLogSources([]*LogSource{})
	assert.Equal(t, 0, len(sources.GetSources()))
	sources = NewLogSources([]*LogSource{NewLogSource("", nil)})
	assert.Equal(t, 1, len(sources.GetSources()))
}

func TestGetValidSources(t *testing.T) {
	source1 := NewLogSource("", nil)
	source2 := NewLogSource("", nil)
	sources := NewLogSources([]*LogSource{source1, source2})
	assert.Equal(t, 2, len(sources.GetValidSources()))
	source1.Status.Error(errors.New("invalid"))
	assert.Equal(t, 1, len(sources.GetValidSources()))
	source1.Status.Success()
	assert.Equal(t, 2, len(sources.GetValidSources()))
}

func TestGetValidSourcesWithType(t *testing.T) {
	source1 := NewLogSource("", nil)
	source2 := NewLogSource("", &LogsConfig{})
	source3 := NewLogSource("", &LogsConfig{Type: "foo"})
	source4 := NewLogSource("", &LogsConfig{Type: "bar"})
	sources := NewLogSources([]*LogSource{source1, source2, source3, source4})
	assert.Equal(t, 1, len(sources.GetValidSourcesWithType("foo")))
}
