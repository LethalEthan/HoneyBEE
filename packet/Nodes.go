package packet

type NodeFormat struct {
	Flags          byte
	ChildrenCount  int32
	Children       []int32
	RedirectNode   int32  //Opt
	Name           string //Opt
	Parser         string //Opt
	Properties     string //Opt
	SuggestionType string //Opt - only if Flags & 0x10
}
