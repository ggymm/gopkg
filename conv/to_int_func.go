package conv

// ------------------------------------------------ int ----------------------------------------------------------------

func ToIntE(value any) (int, error) {
	return toSignedE[int](value)
}

func ToInt(value any) int {
	v, _ := ToIntE(value)
	return v
}

func ToIntOrDefault(value any, defaultValue int) int {
	v, err := ToIntE(value)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ int16 ----------------------------------------------------------------

func ToInt16E(value any) (int16, error) {
	return toSignedE[int16](value)
}

func ToInt16(value any) int16 {
	v, _ := ToInt16E(value)
	return v
}

func ToInt16OrDefault(value any, defaultValue int16) int16 {
	v, err := ToInt16E(value)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ int32 ----------------------------------------------------------------

func ToInt32E(value any) (int32, error) {
	return toSignedE[int32](value)
}

func ToInt32(value any) int32 {
	v, _ := ToInt32E(value)
	return v
}

func ToInt32OrDefault(value any, defaultValue int32) int32 {
	v, err := ToInt32E(value)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ int64 ----------------------------------------------------------------

func ToInt64E(value any) (int64, error) {
	return toSignedE[int64](value)
}

func ToInt64(value any) int64 {
	v, _ := ToInt64E(value)
	return v
}

func ToInt64OrDefault(value any, defaultValue int64) int64 {
	v, err := ToInt64E(value)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
