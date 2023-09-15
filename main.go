package main

import (
    "encoding/binary"
    "fmt"
    "math"
)

func main() {
    msg := "hi🥰"
    fmt.Println(conceal(msg))
    fmt.Println(reveal(conceal(msg)))
}

func conceal(message string) float64 {
    messageToBytes := []byte(message)
    var nan uint64 = 0x7FF0000000000000
    newNan := nan | uint64(len(messageToBytes)<<48)

    for i, toByte := range messageToBytes {
        newNan |= (uint64(toByte)) << (40 - i*8)
    }
    return math.Float64frombits(newNan)
}

func reveal(nan float64) string {
    var floatToBits = math.Float64bits(nan)
    byteArray := make([]byte, 8)
    binary.BigEndian.PutUint64(byteArray, floatToBits)
    return string(byteArray[2:])
}
