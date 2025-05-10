package pxnGoUtils;

import(
	Fmt      "fmt"
	Strings  "strings"
	FilePath "path/filepath"
	OS       "os"
	OSUser   "os/user"
	Errors   "errors"
);



var DefaultConfigSearchPaths = []string{
	"./",
	"/",
	"/etc",
	"~/",
};



func IsFile(file string) bool {
	info, err := OS.Stat(file);
	if err != nil {
		if Errors.Is(err, OS.ErrNotExist) { return false; }
		panic(err);
	}
	return info.Mode().IsRegular();
}

func FindFile(file string, paths...string) string {
	for i := range paths {
		p := FilePath.Join(paths[i], file);
		if IsFile(p) { return p; }
	}
	return "";
}

func ExpandPath(path string) string {
	user, err := OSUser.Current();
	if err != nil { panic(err); }
	if path == "~" { return user.HomeDir; }
	if Strings.HasPrefix(path, "~/") {
		return FilePath.Join(user.HomeDir, path[2:]);
	}
	abs, err := FilePath.Abs(path);
	if err != nil { panic(err); }
	return abs;
}



func FormatByteSize(size int64) string {
	if size > 1000000000000 { return Fmt.Sprintf("%dT", size / 1000000000000); }
	if size >    1000000000 { return Fmt.Sprintf("%dG", size /    1000000000); }
	if size >       1000000 { return Fmt.Sprintf("%dM", size /       1000000); }
	if size >          1000 { return Fmt.Sprintf("%dK", size /          1000); }
	return Fmt.Sprintf("%d", size);
}
