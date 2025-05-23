package html;



type PairURL struct {
	Pub string
	Dev string
}

var (
	// bootstrap
	URL_BootstrapJS = PairURL {
		Pub: "https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.3.3/js/bootstrap.min.js",
		Dev: "https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.3.3/js/bootstrap.js",
	}
	URL_BootstrapCSS = PairURL {
		Pub: "https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.3.3/css/bootstrap.min.css",
		Dev: "https://cdnjs.cloudflare.com/ajax/libs/bootstrap/5.3.3/css/bootstrap.css",
	}
	// font-awesome
	URL_FontAwesomeJS = PairURL {
		Pub: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/js/all.min.js",
		Dev: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/js/all.js",
	}
	URL_FontAwesomeCSS = PairURL {
		Pub: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/css/all.min.css",
		Dev: "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/css/all.css",
	}
	// echarts
	URL_EChartsJS = PairURL {
		Pub: "https://cdnjs.cloudflare.com/ajax/libs/echarts/5.6.0/echarts.min.js",
		Dev: "https://cdnjs.cloudflare.com/ajax/libs/echarts/5.6.0/echarts.js",
	}
);



// bootstrap
func (build *Builder) AddBootstrap() *Builder {
	build.AddTopJS(PubDevURL(build.IsDev, URL_BootstrapJS ));
	build.AddCSS(  PubDevURL(build.IsDev, URL_BootstrapCSS));
	return build;
}

// font-awesome
func (build *Builder) AddFontAwesome() *Builder {
	build.AddTopJS(PubDevURL(build.IsDev, URL_FontAwesomeJS ));
	build.AddCSS(  PubDevURL(build.IsDev, URL_FontAwesomeCSS));
	return build;
}

// echarts
func (build *Builder) AddECharts() *Builder {
	build.AddTopJS(PubDevURL(build.IsDev, URL_EChartsJS));
	return build;
}



func PubDevURL(isdev bool, url PairURL) string {
	if isdev { return url.Dev;
	} else {   return url.Pub; }
}
