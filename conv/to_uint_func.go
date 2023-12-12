package conv

// ------------------------------------------------ uint ----------------------------------------------------------------

func ToUintE(value any) (uint, error) {
	return ToUnsignedE[uint](value)
}

func ToUint(value any) uint {
	v, _ := ToUintE(value)
	return v
}

func ToUintOrDefault(value any, defaultValue uint) uint {
	v, err := ToUintE(value)
	if err != nil {
		return defaultValue
	}
	return v
}

// ------------------------------------------------ uint16 ----------------------------------------------------------------

func ToUint16E(value any) (uint16, error) {
	return ToUnsignedE[uint16](value)
}

func ToUint16(value any) uint16 {
	v, _ := ToUint16E(value)
	return v
}

func ToUint16OrDefault(value any, defaultValue uint16) uint16 {
	v, err := ToUint16E(value)
	if err != nil {
		return defaultValue
	}
	return v
}

// ------------------------------------------------ uint32 ----------------------------------------------------------------

func ToUint32E(value any) (uint32, error) {
	return ToUnsignedE[uint32](value)
}

func ToUint32(value any) uint32 {
	v, _ := ToUint32E(value)
	return v
}

func ToUint32OrDefault(value any, defaultValue uint32) uint32 {
	v, err := ToUint32E(value)
	if err != nil {
		return defaultValue
	}
	return v
}

// ------------------------------------------------ uint64 ----------------------------------------------------------------

func ToUint64E(value any) (uint64, error) {
	return ToUnsignedE[uint64](value)
}

func ToUint64(value any) uint64 {
	v, _ := ToUint64E(value)
	return v
}

func ToUint64OrDefault(value any, defaultValue uint64) uint64 {
	v, err := ToUint64E(value)
	if err != nil {
		return defaultValue
	}
	return v
}
