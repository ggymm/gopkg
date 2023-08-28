package convert

// ------------------------------------------------ int ----------------------------------------------------------------

func ToIntE(value any, options ...Options) (int, error) {
	return toSignedE[int](value, options...)
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

func ToInt16E(value any, options ...Options) (int16, error) {
	return toSignedE[int16](value, options...)
}

func ToInt16(value any, options ...Options) int16 {
	v, _ := ToInt16E(value, options...)
	return v
}

func ToInt16OrDefault(value any, defaultValue int16, options ...Options) int16 {
	v, err := ToInt16E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ int32 ----------------------------------------------------------------

func ToInt32E(value any, options ...Options) (int32, error) {
	return toSignedE[int32](value, options...)
}

func ToInt32(value any, options ...Options) int32 {
	v, _ := ToInt32E(value, options...)
	return v
}

func ToInt32OrDefault(value any, defaultValue int32, options ...Options) int32 {
	v, err := ToInt32E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ int64 ----------------------------------------------------------------

func ToInt64E(value any, options ...Options) (int64, error) {
	return toSignedE[int64](value, options...)
}

func ToInt64(value any, options ...Options) int64 {
	v, _ := ToInt64E(value, options...)
	return v
}

func ToInt64OrDefault(value any, defaultValue int64, options ...Options) int64 {
	v, err := ToInt64E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
