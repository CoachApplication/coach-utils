package utils

/**
 * uniqueStringSlice is a slice wrapper that prevents duplicate values in the slice
 */
type UniqueStringSlice struct {
	s []string
}

func (uss *UniqueStringSlice) Add(val string) {
	for _, has := range uss.s {
		if has == val {
			return
		}
	}
	uss.s = append(uss.s, val)
}
func (uss *UniqueStringSlice) Merge(vals []string) {
	for _, val := range vals {
		uss.Add(val)
	}
}
func (uss *UniqueStringSlice) Slice() []string {
	return uss.s
}
