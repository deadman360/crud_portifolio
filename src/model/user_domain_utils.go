package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

var (
	journey = zap.String("journey", "UserDomain")
)

func (ud *userDomain) GetBSONValue() (interface{}, error) {
	b, err := bson.Marshal(ud)
	if err != nil {
		logger.Error("Error parsing userDomain to BSON", err, journey)
		return nil, err
	}
	var bsonValue interface{}
	if err := bson.Unmarshal(b, &bsonValue); err != nil {
		logger.Error("Error trying to unmarshal bson document", err, journey)
		return nil, err
	}
	return bsonValue, nil
}
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
