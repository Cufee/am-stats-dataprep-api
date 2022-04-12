package shared

type DataType string

const (
	WN8                DataType = "wn8"
	Battles            DataType = "battles"
	Winrate            DataType = "winrate"
	Accuracy           DataType = "accuracy"
	AvgDamage          DataType = "avgDamage"
	WinrateWithBattles DataType = "winrateWithBattles"
)
