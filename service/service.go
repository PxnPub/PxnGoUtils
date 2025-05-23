package service;

import(
	OS    "os"
	Time  "time"
	TrapC "github.com/PxnPub/pxnGoUtils/trapc"
);



func Pre() *TrapC.TrapC  {
	print("\n");
	trapc := TrapC.New();
	return trapc;
}

func Post(trapc *TrapC.TrapC) {
	Time.Sleep(Time.Duration(250) * Time.Millisecond);
	print("\n"); trapc.Wait();
	print(" ~end~ \n");
	print("\n"); OS.Exit(0);
}
