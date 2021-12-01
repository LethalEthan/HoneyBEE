package packet

import "HoneyBEE/jsonstruct"

//Play 0x55
type Teams_CB struct {
	TeamName   string
	Mode       byte
	ActionData interface{}
}

type TeamStruct struct {
	TeamDisplayName   *jsonstruct.ChatComponent
	FriendlyFlags     byte
	NameTagVisibility string
	CollisionRule     string
	TeamColour        int32
	TeamPrefix        jsonstruct.ChatComponent
	TeamSuffix        jsonstruct.ChatComponent
	EntityArray       *TeamEntityArray //pointer so I can reuse this struct for mode 2
}

type TeamEntityArray struct {
	EntityCount int32
	Entities    []string
}

const (
	TeamAlways            = "always"
	TeamHideForOtherTeams = "hideForOtherTeams"
	TeamHideForOwnTeam    = "hideForOwnTeam"
	TeamsNever            = "never"
	TeamPushOtherTeam     = "pushOtherTeams"
	TeamPushOwnTeam       = "pushOwnTeam"
)

func (Team Teams_CB) EncodeTeam(PW *PacketWriter) {
	PW.ResetData(0x55)
	PW.WriteString(Team.TeamName)
	PW.WriteUByte(Team.Mode)
	switch Team.Mode {
	case 0:
		CreateTeam := Team.ActionData.(TeamStruct)
		PW.WriteArray(CreateTeam.TeamDisplayName.MarshalChatComponent())
		PW.WriteUByte(CreateTeam.FriendlyFlags)
		PW.WriteString(CreateTeam.NameTagVisibility)
		PW.WriteString(CreateTeam.CollisionRule)
		PW.WriteVarInt(CreateTeam.TeamColour)
		PW.WriteArray(CreateTeam.TeamPrefix.MarshalChatComponent())
		PW.WriteArray(CreateTeam.TeamSuffix.MarshalChatComponent())
		PW.WriteVarInt(CreateTeam.EntityArray.EntityCount)
		PW.WriteStringArray(CreateTeam.EntityArray.Entities)
	case 1:
		return
	case 2:
		UpdateTeam := Team.ActionData.(TeamStruct)
		PW.WriteArray(UpdateTeam.TeamDisplayName.MarshalChatComponent())
		PW.WriteUByte(UpdateTeam.FriendlyFlags)
		PW.WriteString(UpdateTeam.NameTagVisibility)
		PW.WriteString(UpdateTeam.CollisionRule)
		PW.WriteVarInt(UpdateTeam.TeamColour)
		PW.WriteArray(UpdateTeam.TeamPrefix.MarshalChatComponent())
		PW.WriteArray(UpdateTeam.TeamSuffix.MarshalChatComponent())
	case 3:
		Entities := Team.ActionData.(TeamEntityArray)
		PW.WriteVarInt(Entities.EntityCount)
		PW.WriteStringArray(Entities.Entities)
	case 4:
		Entities := Team.ActionData.(TeamEntityArray)
		PW.WriteVarInt(Entities.EntityCount)
		PW.WriteStringArray(Entities.Entities)
	}
}

func CreateTeam() {

}

func UpdateTeam() {

}

func AddEntityToTeam() {

}

func RemoveEntityFromTeam() {

}
