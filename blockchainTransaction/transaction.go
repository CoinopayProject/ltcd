package blockchainTransaction

import (
	"github.com/ltcsuite/ltcd/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockchainTransaction struct {
	Id primitive.ObjectID `bson:"_id"`
	repository.DatabaseObject
	Amount                 primitive.Decimal128
	TransactionId          string
	WitnessHash            string
	ScriptClass            string
	Addresses              []string
	BlockHash              string
	RequiredSignatureCount int
	Coin                   string
}
