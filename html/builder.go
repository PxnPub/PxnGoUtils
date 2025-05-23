package html;

import(
	Fmt "fmt"
);



//out.Header().Set("Content-Type", "text/html");
//out.Header().Set("Content-Type", "application/json");



type Builder struct {
	IsDev  bool
	Title  string
	CSS    []string
	RawCSS string
	JS_Top []string
	JS_Bot []string
}



func NewBuilder() *Builder {
	return &Builder{};
}



func (build *Builder) Render(contents string) string {
	return build.RenderTop() + "\n" +
		contents +
		"\n" + build.RenderBottom();
}

func (build *Builder) RenderTop() string {
	out := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width,initial-scale=1.0" />
<meta http-equiv="Cache-Control" content="max-age=3600, must-revalidate" />
`;
	if build.Title != "" {
		out += Fmt.Sprintf("<title>%s</title>", build.Title);
	}
	// css
	if len(build.CSS) != 0 {
		for _, file := range build.CSS {
			out += Fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, file) + "\n";
		}
	}
	if build.RawCSS != "" {
		out += "\n<style type=\"text/css\">\n" + build.RawCSS + "\n</style>\n";
	}
	// js
	if len(build.JS_Top) != 0 {
		for _, file := range build.JS_Top {
			out += Fmt.Sprintf(`<script src="%s" defer></script>`, file) + "\n";
		}
	}
	return out + "</head>\n<body>\n\n\n";
}

func (build *Builder) RenderBottom() string {
	out := "";
//TODO: bottom js
	return out + "\n\n</body>\n</html>\n";
}



// title
func (build *Builder) SetTitle(title string) *Builder {
	build.Title = title;
	return build;
}



// css
func (build *Builder) AddCSS(path string) *Builder {
	build.CSS = append(build.CSS, path);
	return build;
}

func (build *Builder) AddRawCSS(css string) *Builder {
	build.RawCSS += "\n" + css + "\n";
	return build;
}



// js
func (build *Builder) AddTopJS(path string) *Builder {
	build.JS_Top = append(build.JS_Top, path);
	return build;
}

func (build *Builder) AddBotJS(path string) *Builder {
	build.JS_Bot = append(build.JS_Bot, path);
	return build;
}
