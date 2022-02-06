/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-02-06 18:37:41
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-06 20:27:19
 */
package myleetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	nodeList := make(map[int]int)
	nodeMap := make(map[int][]int)
	for i := range prerequisites {
		for j := range prerequisites[i] {
			nodeList[prerequisites[i][j]] = 0
		}
		if z, ok := nodeMap[prerequisites[i][0]]; ok {
			z = append(z, prerequisites[i][1])
			nodeMap[prerequisites[i][0]] = z
		}
	}
	//0 未扫描 1扫描中 2扫描完成
	for k := range nodeList {
		if dfs(k, nodeMap, nodeList) {
			return true
		}
	}
	return false
}

func dfs(pos int, nodeMap map[int][]int, nodeList map[int]int) bool {
	if nodeList[pos] == 0 {
		if z, ok := nodeMap[pos]; ok {
			for i := range z {
				nodeList[pos] = 1
				if dfs(z[i], nodeMap, nodeList) {
					return true
				}
			}
		}
		nodeList[pos] = 2
		return false
	} else if nodeList[pos] == 1 {
		return true
	} else {
		return false
	}
}
