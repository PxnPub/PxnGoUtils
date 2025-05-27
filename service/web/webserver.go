package web;

import(
	Log      "log"
	Ctx      "context"
	Time     "time"
	HTTP     "net/http"
	Atomic   "sync/atomic"
	Gorilla  "github.com/gorilla/mux"
	TrapC    "github.com/PxnPub/pxnGoUtils/trapc"
	NetUtils "github.com/PxnPub/pxnGoUtils/net"
);



const DefaultBindWeb = "tcp://127.0.0.1:8000";



type WebServer struct {
	TrapC   *TrapC.TrapC
	Srv     *HTTP.Server
	Mux     *Gorilla.Router
	Bind    string
	StatReq Atomic.Uint64
}



func NewWebServer(trapc *TrapC.TrapC, bind string) *WebServer {
	server := WebServer{
		TrapC: trapc,
		Mux:   Gorilla.NewRouter(),
		Bind:  bind,
	};
	server.Mux.Use(middleware_stats(&server));
	return &server;
}



func (server *WebServer) Start() {
	server.Srv = &HTTP.Server{
		Addr:    server.Bind,
		Handler: server.Mux,
	};
	go func () {
		server.TrapC.WaitGroup.Add(1);
		defer server.TrapC.WaitGroup.Done();
		server.TrapC.AddStopHook(func() {
			if err := server.Srv.Shutdown(Ctx.Background()); err != nil {
				panic(err);
			}
		});
		NetUtils.RemoveOldSocket(server.Bind);
		listen, err := NetUtils.NewSock(server.Bind);
		if err != nil { panic(err); }
		Log.Printf("[%s] Listening..", server.Bind);
		if err := server.Srv.Serve(*listen); err != nil {
			switch err.Error() {
				case "http: Server closed":
					Log.Printf("[%s] Listener closed.", server.Bind);
					break;
				default: panic(err);
			}
		}
	}();
	sleep, _ := Time.ParseDuration("100ms");
	Time.Sleep(sleep);
}



func middleware_stats(server *WebServer) Gorilla.MiddlewareFunc {
	return func(next HTTP.Handler) HTTP.Handler {
		return HTTP.HandlerFunc(func(out HTTP.ResponseWriter, in *HTTP.Request) {
			cnt := server.StatReq.Add(1);
			Log.Printf("REQUEST[%d] %s\n", cnt, in.RequestURI);
			next.ServeHTTP(out, in);
		});
	}
}



func AddRouteStatic(mux *Gorilla.Router) {
	fs := HTTP.FileServer(HTTP.Dir("./static"));
	mux.PathPrefix("/static/").Handler(HTTP.StripPrefix("/static/", fs));
}
