// Code generated by "stringer -type Status"; DO NOT EDIT.

package todo

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Yet-1]
	_ = x[Done-2]
}

const _Status_name = "YetDone"

var _Status_index = [...]uint8{0, 3, 7}

func (i Status) String() string {
	i -= 1
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
