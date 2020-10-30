package main

import (
	"encoding/gob"

	"github.com/Guiguerreiro39/go-auction-house/input"
)

func bid(enc *gob.Encoder) {
	bid := input.Bid()
	enc.Encode(bid)
}

func joinRoom(enc *gob.Encoder, dec *gob.Decoder) {
	var success bool
	id := input.JoinRoom()

	enc.Encode(id)
	dec.Decode(&success)
	receiver()

	if success {
		insideRoom = true
	}
}
