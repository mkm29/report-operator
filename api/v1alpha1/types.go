package v1alpha1

import (
	"fmt"
	"strings"
)

const (
	TTLReportAnnotation = "trivy-operator.aquasecurity.github.io/report-ttl"
)

// Severity level of a vulnerability or a configuration audit check.
// +enum
type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityHigh     Severity = "HIGH"
	SeverityMedium   Severity = "MEDIUM"
	SeverityLow      Severity = "LOW"

	SeverityUnknown Severity = "UNKNOWN"
)

// StringToSeverity returns the enum constant of Severity with the specified
// name. The name must match exactly an identifier used to declare an enum
// constant. (Extraneous whitespace characters are not permitted.)
func StringToSeverity(name string) (Severity, error) {
	s := strings.ToUpper(name)
	switch s {
	case "CRITICAL", "HIGH", "MEDIUM", "LOW", "NONE", "UNKNOWN":
		return Severity(s), nil
	default:
		return "", fmt.Errorf("unrecognized name literal: %s", name)
	}
}

const ScannerNameTrivy = "Trivy"

// Scanner is the spec for a scanner generating a security assessment report.
type Scanner struct {
	// Name the name of the scanner.
	Name string `json:"name"`

	// Vendor the name of the vendor providing the scanner.
	Vendor string `json:"vendor"`

	// Version the version of the scanner.
	Version string `json:"version"`
}

type SnykCvssDetails struct {
	Assigner        string  `deepcopier:"assigner"`
	Severity        string  `deepcopier:"severity"`
	CvssV3Vector    string  `deepcopier:"V3Vector"`
	CvssV3BaseScore float64 `deepcopier:"V3Score"`
}

type SnykVulnerability struct {
	ID          string            `deepcopier:"vulnerabilityID"`
	Title       string            `deepcopier:"title"`
	CVSSv3      float64           `deepcopier:"score"`
	CvssDetails []SnykCvssDetails `deepcopier:"cvss.nvd"`
	Description string            `deepcopier:"description"`
	Severity    Severity          `deepcopier:"severity"`
	Language    string            `deepcopier:"language"`
	Package     string            `deepcopier:"package"`
	Version     string            `deepcopier:"version"`
	FixedIn     string            `deepcopier:"fixedIn"`
	References  []string          `deepcopier:"links"`
}

type SnykReport struct {
	// Vulnerabilities is a list of vulnerabilities found in the scanned image.
	Vulnerabilities []SnykVulnerability `deepcopier:"vulnerabilities"`
}
