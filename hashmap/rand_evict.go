// Copyright 2012-2013 Apcera Inc. All rights reserved.

package hashmap

import (
	"math/rand"
	"time"
)

// We use init to setup the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RemoveRandom can be used for a random policy eviction.
// This is stochastic but very fast and does not impede
// performance like LRU, LFU or even ARC based implementations.
func (h *HashMap) RemoveRandom() {
	if h.used == 0 {
		return
	}
	index := (rand.Int()) & int(h.msk)
	// Walk forward til we find an entry
	for i := index; i < len(h.bkts); i++ {
		e := &h.bkts[i]
		if *e != nil {
			*e = (*e).next
			h.used -= 1
			return
		}
	}
	// If we are here we hit end and did not remove anything,
	// use the index and walk backwards.
	for i := index; i >= 0; i-- {
		e := &h.bkts[i]
		if *e != nil {
			*e = (*e).next
			h.used -= 1
			return
		}
	}
	panic("Should not reach here..")
}

// RemoveRandom can be used for a random policy eviction.
// This is stochastic but very fast and does not impede
// performance like LRU, LFU or even ARC based implementations.
func RemoveRandom(h map[string][]interface{}) {
	if len(h) == 0 {
		return
	}
	// As of Go 1, map iteration order is randomized.
	// TODO: Determine performance impact of range. Also, relying on this
	// randomization may not be reliable as it could change in the future.
	var key string
	for key, _ = range h {
		break
	}
	delete(h, key)
}
