/*
Copyright 2026 The pdfcpu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package format

import (
	"testing"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func TestTextPercent(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		want   string
		unique bool
	}{
		{"TrailingPercent", "100%", "100%", false},
		{"PercentBeforeSpace", "100% done", "100% done", false},
		{"EscapedPercent", "50%% off", "50% off", false},
		{"EscapedDirective", "%%t", "%t", false},
		{"EscapedPercentThenDirective", "%%%P", "%3", true},
		{"PageNr", "%p of %P", "2 of 3", true},
		{"PageNrOffset", "%p10", "12", true},
		{"Version", "%v", model.VersionStr, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, unique := Text(tt.text, "", 2, 3)
			if got != tt.want || unique != tt.unique {
				t.Fatalf("Text(%q) = %q, %t, want %q, %t", tt.text, got, unique, tt.want, tt.unique)
			}
		})
	}
}
