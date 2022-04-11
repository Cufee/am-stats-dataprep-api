package shared

type LayoutKind string

const (
	WN8                LayoutKind = "wn8"
	Battles            LayoutKind = "battles"
	Winrate            LayoutKind = "winrate"
	Accuracy           LayoutKind = "accuracy"
	AvgDamage          LayoutKind = "avgDamage"
	WinrateWithBattles LayoutKind = "winrateWithBattles"
)
