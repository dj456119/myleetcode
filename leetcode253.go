/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-01 14:47:59
 * @LastEditors: cm.d
 * @LastEditTime: 2022-01-01 15:23:09
 */
package myleetcode

import (
	"fmt"
	"sort"
)

type Meeting struct {
	Start int
	End   int
	All   [][]int
}

func NewMeeting(meet []int) *Meeting {
	m := Meeting{
		Start: meet[0],
		End:   meet[1],
		All:   [][]int{meet},
	}
	return &m
}

func (m *Meeting) isOK(meet []int) bool {
	if meet[0] >= m.End || meet[1] <= m.Start {
		return true
	}
	return false
}

func (m *Meeting) Add(meet []int) {
	fmt.Println(m, meet)
	if meet[0] >= m.End {
		m.End = meet[1]
	}
	if meet[1] <= m.Start {
		m.Start = meet[0]
	}
	m.All = append(m.All, meet)
}

func minMeetingRooms(intervals [][]int) int {
	ms := []*Meeting{}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return false
	})

	for _, j := range intervals {
		if len(ms) == 0 {
			m := NewMeeting(j)
			ms = append(ms, m)

		} else {
			flag := false
			for _, m := range ms {
				b := m.isOK(j)
				if b {
					m.Add(j)
					flag = true
					break
				}
			}
			if !flag {
				m := NewMeeting(j)
				ms = append(ms, m)
			}
		}
	}
	return len(ms)
}
