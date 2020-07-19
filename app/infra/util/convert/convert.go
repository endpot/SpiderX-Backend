package convert

func ParseBoolToUint8(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}

func ParseUint8ToBool(u uint8) bool {
	if u > 0 {
		return true
	}

	return false
}
