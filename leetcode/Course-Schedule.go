
func canFinish(numCourses int, prerequisites [][]int) bool {

	var g map[int][]int = make(map[int][]int)
	for _, y := range prerequisites {
		g[y[1]] = append(g[y[1]], y[0])
	}
	var q []int
	var inDegree []int = make([]int, numCourses)
	for _, v := range g {
		for _, i := range v {
			inDegree[i]++
		}
	}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}
	var count int
	for len(q) != 0 {
		for _, v := range g[q[0]] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
		q = q[1:]
		count++
	}
	if count != numCourses {
		return false
	}
	return true

}
