package compact

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
	"sync"
)

type TileQuery struct {
	Z, Y, X int
}

func (tq *TileQuery) bundleFilePath(size int) string {
	var buf [22]byte

	baseRow := (tq.Y / size) * size
	baseCol := (tq.X / size) * size

	const hex = "0123456789abcdef"

	buf[0] = '/'
	buf[1] = 'L'
	buf[2] = '0' + byte(tq.Z/10)
	buf[3] = '0' + byte(tq.Z%10)

	buf[4] = '/'
	buf[5] = 'R'
	buf[6] = hex[(baseRow>>12)&0xF]
	buf[7] = hex[(baseRow>>8)&0xF]
	buf[8] = hex[(baseRow>>4)&0xF]
	buf[9] = hex[baseRow&0xF]

	buf[10] = 'C'
	buf[11] = hex[(baseCol>>12)&0xF]
	buf[12] = hex[(baseCol>>8)&0xF]
	buf[13] = hex[(baseCol>>4)&0xF]
	buf[14] = hex[baseCol&0xF]

	copy(buf[15:], ".bundle")
	return string(buf[:])
}

func (tq *TileQuery) getTileData(file *os.File, tileOffset uint32, tileSize uint32) ([]byte, error) {
	buffer := make([]byte, tileSize)

	_, err := file.ReadAt(buffer, int64(tileOffset))
	if err != nil && err != io.EOF {
		return nil, err
	}

	return buffer, nil
}

func (tq *TileQuery) getTileRecord(file *os.File, size int) (uint32, uint32, error) {
	tileIndexOffset := int64(64 + 8*(size*(tq.Y%size)+(tq.X%size)))

	var buffer [8]byte

	_, err := file.ReadAt(buffer[:], tileIndexOffset)
	if err != nil && err != io.EOF {
		return 0, 0, err
	}

	tileOffset := binary.LittleEndian.Uint32(buffer[0:4])
	tileSize := binary.LittleEndian.Uint32(buffer[4:8])

	return tileOffset, tileSize, nil
}

var (
	bundleCache   = make(map[string]*os.File)
	bundleCacheMu sync.RWMutex
)

func openBundleFile(path string) (*os.File, error) {
	bundleCacheMu.RLock()

	cachedFile, ok := bundleCache[path]
	bundleCacheMu.RUnlock()
	if ok {
		return cachedFile, nil
	}

	bundleCacheMu.Lock()
	defer bundleCacheMu.Unlock()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	bundleCache[path] = file

	return file, nil
}

func (tq *TileQuery) GetCompactCache(size int, tileDir string) ([]byte, error) {
	bundleFile := tq.bundleFilePath(size)

	file, err := openBundleFile(tileDir + bundleFile)
	if err != nil {
		return nil, err
	}

	tileOffset, tileSize, err := tq.getTileRecord(file, size)
	if err != nil {
		return nil, err
	}

	if tileSize == 0 {
		return nil, errors.New("no tilecontent in bundle")
	}

	tileData, err := tq.getTileData(file, tileOffset, tileSize)
	if err != nil {
		return nil, err
	}

	return tileData, nil

}
