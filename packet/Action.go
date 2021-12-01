package packet

import "HoneyBEE/jsonstruct"

type ActionComponet struct {
	ActionID int
}
type ActionAdd struct {
	Title    jsonstruct.ChatComponent
	Health   float32
	Colour   int
	Division int
}

/*Division Field
0	No division
1	6 notches
2	10 notches
3	12 notches
4	20 notches
*/
