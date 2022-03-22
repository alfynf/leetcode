package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(n int) bool {

}

func firstBadVersion(n int) int {
	var min int = 2147483647
	for i := 1; i <= n; i++ {
		if isBadVersion(i) == true && i <= min {
			min = i
			break
		}
	}
	return min
}

func main() {
	fmt.Println(firstBadVersion(5))
}
