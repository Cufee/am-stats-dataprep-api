package logic

func SessionOfOverNegOne(values Values) bool {
	if val, ok := values[SessionOf].(float64); ok && val >= 0 {
		return true
	}
	return false
}

func SessionOfOverZero(values Values) bool {
	if val, ok := values[SessionOf].(float64); ok && val > 0 {
		return true
	}
	return false
}

func SessionValueOverNegOne(values Values) bool {
	if val, ok := values[SessionValue].(float64); ok && val >= 0 {
		return true
	}
	return false
}

func SessionValueOverZero(values Values) bool {
	if val, ok := values[SessionValue].(float64); ok && val > 0 {
		return true
	}
	return false
}

func AllTimeOfOverNegOne(values Values) bool {
	if val, ok := values[AllTimeOf].(float64); ok && val >= 0 {
		return true
	}
	return false
}
func AllTimeOfOverZero(values Values) bool {
	if val, ok := values[AllTimeOf].(float64); ok && val > 0 {
		return true
	}
	return false
}

func AllTimeValueOverNegOne(values Values) bool {
	if val, ok := values[AllTimeValue].(float64); ok && val >= 0 {
		return true
	}
	return false
}

func AllTimeValueOverZero(values Values) bool {
	if val, ok := values[AllTimeValue].(float64); ok && val > 0 {
		return true
	}
	return false
}

func SessionAndAllTimeOfOverZero(values Values) bool {
	if val, ok := values[SessionOf].(float64); !ok || val <= 1 {
		return false
	}
	if val, ok := values[AllTimeOf].(float64); !ok || val <= 0 {
		return false
	}
	return true
}

func SessionAndAllTimeValueOverZero(values Values) bool {
	if val, ok := values[SessionValue].(float64); !ok || val <= 1 {
		return false
	}
	if val, ok := values[AllTimeValue].(float64); !ok || val <= 0 {
		return false
	}
	return true
}
