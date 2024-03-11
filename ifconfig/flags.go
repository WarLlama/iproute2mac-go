package ifconfig

type Flag struct {
	Name    string
	Bitmask int
}

const (
	FlagUP               = "UP"
	FlagValueUP          = 0x1
	FlagBROADCAST        = "BROADCAST"
	FlagValueBROADCAST   = 0x2
	FlagDEBUG            = "DEBUG"
	FlagValueDEBUG       = 0x4
	FlagLOOPBACK         = "LOOPBACK"
	FlagValueLOOPBACK    = 0x8
	FlagPOINTOPOINT      = "POINTOPOINT"
	FlagValuePOINTOPOINT = 0x10
	FlagNOTRAILERS       = "SMART"
	FlagValueNOTRAILERS  = 0x20
	FlagRUNNING          = "RUNNING"
	FlagValueRUNNING     = 0x40
	FlagNOARP            = "NOARP"
	FlagValueNOARP       = 0x80
	FlagPROMISC          = "PROMISC"
	FlagValuePROMISC     = 0x100
	FlagALLMULTI         = "ALLMULTI"
	FlagValueALLMULTI    = 0x200
	FlagOACTIVE          = "OACTIVE"
	FlagValueOACTIVE     = 0x400
	FlagSIMPLEX          = "SIMPLEX"
	FlagValueSIMPLEX     = 0x800
	FlagLINK0            = "LINK0"
	FlagValueLINK0       = 0x1000
	FlagLINK1            = "LINK1"
	FlagValueLINK1       = 0x2000
	FlagLINK2            = "LINK2"
	FlagValueLINK2       = 0x4000
	FlagMULTICAST        = "MULTICAST"
	FlagValueMULTICAST   = 0x8000
)

var (
	flags = []Flag{
		{
			Name:    FlagUP,
			Bitmask: FlagValueUP,
		},
		{
			Name:    FlagBROADCAST,
			Bitmask: FlagValueBROADCAST,
		},
		{
			Name:    FlagDEBUG,
			Bitmask: FlagValueDEBUG,
		},
		{
			Name:    FlagLOOPBACK,
			Bitmask: FlagValueLOOPBACK,
		},
		{
			Name:    FlagPOINTOPOINT,
			Bitmask: FlagValuePOINTOPOINT,
		},
		{
			Name:    FlagNOTRAILERS,
			Bitmask: FlagValueNOTRAILERS,
		},
		{
			Name:    FlagRUNNING,
			Bitmask: FlagValueRUNNING,
		},
		{
			Name:    FlagNOARP,
			Bitmask: FlagValueNOARP,
		},
		{
			Name:    FlagPROMISC,
			Bitmask: FlagValuePROMISC,
		},
		{
			Name:    FlagALLMULTI,
			Bitmask: FlagValueALLMULTI,
		},
		{
			Name:    FlagOACTIVE,
			Bitmask: FlagValueOACTIVE,
		},
		{
			Name:    FlagSIMPLEX,
			Bitmask: FlagValueSIMPLEX,
		},
		{
			Name:    FlagOACTIVE,
			Bitmask: FlagValueOACTIVE,
		},
		{
			Name:    FlagMULTICAST,
			Bitmask: FlagValueMULTICAST,
		},
		{
			Name:    FlagLINK0,
			Bitmask: FlagValueLINK0,
		},
		{
			Name:    FlagLINK1,
			Bitmask: FlagValueLINK1,
		},
		{
			Name:    FlagLINK2,
			Bitmask: FlagValueLINK2,
		},
	}
)

// TODO: flag is actually a bitmask, but I can't figure out how to read that and compare
func ParseFlags(flag int) []string {
	var flagList []string
	for _, f := range flags {
		if f.Bitmask&flag == f.Bitmask {
			flagList = append(flagList, f.Name)
		}
	}
	return flagList
}
