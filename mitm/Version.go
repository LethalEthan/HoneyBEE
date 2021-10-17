package mitm

//ProtocolToVer - protocol/version map
var ProtocolToVer map[int32]string

//ProtocolToVersionInit - Initialise protocol/version map
func ProtocolToVersionInit() {
	ProtocolToVer = map[int32]string{
		107: "1.9",
		108: "1.9.1",
		109: "1.9.2",
		110: "1.9.3/1.9.4",
		210: "1.10/1.10.1/1.10.2",
		315: "1.11",
		316: "1.11.1/1.11.2",
		335: "1.12",
		338: "1.12.1",
		340: "1.12.2",
		393: "1.13",
		401: "1.13.1",
		404: "1.13.2", //not found hehe
		477: "1.14",
		480: "1.14.1",
		485: "1.14.2",
		490: "1.14.3",
		498: "1.14.4",
		573: "1.15",
		575: "1.15.1",
		578: "1.15.2",
		735: "1.16",
		736: "1.16.1",
		751: "1.16.2",
		753: "1.16.3",
		754: "1.16.4/1.16.5",
		755: "1.17",
		756: "1.17.1",
	}
}
