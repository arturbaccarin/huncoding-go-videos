package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.GetPassword()))
	ud.SetPassword(hex.EncodeToString(hash.Sum(nil)))
}
