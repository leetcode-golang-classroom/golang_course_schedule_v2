package sol

type Courses []int

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	cycle := make(map[int]struct{})
	visit := make(map[int]struct{})
	preCourseMap := make(map[int]Courses, numCourses)
	for _, dependency := range prerequisites {
		preCourseMap[dependency[0]] = append(preCourseMap[dependency[0]], dependency[1])
	}
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := cycle[course]; ok {
			return false
		}
		if _, ok := visit[course]; ok {
			return true
		}
		// add to cycle
		cycle[course] = struct{}{}
		for _, preCourse := range preCourseMap[course] {
			if !dfs(preCourse) {
				return false
			}
		}
		// if not cycle found
		delete(cycle, course)
		visit[course] = struct{}{}
		result = append(result, course)
		return true
	}
	for idx := 0; idx < numCourses; idx++ {
		if !dfs(idx) {
			return []int{}
		}
	}

	return result
}
