package tree

// type Record struct {
// 	ID, Parent int
// }
//
// type ByParent []Record
//
// type Node struct {
// 	ID       int
// 	Children []*Node
// }
//
// type Mismatch struct{}
//
// func (m Mismatch) Error() string {
// 	return "c"
// }
//
// func (r ByParent) Len() int      { return len(r) }
// func (r ByParent) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
// func (r ByParent) Less(i, j int) bool {
//
// 	if r[i].ID == 0 {
// 		return true
// 	}
//
// 	if r[j].ID == 0 {
// 		return false
// 	}
//
// 	return r[i].Parent < r[j].Parent
// }

// func Build(records []Record) (*Node, error) {
//   if len(records) == 0 {
//     return nil, nil
//   }
//   root := &Node{}
//   todo := []*Node{root}
//
//   for _, record := range records {
//
//   }
// }
// func main() {
// 	records := []Record{
// 		{ID: 0},
// 		{2, 0},
// 		{3, 2},
// 		{1, 0},
// 	}
//
// 	fmt.Println(records)
//
// 	sort.Sort(ByParent(records))
// 	fmt.Println("Sorted:")
// 	fmt.Println(records)
// }
