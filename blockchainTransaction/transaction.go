package blockchainTransaction

import (
	"github.com/ltcsuite/ltcd/repository"
	"github.com/ltcsuite/ltcd/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlockchainTransactionInput struct {
	TxIn wire.TxIn
	repository.DatabaseObject
	TransactionId string
	WitnessHash   string
	BlockHash     string
	Coin          string
}

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
