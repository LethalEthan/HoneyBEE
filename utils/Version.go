package utils

import "fmt"

const (
	PrimaryMinecraftVersion         = "1.17.1" //Primary Supported MC version
	PrimaryMinecraftProtocolVersion = 756      //Primary Supported MC protocol Version
	//HoneyBEEVersion                       = "1.1.1 (Build 141)"
	MajorVersion = 0
	MinorVersion = 0
	PatchVersion = 1
	BuildVersion = 13
	StateVersion = "ALPHA" //PreAlpha, Alpha, Beta, Release Candidate, Release
)

//GetVersion - Returns Major, Minor, Patch, VersionState and Build number
func GetVersion() (int, int, int, int, string) {
	return MajorVersion, MinorVersion, PatchVersion, BuildVersion, StateVersion
}

//GetVersionString - Returns stringified version
func GetVersionString() string {
	version := fmt.Sprintf("%d.%d.%d.%d STATE: %s", MajorVersion, MinorVersion, PatchVersion, BuildVersion, StateVersion)
	return version
}
