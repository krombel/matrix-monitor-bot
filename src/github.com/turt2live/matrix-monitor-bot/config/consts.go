package config

import (
	"time"
)

const RemoteSendDelayThreshold = 1 * time.Second
const ReceiveDelayThreshold = 5 * time.Second
const RttWarningThreshold = 10 * time.Second
const ProcessingDelayThreshold = 10*time.Millisecond // This is pretty relaxed

// Metrics:
// [ bot ] --A-> [ matrix.org ] --B-> [ t2bot.io ] --C-> [ bot (G) ]
//                                                          |
// [ bot ] <-F-- [ matrix.org ] <-E-- [ t2bot.io ] <-D------+
// A: Ping remote send delay
// B: Ping federation delay
// C: Ping sync delay
// D: Pong send delay
// E: Pong federation delay
// F: Pong sync delay
// G: The processing delay for a ping