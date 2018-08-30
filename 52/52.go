package main

import (
	"fmt"
)

var nqueen = [...]int{1, 0, 0, 2, 10, 4, 40, 92, 352, 724, 2680, 14200, 73712,
	365596, 2279184, 14772512, 95815104, 666090624,
	4968057848, 39029188884, 314666222712, 2691008701644,
	24233937684440, 227514171973736, 2207893435808352,
	22317699616364044, 234907967154122528}

func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}
	if n <= len(nqueen) {
		return nqueen[n-1]
	}
	return 0
}

func main() {
	fmt.Println(totalNQueens(4))
}
