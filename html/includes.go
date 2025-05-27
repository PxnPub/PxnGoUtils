package html;

import(
	Strings "strings"
);



var (
	// bootstrap
	URL_BootstrapCSS = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.6/dist/css/bootstrap.min.css"
	URL_BootstrapJS  = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.6/dist/js/bootstrap.bundle.min.js"

	// jquery
	URL_JQueryJS = "https://cdnjs.cloudflare.com/ajax/libs/jquery/3.7.1/jquery.min.js"

	// datatables
	URL_DataTablesJS           = "https://cdn.datatables.net/2.3.1/js/dataTables.min.js"
	URL_DataTablesBootstrapJS  = "https://cdn.datatables.net/2.3.1/js/dataTables.bootstrap5.min.js"
	URL_DataTablesBootstrapCSS = "https://cdn.datatables.net/2.3.1/css/dataTables.bootstrap5.min.css"
	URL_DataTablesScrollerJS   = "https://cdn.datatables.net/scroller/2.4.3/js/dataTables.scroller.min.js"
	URL_DataTablesScrollerCSS  = "https://cdn.datatables.net/scroller/2.4.3/css/scroller.bootstrap5.min.css"
	URL_DataTablesPageResizeJS = "https://cdn.datatables.net/plug-ins/2.3.1/features/pageResize/dataTables.pageResize.min.js"

	// font-awesome
	URL_FontAwesomeCSS = "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/css/all.min.css"
	URL_FontAwesomeJS  = "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.7.2/js/all.min.js"

	// echarts
	URL_EChartsJS = "https://cdnjs.cloudflare.com/ajax/libs/echarts/5.6.0/echarts.min.js"
);



func PubDevURL(isdev bool, url string) string {
	if isdev {
		if Strings.HasSuffix(url, ".min.css") { return Strings.TrimSuffix(url, ".min.css") + ".css"; }
		if Strings.HasSuffix(url, ".min.js" ) { return Strings.TrimSuffix(url, ".min.js" ) + ".js";  }
	}
	return url;
}
