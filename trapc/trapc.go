package trapc;

import(
	OS      "os"
	Signal  "os/signal"
	Sync    "sync"
	SysCall "syscall"
	Time    "time"
);



type TrapC struct {
	WaitGrp *Sync.WaitGroup
	StopChn []chan bool
}



func New() *TrapC {
	var waitgroup Sync.WaitGroup;
	stopchans := make([]chan bool, 0);
	sig := make(chan OS.Signal, 1);
	Signal.Notify(sig, SysCall.SIGINT, SysCall.SIGTERM);
	loops   := 0;
	timeout := 0;
	trapc := TrapC{
		WaitGrp: &waitgroup,
		StopChn: stopchans,
	};
	// timeout loop
	go func() {
		for {
			Time.Sleep(Time.Second);
			timeout++;
			if timeout >= 10 {
				timeout = 0;
				if loops > 0 { loops--; }
			}
		}
	}();
	// ctrl+c loop
	go func() {
		for {
			<-sig;
			switch loops {
			case 0:
				print("\rStopping..     \n");
				for i := range trapc.StopChn {
					trapc.StopChn[i] <- true;
				}
				break;
			case 1:
				print("\rTerminate?     \n");
				break;
			default:
				if loops > 0 {
					print("\rTerminate!     \n");
					OS.Exit(0);
				}
				break;
			}
			timeout = 0;
			loops++;
		}
	}();
	return &trapc;
}



func (trapc *TrapC) NewStopChan() chan bool {
	stopchan := make(chan bool, 1);
	trapc.StopChn = append(trapc.StopChn, stopchan);
	return stopchan;
}



func (trapc *TrapC) Wait() {
	trapc.WaitGrp.Wait();
}
