package pkg

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"
)

type Sha1Stream struct {
	_sha1 hash.Hash
}

func (ss *Sha1Stream) Update(data []byte) {
	if ss._sha1 == nil {
		ss._sha1 = sha1.New()
	}
	ss._sha1.Write(data)
}

func (ss *Sha1Stream) Sum() string {
	return hex.EncodeToString(ss._sha1.Sum([]byte("")))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

func FileSha1(file *os.File) (string, error) {
	_sha1 := sha1.New()
	_, err := io.Copy(_sha1, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(_sha1.Sum(nil)), nil
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file *os.File) (string, error) {
	_md5 := md5.New()
	_, err := io.Copy(_md5, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(_md5.Sum(nil)), nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSize(filename string) (int64, error) {
	var result int64
	err := filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}
