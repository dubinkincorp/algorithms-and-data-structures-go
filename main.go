package main

import "fmt"

//column H is populated with 0
var h uint64 = 9187201950435737471
//columns G and H are populated with 0
var gh uint64 = 4557430888798830399
//columns A and B are populated with 0
var ab uint64 = 18229723555195321596
//column A is populated with 0
var a uint64 = 18374403900871474942

func kingsNextMoveVariants(position int) uint64 {

	var k uint64 = 1 << position

	//k on A column
	var kA = k & a
	//k on H column
	var kH = k & h

	return kA<<7 | k<<8 | kH<<9 | kA>>1 | kH<<1 | kA>>9 | k>>8 | kH>>7
}

func knightsNextMoveVariants(position int) uint64 {

	var k uint64 = 1 << position

	//k on H column
	var kH = k & h
	//k on G column
	var kG = k & gh
	//k on B column
	var kB = k & ab
	//k on A column
	var kA = k & a

	return kB << 6 | kB >> 10 | kA << 15 | kA >> 17 | kH<< 17 | kH>> 15 | kG<< 10 | kG>> 6
}

func populationCnt(num uint64) int {

	var cnt = 0

	for ; num > 0;  {

		if (num & 1) == 1 {
			cnt++
		}

		num >>= 1
	}

	return cnt
}

func populationCnt2(num uint64) int {

	var cnt = 0

	for ; num > 0; {

		cnt++

		num &= num - 1
	}

	return cnt
}

func main() {
	var kingsResult = kingsNextMoveVariants(28)
	fmt.Print("kingsResult = ")
	fmt.Println(kingsResult)
	fmt.Print("kings population count = ")
	fmt.Println(populationCnt(kingsResult))
	fmt.Print("kings population count = ")
	fmt.Println(populationCnt2(kingsResult))

	var knightsResult = knightsNextMoveVariants(28)
	fmt.Print("knightsResult = ")
	fmt.Println(knightsResult)
	fmt.Print("knights population count = ")
	fmt.Println(populationCnt(knightsResult))
	fmt.Print("knights population count = ")
	fmt.Println(populationCnt2(knightsResult))
}
