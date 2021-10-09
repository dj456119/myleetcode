package myleetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 0 {
		return -1
	}
	if n == 1 {
		if isBadVersion(1) {
			return 1
		}
	}
	start := 1
	end := n
	for {
		middle := (start + end) / 2
		if isBadVersion(middle) {
			end = middle
		} else {
			start = middle
		}
		if start+1 == end {
			if isBadVersion(start) {
				return start
			} else if isBadVersion(end) {
				return end
			}
		}
		if start == end && isBadVersion(start) {
			return start
		}
	}
}

func isBadVersion(version int) bool {
	return version == 1
}

func FirstBadVersion(n int) int {
	return firstBadVersion(n)
}
