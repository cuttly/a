package main

// Write a function that takes a byte, swaps its halfs (like the example) and returns the result.

// Example: 1byte

// 0100 | 0001
//     \ /
//     / \
// 0001 | 0100

// SOLUTION 1:
func SwapBits(octet byte) byte {
	octet = octet << 4
	octet = octet >> 4
	return octet
}

// SOLUTION 2:
// func swapBits(n byte) byte {
// 	return (n >> 4) | (n << 4)
// }

// USELESS BELOW

// FROM TESTING FILES:
// func challengeBytes(fn1, fn2 interface{}, args ...interface{}) {
// 	st1 := challenge.Monitor(fn1, args)
// 	st2 := challenge.Monitor(fn2, args)
// 	if !reflect.DeepEqual(st1.Results, st2.Results) {
// 		challenge.Fatalf("%s(%08b) == %08b instead of %08b\n",
// 			"SwapBits",
// 			args[0].(byte),
// 			st1.Results[0].(byte),
// 			st2.Results[0].(byte),
// 		)
// 	}
// }

// func main() {
// 	args := []byte{0x24, 0x14, 0x11, 0x22, 0xd2, 0x15, 0xff, 0x0, 0x35, 0x58, 0x43}

// 	for i := 0; i < 10; i++ {
// 		n := random.IntBetween(0, 255)
// 		args = append(args, byte(n))
// 	}

// 	for _, v := range args {
// 		challengeBytes(student.SwapBits, swapBits, v)
// 	}
// }
