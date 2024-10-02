package utils

import (
	"math/rand"
)

func AssignGate() string {
	gates := []string{"A1", "A2", "B1", "B2", "C1", "C2"}
	return gates[rand.Intn(len(gates))]
}
