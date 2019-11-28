package tfsec

import (
	"testing"

	"github.com/liamg/tfsec/internal/app/tfsec/scanner"

	"github.com/liamg/tfsec/internal/app/tfsec/checks"
)

func Test_GkeAbacEnabled(t *testing.T) {

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode scanner.CheckCode
		mustExcludeResultCode scanner.CheckCode
	}{
		{
			name: "check google_container_cluster with enable_legacy_abac set to true",
			source: `
resource "google_container_cluster" "gke" {
	enable_legacy_abac = "true"
	
}`,
			mustIncludeResultCode: checks.GkeAbacEnabled,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := scanSource(test.source)
			assertCheckCode(t, test.mustIncludeResultCode, test.mustExcludeResultCode, results)
		})
	}

}
