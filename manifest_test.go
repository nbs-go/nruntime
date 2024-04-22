package nruntime_test

import (
	"github.com/nbs-go/nruntime"
	"testing"
	"time"
)

const (
	// Define Application Manifest
	appName        = "TestApp"
	appVersion     = "v1.0.0"
	buildSignature = "9f816a75ca6fcd264c791a70a86d8a4a498cf22d"
)

func init() {
	nruntime.NewRuntime(appName, appVersion, buildSignature)
}

func TestManifestAppName(t *testing.T) {
	actual := nruntime.AppName()
	if actual != appName {
		t.Errorf("invalid AppName returned. Expected=%s Actual=%s", appName, actual)
	}
}

func TestManifestAppVersion(t *testing.T) {
	actual := nruntime.AppVersion()
	if actual != appVersion {
		t.Errorf("invalid AppVersion returned. Expected=%s Actual=%s", appVersion, actual)
	}
}

func TestManifestBuildSignature(t *testing.T) {
	actual := nruntime.BuildSignature()
	if actual != buildSignature {
		t.Errorf("invalid BuildSignature returned. Expected=%s Actual=%s", buildSignature, actual)
	}
}

func TestManifestStartedAt(t *testing.T) {
	actual := nruntime.StartedAt()
	now := time.Now()
	if actual.After(now) {
		t.Errorf("Unexpected captured StartedAt must be less than now. Now=%s Actual=%s", now, actual)
	}
}
