package model

import (
	"crypto/md5"
	"encoding/hex"

	"go.uber.org/zap"
)

var (
	journey = zap.String("journey", "UserDomain")
)

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
