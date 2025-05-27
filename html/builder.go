package html;

import(
	Fmt  "fmt"
	HTTP "net/http"
);



const(
	MIME_TEXT = "text/plain"
	MIME_HTML = "text/html"
	MIME_JSON = "application/json"
);



type Builder struct {
	IsDev         bool
	Title         string
	AppendHead    string
	AppendHeader  string
	AppendFooter  string
	IsBootstrap   bool
	IsJQuery      bool
	IsDataTables  bool
	IsFontAwesome bool
	IsECharts     bool
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
	if build.Title != ""   { out += Fmt.Sprintf("<title>%s</title>\n", build.Title); }
	if build.IsBootstrap   { out += Fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, URL_BootstrapCSS          ) + "\n"; }
	if build.IsDataTables  { out += Fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, URL_DataTablesBootstrapCSS) + "\n"; }
	if build.IsDataTables  { out += Fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, URL_DataTablesScrollerCSS ) + "\n"; }
	if build.IsFontAwesome { out += Fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, URL_FontAwesomeCSS        ) + "\n"; }
	if build.IsJQuery      { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_JQueryJS              ) + "\n"; }
	if build.IsBootstrap   { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_BootstrapJS           ) + "\n"; }
	if build.IsDataTables  { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_DataTablesJS          ) + "\n"; }
	if build.IsDataTables  { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_DataTablesBootstrapJS ) + "\n"; }
	if build.IsDataTables  { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_DataTablesScrollerJS  ) + "\n"; }
	if build.IsDataTables  { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_DataTablesPageResizeJS) + "\n"; }
	if build.IsFontAwesome { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_FontAwesomeJS         ) + "\n"; }
	if build.IsECharts     { out += Fmt.Sprintf(`<script src="%s"></script>`,          URL_EChartsJS             ) + "\n"; }
	if build.AppendHead != "" { out += build.AppendHead; }
	out += "</head>\n<body>\n\n\n";
	if build.AppendHeader != "" {
		out += build.AppendHeader + "\n\n\n";
	}
	return out;
}

func (build *Builder) RenderBottom() string {
	out := "";
	if build.AppendFooter != "" {
		out += "\n\n\n" + build.AppendFooter + "\n";
	}
	return out + "\n\n</body>\n</html>\n";
}



func SetContentType(out HTTP.ResponseWriter, mime string) {
	switch mime {
		case "text": mime = MIME_TEXT;
		case "html": mime = MIME_HTML;
		case "json": mime = MIME_JSON;
	}
	out.Header().Set("Content-Type", mime);
}



// title
func (build *Builder) SetTitle(title string) *Builder {
	build.Title = title;
	return build;
}



// css
func (build *Builder) AddCSS(path string) *Builder {
	build.AppendHeader += Fmt.Sprintf(`<link rel="stylesheet" href="%s" />`, PubDevURL(build.IsDev, path)) + "\n";
	return build;
}

func (build *Builder) AddRawCSS(css string) *Builder {
	build.AppendHeader += `<style type="text/css">` + "\n" + css + "\n</style>\n";
	return build;
}



// js
func (build *Builder) AddTopJS(path string) *Builder {
	build.AppendHeader += Fmt.Sprintf(`<script src="%s"></script>`, PubDevURL(build.IsDev, path)) + "\n";
	return build;
}

func (build *Builder) AddBotJS(path string) *Builder {
	build.AppendFooter += Fmt.Sprintf(`<script src="%s"></script>`, PubDevURL(build.IsDev, path)) + "\n";
	return build;
}



func (build *Builder) WithBootstrap() *Builder {
	build.IsBootstrap = true;
	return build;
}

func (build *Builder) WithJQuery() *Builder {
	build.IsJQuery = true;
	return build;
}

func (build *Builder) WithDataTables() *Builder {
	build.IsDataTables = true;
	return build;
}

func (build *Builder) WithFontAwesome() *Builder {
	build.IsFontAwesome = true;
	return build;
}

func (build *Builder) WithECharts() *Builder {
	build.IsECharts = true;
	return build;
}
