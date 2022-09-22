package main

// Might or might not work
func main() {
	reversebits(b00100110)
}

// reversebits(00100110) -> 01100100
func reversebits(octet byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		if octet&(1<<i) != 0 {
			result |= 1 << (7 - i)
		}
	}
	return result
}

// from the internet
// func reverseBits(num uint32) uint32 {
//     var ret = uint32(0)
//     var power = uint32(31)
//     for num != 0 {
//         ret += (num & 1) << power
//         num = num >> 1
//         power -= 1
//     }
//     return ret
// }

// from leetcode

// func reverseBits(num uint32) uint32 {
// 	var reversed uint32

// 	// use a counter instead of making the input smaller

// 	var index uint32
// 	for index < 32 {
// 		// fetch the current bit
// 		lastBit := getBitAtIndex(num, index)

// 		// push it in the reversed number (push left)

// 		// -- make space for the new number
// 		reversed = reversed << 1
// 		// -- copy the last bit
// 		reversed = reversed | lastBit

// 		index++
// 	}

// 	return reversed
// }

// func getBitAtIndex(num uint32, index uint32) uint32 {
// 	tmp := num >> index
// 	return tmp & 1
// }
