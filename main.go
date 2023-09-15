package main

import "log"

func main() {
	conceal("hello")
}

func conceal(message string) {
	var nan uint64 = 0x7FF8000000000001
	b := []byte(message)
	for i, b2 := range b {
		nan |= (uint64(b2) << uint64(i*3))
	}
	log.Printf("%x", nan)
	//binary.Read(&r, binary.BigEndian, message)

}

//0001
//0.001 2x3
//0b010000001000110000000000000000000
//00110000000000000000000 = 0.625
//1.27
// 0001
// 0.25 = 0.0100011
