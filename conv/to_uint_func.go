package conv

// ------------------------------------------------ uint ----------------------------------------------------------------

func ToUintE(value any, options ...Options) (uint, error) {
	return ToUnsignedE[uint](value, options...)
}

func ToUint(value any, options ...Options) uint {
	v, _ := ToUintE(value, options...)
	return v
}

func ToUintOrDefault(value any, defaultValue uint, options ...Options) uint {
	v, err := ToUintE(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}

// ------------------------------------------------ uint16 ----------------------------------------------------------------

func ToUint16E(value any, options ...Options) (uint16, error) {
	return ToUnsignedE[uint16](value, options...)
}

func ToUint16(value any, options ...Options) uint16 {
	v, _ := ToUint16E(value, options...)
	return v
}

func ToUint16OrDefault(value any, defaultValue uint16, options ...Options) uint16 {
	v, err := ToUint16E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}

// ------------------------------------------------ uint32 ----------------------------------------------------------------

func ToUint32E(value any, options ...Options) (uint32, error) {
	return ToUnsignedE[uint32](value, options...)
}

func ToUint32(value any, options ...Options) uint32 {
	v, _ := ToUint32E(value, options...)
	return v
}

func ToUint32OrDefault(value any, defaultValue uint32, options ...Options) uint32 {
	v, err := ToUint32E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}

// ------------------------------------------------ uint64 ----------------------------------------------------------------

func ToUint64E(value any, options ...Options) (uint64, error) {
	return ToUnsignedE[uint64](value, options...)
}

func ToUint64(value any, options ...Options) uint64 {
	v, _ := ToUint64E(value, options...)
	return v
}

func ToUint64OrDefault(value any, defaultValue uint64, options ...Options) uint64 {
	v, err := ToUint64E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}
