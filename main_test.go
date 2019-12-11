package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/brightpuddle/goaci/backup"
	"github.com/stretchr/testify/assert"
)

func TestBackupQuery(t *testing.T) {

	// Class query
	res, _ := backupQuery(Args{
		Target: "./testdata/config.tar.gz",
		Class:  "fvTenant",
	})
	if !assert.Equal(t, 2, len(res.Array())) {
		fmt.Println(res.Get("@pretty"))
	}

	// DN query
	res, _ = backupQuery(Args{
		Target: "./testdata/config.tar.gz",
		Dn:     "uni/tn-a",
	})
	if !assert.Equal(t, "a", res.Get("fvTenant.attributes.name").Str) {
		fmt.Println(res.Get("@pretty"))
	}

	// No class or DN
	_, err := backupQuery(Args{Target: "./testdata/config.tar.gz"})
	assert.Error(t, err)

	// File doesn't exist
	_, err = backupQuery(Args{Target: "does.not.exist"})
	assert.Error(t, err)
}

func TestPrintResult(t *testing.T) {
	// Single object
	client, _ := backup.NewClient("./testdata/config.tar.gz")
	res, _ := client.GetDn("uni/tn-a")
	b := bytes.Buffer{}
	printResult(res, &b)
	assert.Contains(t, b.String(), "uni/tn-a")

	// Array
	res, _ = client.GetClass("fvTenant")
	b = bytes.Buffer{}
	printResult(res, &b)
	assert.Contains(t, b.String(), "uni/tn-a")
	assert.Contains(t, b.String(), "uni/tn-b")
}
