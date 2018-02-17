package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type ByParent []Record

// type Mismatch struct{}

// func (m Mismatch) Error() string {
// 	return "c"
// }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(ByParent(records))

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("No root node")
	}

	// Initialize parent catalogue
	parents := make(map[int]*Node)
	currParent := 0

	// Initialize root and duplicate tracker
	root := &Node{ID: 0}
	parents[0] = root
	possibleDup := records[0]

	// Process all records except root node
	for i := 1; i < len(records); i++ {
		curr := records[i]

		// check if record is valid
		err := checkValid(curr)
		if err != nil {
			return nil, err
		}

		//check if record is duplicate
		if curr.ID == possibleDup.ID &&
			curr.Parent == possibleDup.Parent {
			return nil, errors.New("duplicate")
		}

		if _, ok := parents[curr.Parent]; !ok {
			return nil, errors.New("No parent for this node")
		}

		if currParent != curr.Parent {
			if currParent+1 == curr.Parent {
				currParent++
			} else {
				return nil, errors.New("Not continuous")
			}
		}

		possibleDup = curr
		currNode := &Node{ID: curr.ID}
		parents[curr.ID] = currNode
		parent := parents[curr.Parent]
		parent.Children = append(parent.Children, currNode)
	}

	return root, nil
}

func checkValid(r Record) error {
	if r.Parent > r.ID {
		return errors.New("b")
	}

	return nil
}

func (r ByParent) Len() int      { return len(r) }
func (r ByParent) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r ByParent) Less(i, j int) bool {
	if r[i].ID == 0 {
		return true
	}

	if r[j].ID == 0 {
		return false
	}

	if r[i].Parent == r[j].Parent {
		return r[i].ID < r[j].ID
	}

	return r[i].Parent < r[j].Parent
}
