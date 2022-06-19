# golang_course_schedule_v2

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you **must** take course `bi` first if you want to take course `ai`.

- For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return *the ordering of courses you should take to finish all courses*. If there are many valid answers, return **any** of them. If it is impossible to finish all courses, return **an empty array**.

## Examples

**Example 1:**

```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

```

**Example 2:**

```
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

```

**Example 3:**

```
Input: numCourses = 1, prerequisites = []
Output: [0]

```

**Constraints:**

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= numCourses * (numCourses - 1)`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- `ai != bi`
- All the pairs `[ai, bi]` are **distinct**.

## 解析

類似[**207. Course Schedule**](https://www.notion.so/207-Course-Schedule-5338524d2e9b4521af68993b939d04e5) 

給定一個正整數 numCourses, 代表 0 到 numCourses - 1 課程

還有一個矩陣  prerequisites 每個 entry [a_i, b_i] 代表 $b_i$ 課程必須要先完成才能完成 $a_i$

題目要求寫出一個演算法找出在給定的 numCourses, 還有 prerequisites 條件下

可能的完成 order list 

從 prerequisites 可以建構出 每個 courses 對應的 dependency map

然後從這個 dependency map

做 DFS 找出可能的 dependency list

如果遇到 dependency cycle 則回傳 empty list

如果循序找到符合的 list 則逐步加入 list 並回傳該 list

如下圖：

![](https://i.imgur.com/jfFpyh1.png)

## 程式碼
```go
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

```
## 困難點

1. 需要理解如何透過 DFS 累計可能的 Path
2. 需要檢查 dependency cycle
3. 透過 HashSet 去避免重複走訪

## Solve Point

- [x]  需要理解如何透過 DFS 累計可能的 Path
- [x]  需要檢查 dependency cycle
- [x]  透過 HashSet 去避免重複走訪