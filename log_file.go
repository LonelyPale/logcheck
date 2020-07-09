package logcheck

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) (fbs []byte, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = f.Close()
	}()

	return ioutil.ReadAll(f)
}

func CountLine(bs []byte) (uint64, error) {
	fpath := "/Users/wyb/data/gov2-data-mainnet-mysql/log/netsync.20200507"
	fbs, err := ReadFile(fpath)
	if err != nil {
		return 0, err
	}

	count := uint64(0)
	bslen := len(fbs)

	if bslen == 0 {
		return 0, nil
	}

	for _, b := range fbs {
		if b == 0x0a {
			count++
		}
	}

	if fbs[bslen-1] != 0x0a {
		count++
	}

	return count, nil
}
