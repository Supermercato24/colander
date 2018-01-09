package configuration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	soaPath   = filepath.Join(PathConfig, soaFile)
	soaMockup = []byte(`<?php
return [
    'UA' => 'gdo_SOA',
    'cp24' => 'https://0.bar.it',
    'cp24_token' => 'Basic Z0zo=',
     's24' => 'https://1.bar.it',
    's24_token' => 'Basic Z1zo=',
    'sm' => 'https://2.bar.it',
    'sm_token' => 'Basic Z2zo=',
];`)
)

func SoaModel(t *testing.T, soa *soa) {
	// UA
	assert.NotEmpty(t, soa.UA())

	// cp24
	assert.NotEmpty(t, soa.Cp24.Domain())
	assert.NotEmpty(t, soa.Cp24.Token())

	// s24
	assert.NotEmpty(t, soa.S24.Domain())
	assert.NotEmpty(t, soa.S24.Token())

	// sm
	assert.NotEmpty(t, soa.Sm.Domain())
	assert.NotEmpty(t, soa.Sm.Token())

	// Domain
	assert.NotEqual(t, soa.Cp24.Domain(), soa.S24.Domain())
	assert.NotEqual(t, soa.Cp24.Domain(), soa.Sm.Domain())
	assert.Contains(t, soa.Cp24.Domain(), "https")

	assert.NotEqual(t, soa.S24.Domain(), soa.Cp24.Domain())
	assert.NotEqual(t, soa.S24.Domain(), soa.Sm.Domain())
	assert.Contains(t, soa.S24.Domain(), "https")

	assert.NotEqual(t, soa.Sm.Domain(), soa.S24.Domain())
	assert.NotEqual(t, soa.Sm.Domain(), soa.Cp24.Domain())
	assert.Contains(t, soa.Sm.Domain(), "https")

	// Token
	assert.NotEqual(t, soa.Cp24.Token(), soa.S24.Token())
	assert.NotEqual(t, soa.Cp24.Token(), soa.Sm.Token())
	assert.Contains(t, soa.Cp24.Token(), "Basic")

	assert.NotEqual(t, soa.S24.Token(), soa.Cp24.Token())
	assert.NotEqual(t, soa.S24.Token(), soa.Sm.Token())
	assert.Contains(t, soa.Cp24.Token(), "Basic")

	assert.NotEqual(t, soa.Sm.Token(), soa.S24.Token())
	assert.NotEqual(t, soa.Sm.Token(), soa.Cp24.Token())
	assert.Contains(t, soa.Sm.Token(), "Basic")
}

func SoaMockup(t *testing.T) {
	cwd, _ := os.Getwd()
	DirProject = cwd
	BuildProject()

	fd, err := os.Create(soaPath)
	assert.NoError(t, err)
	fd.Write(soaMockup)
	assert.NoError(t, fd.Close())
}

func TestSoa(t *testing.T) {
	assert.True(t, t.Run("mockup", SoaMockup))
	assert.True(t, assert.NotPanics(t, func() { SoaConfiguration() }))
	assert.NotPanics(t, func() { SoaModel(t, SoaConfiguration()) })
	assert.NoError(t, os.Remove(soaPath))
}

func BenchmarkSoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SoaConfiguration()
	}
}
