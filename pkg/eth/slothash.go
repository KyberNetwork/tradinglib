package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func OverrideSet(stateDiff map[common.Hash]common.Hash, storageSlot int, key common.Hash) {
	// for a set the first slot is used to store the array to support enumerable.
	// The slot right after it is used to store the map (need to override this map).
	OverrideMap(stateDiff, storageSlot+1, key, common.BigToHash(common.Big1))
}

func OverrideVariable(stateDiff map[common.Hash]common.Hash, storageSlot int, value common.Hash) {
	slotHash := common.BigToHash(big.NewInt(int64(storageSlot)))
	stateDiff[slotHash] = value
}

func OverrideMap(stateDiff map[common.Hash]common.Hash, storageSlot int, key, value common.Hash) {
	overrideSlot32 := common.BigToHash(big.NewInt(int64(storageSlot)))
	overrideSlotHash := crypto.Keccak256Hash(key[:], overrideSlot32[:])
	stateDiff[overrideSlotHash] = value
}
