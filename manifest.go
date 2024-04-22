package nruntime

import "time"

// Application manifest

var mAppName string
var mAppVersion string
var mBuildSignature string
var mStartedAt time.Time

// NewRuntime set Application Manifest values and capture runtime starting time
func NewRuntime(appName, appVersion, buildSignature string) {
	mAppName = appName
	mAppVersion = appVersion
	mBuildSignature = buildSignature
	mStartedAt = time.Now()
}

// AppName get Application Name value
func AppName() string {
	return mAppName
}

// AppVersion get Application Version value
func AppVersion() string {
	return mAppVersion
}

// BuildSignature get Build Signature value
func BuildSignature() string {
	return mBuildSignature
}

// StartedAt get Application runtime starting time
func StartedAt() time.Time {
	return mStartedAt
}
