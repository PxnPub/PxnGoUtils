package trapc;

import(
	OS      "os"
	Signal  "os/signal"
	Sync    "sync"
	Atomic  "sync/atomic"
	SysCall "syscall"
	Time    "time"
	ThdSafe "github.com/hayageek/threadsafe"
);



type TrapState int8;
const (
	State_OK   TrapState = iota
	State_Stop
	State_Warn
	State_Term
);

type StopHook func();

type TrapC struct {
	WaitGroup *Sync.WaitGroup
	StopChans ThdSafe.Slice[chan bool]
	StopHooks ThdSafe.Slice[StopHook]
	Stopping  Atomic.Bool
	State     TrapState
	Timeout   int8
}



func New() *TrapC {
	var waitgroup Sync.WaitGroup;
	signals := make(chan OS.Signal, 1);
	Signal.Notify(signals, SysCall.SIGINT, SysCall.SIGTERM);
	trapc := TrapC{
		WaitGroup: &waitgroup,
	};
	// ctrl+c loop
	go func() {
		timer := Time.NewTicker(Time.Second);
		for {
			select {
			case <-timer.C:
				trapc.Timeout++;
				if trapc.Timeout >= 10 {
					trapc.Timeout = 0;
					if trapc.State > State_OK {
						trapc.State--;
					}
				}
			case <-signals:
				trapc.State++;
				trapc.Timeout = 0;
				switch trapc.State {
				case State_Stop:
					print("\rStopping..     \n");
					trapc.Stop();
					break;
				case State_Warn:
					print("\rTerminate?     \n");
					break;
				default:
					if trapc.State < State_OK {
						trapc.State = State_OK;
					} else
					if trapc.State > State_Warn {
						print("\rTerminated!!!     \n");
						OS.Exit(0);
					}
					break;
				}
			}
		}
	}();
	return &trapc;
}



func (trapc *TrapC) NewStopChan() chan bool {
	stopchan := make(chan bool, 1);
	if trapc.Stopping.Load() {
		stopchan <-true;
	}
	trapc.StopChans.Append(stopchan);
	return stopchan;
}

func (trapc *TrapC) Wait() {
	trapc.WaitGroup.Wait();
}

func (trapc *TrapC) AddStopHook(hook StopHook) {
	if trapc.Stopping.Load() {
		hook();
	} else {
		trapc.StopHooks.Append(hook);
	}
}



func (trapc *TrapC) Stop() {
	trapc.Stopping.Store(true);
	for ; trapc.StopHooks.Length()>0; {
		stopchan, ok := trapc.StopChans.Get(0);
		if !ok { break; }
		trapc.StopChans.Remove(0);
		stopchan <-true;
	}
	for ; trapc.StopHooks.Length()>0; {
		hook, ok := trapc.StopHooks.Get(0);
		if !ok { break; }
		trapc.StopHooks.Remove(0);
		hook();
	}
}

func (trapc *TrapC) IsStopping() bool {
	return trapc.Stopping.Load();
}
