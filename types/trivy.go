package types

import "time"

// types used with trivy scanner

type CVEReport struct {
	SchemaVersion int       `json:"SchemaVersion"`
	ArtifactName  string    `json:"ArtifactName"`
	ArtifactType  string    `json:"ArtifactType"`
	Metadata      Metadata  `json:"Metadata"`
	Results       []Results `json:"Results"`
}

type Rootfs struct {
	Type    string      `json:"type"`
	DiffIds interface{} `json:"diff_ids"`
}

type Config struct {
}

type ImageConfig struct {
	Architecture string    `json:"architecture"`
	Created      time.Time `json:"created"`
	Os           string    `json:"os"`
	Rootfs       Rootfs    `json:"rootfs"`
	Config       Config    `json:"config"`
}

type Metadata struct {
	ImageConfig ImageConfig `json:"ImageConfig"`
}
type Layer struct {
}

type Nvd struct {
	V2Vector string  `json:"V2Vector"`
	V3Vector string  `json:"V3Vector"`
	V2Score  float64 `json:"V2Score"`
	V3Score  float64 `json:"V3Score"`
}

type Redhat struct {
	V3Vector string  `json:"V3Vector"`
	V3Score  float64 `json:"V3Score"`
}

type Cvss struct {
	Nvd    Nvd    `json:"nvd"`
	Redhat Redhat `json:"redhat"`
}

type Vulnerabilities struct {
	VulnerabilityID  string    `json:"VulnerabilityID"`
	PkgName          string    `json:"PkgName"`
	InstalledVersion string    `json:"InstalledVersion"`
	FixedVersion     string    `json:"FixedVersion"`
	Layer            Layer     `json:"Layer"`
	SeveritySource   string    `json:"SeveritySource"`
	PrimaryURL       string    `json:"PrimaryURL"`
	Title            string    `json:"Title"`
	Description      string    `json:"Description"`
	Severity         Severity  `json:"Severity"`
	Cvss             Cvss      `json:"CVSS"`
	References       []string  `json:"References"`
	PublishedDate    time.Time `json:"PublishedDate"`
	LastModifiedDate time.Time `json:"LastModifiedDate"`
}

type Results struct {
	Target          string            `json:"Target"`
	Class           string            `json:"Class"`
	Type            string            `json:"Type"`
	Vulnerabilities []Vulnerabilities `json:"Vulnerabilities"`
}
