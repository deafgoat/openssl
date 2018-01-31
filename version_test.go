// Copyright (C) MongoDB, Inc. 2018-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package openssl

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v := Version
	if len(v) == 0 {
		t.Fatal("Version string is empty")
	}
	t.Logf("Testing with %s", Version)
}
