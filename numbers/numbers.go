package numbers;

import(
	Fmt     "fmt"
	Strings "strings"
	StrConv "strconv"
);



func ToBase36(val uint64) string {
	result := make([]byte, 13);
	copy(result, "0000000000000");
	str := Strings.ToUpper(StrConv.FormatUint(uint64(val), 36));
	size := len(str);
	copy(result[13-size:], str);
	return string(result);
}

func FromBase36(str string) (uint64, error) {
	if len(str) != 13 { return 0, Fmt.Errorf("Invalid UID value: %s", str); }
	return StrConv.ParseUint(str, 36, 64);
}



func FormatByteSize(size int64) string {
	if size > 1000000000000 { return Fmt.Sprintf("%dT", size / 1000000000000); }
	if size >    1000000000 { return Fmt.Sprintf("%dG", size /    1000000000); }
	if size >       1000000 { return Fmt.Sprintf("%dM", size /       1000000); }
	if size >          1000 { return Fmt.Sprintf("%dK", size /          1000); }
	return Fmt.Sprintf("%d", size);
}
