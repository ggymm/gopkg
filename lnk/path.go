package lnk

import (
	"bufio"
	"encoding/binary"
	"errors"
	"strings"
)

var validCLSID = [16]byte{
	0x01, 0x14, 0x02, 0x00,
	0x00, 0x00,
	0x00, 0x00,
	0xc0, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x46,
}

var (
	// ErrNotALink is returned when the header size is not 76.
	ErrNotALink = errors.New("not a link")

	// ErrInvalidCLSID is returned when the CLSID is not valid
	ErrInvalidCLSID = errors.New("invalid CLSID")
)

// ParsePath 从LNK文件中解析本地路径
func ParsePath(file *bufio.Reader) (string, error) {
	// 验证是否为LNK文件
	var size uint32
	err := binary.Read(file, binary.LittleEndian, &size)
	if err != nil {
		return "", err
	}
	if size != 76 {
		return "", ErrNotALink
	}

	var clsid [16]byte
	_, err = file.Read(clsid[:])
	if err != nil {
		return "", err
	}
	if clsid != validCLSID {
		return "", ErrInvalidCLSID
	}

	// 读取LinkFlags
	var flags uint32
	err = binary.Read(file, binary.LittleEndian, &flags)
	if err != nil {
		return "", err
	}

	hasIdList := flags&(1<<0) != 0
	hasLinkInfo := flags&(1<<1) != 0
	forceNoLinkInfo := flags&(1<<8) != 0

	// 跳过FileAttributes到Reserved3
	_, err = file.Discard(52)
	if err != nil {
		return "", err
	}

	// 如果存在IDList，需要先跳过它
	if hasIdList {
		var idListSize uint16
		err = binary.Read(file, binary.LittleEndian, &idListSize)
		if err != nil {
			return "", err
		}
		_, err = file.Discard(int(idListSize))
		if err != nil {
			return "", err
		}
	}

	if !hasLinkInfo || forceNoLinkInfo {
		return "", nil
	}

	// 跳过LinkInfoSize
	_, err = file.Discard(4)
	if err != nil {
		return "", err
	}

	// 读取LinkInfoHeaderSize
	var infoHeaderSize uint32
	err = binary.Read(file, binary.LittleEndian, &infoHeaderSize)
	if err != nil {
		return "", err
	}

	// 读取LinkInfoFlags
	var infoFlags uint32
	err = binary.Read(file, binary.LittleEndian, &infoFlags)
	if err != nil {
		return "", err
	}

	hasLocalBasePath := infoFlags&(1<<0) != 0
	if !hasLocalBasePath {
		return "", nil
	}

	// 跳过offset字段
	_, err = file.Discard(16)
	if err != nil {
		return "", err
	}

	if infoHeaderSize > 28 {
		// 跳过LocalBasePathOffsetUnicode
		_, err = file.Discard(4)
		if err != nil {
			return "", err
		}
	}

	if infoHeaderSize > 32 {
		// 跳过CommonPathSuffixOffsetUnicode
		_, err = file.Discard(4)
		if err != nil {
			return "", err
		}
	}

	// 跳过VolumeID相关字段
	_, err = file.Discard(16)
	if err != nil {
		return "", err
	}

	// 读取并跳过VolumeLabel
	_, err = file.ReadString('\x00')
	if err != nil {
		return "", err
	}

	// 读取LocalBasePath
	localPath, err := file.ReadString('\x00')
	if err != nil {
		return "", err
	}
	return strings.Trim(localPath, "\x00"), nil
}
