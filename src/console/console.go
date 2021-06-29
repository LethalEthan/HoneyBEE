package console

import (
	"bufio"
	"config"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"nserver"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"server"
	"syscall"
	"utils"

	logging "github.com/op/go-logging"
)

var (
	conf     *config.Config
	shutdown = make(chan os.Signal, 1)
	Log      = logging.MustGetLogger("HoneyGO")
	Panicked bool
	hprof    *os.File
	cprof    *os.File
)

func Console() {
	defer DRECOVER()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "help":
			Log.Warning("There is no help atm :(")
			Log.Warning("This is a simple, quick and dirty way of doing commands, a proper thing is being made later")
		case "shutdown":
			shutdown <- os.Interrupt
		case "stop":
			shutdown <- os.Interrupt
		case "exit":
			shutdown <- os.Interrupt
		case "reload":
			if !conf.DEBUGOPTS.NewServer {
				server.SetNetServerRun(false)
				server.SetRun(false)
				config.ConfigReload()
				server.GCPShutdown <- true
				server.ServerReload()
				server.SetRun(true)
				server.SetNetServerRun(true)
			} else {
				config.ConfigReload()
				//nserver.Conf = config.GConfig
				if conf.Performance.CPU == 0 {
					Log.Info("Setting GOMAXPROCS to all available logical CPU's")
					runtime.GOMAXPROCS(runtime.NumCPU()) //Set it to the value of how many cores
				} else {
					Log.Info("Setting GOMAXPROCS to config: ", conf.Performance.CPU)
					runtime.GOMAXPROCS(conf.Performance.CPU)
				}
				if runtime.NumCPU() <= 3 || conf.Performance.CPU <= 2 {
					Log.Critical("Number of CPU's is less than 3 this could impact performance as this is a heavily threaded application")
				}
				if conf.Server.ClientFrameBuffer == 0 || conf.Server.ReadBufferCap == 0 || conf.Server.RecieveBuf == 0 || conf.Server.SendBuf == 0 || conf.Server.Timeout == 0 {
					panic("Please don't be stupid and set the buffers or timeout as 0 :/")
				}
				Log.Critical("If you changed new server to old server this will not be reloaded or changed!")
			}
		case "GC":
			runtime.GC()
			Log.Info("GC invoked")
		case "mem":
			utils.PrintDebugStats()
		case "SSM":
			Log.Debug(server.StatusCache)
		case "CCM":
			Log.Debug(server.ClientConnectionMap)
		case "panic":
			panic("panicked, you told me to :)")
		case "cpuprofile":
			if config.Cpuprofile != "" {
				pprof.StopCPUProfile()
				cprof.Close() // error handling omitted for example
				Log.Warning("Written CPU Profile")
			} else {
				Log.Critical("cpuprofile flag not specified! not writing a profile")
			}
		case "memprofile":
			if config.Memprofile != "" {
				runtime.GC() // get up-to-date statistics
				if err := pprof.WriteHeapProfile(hprof); err != nil {
					log.Fatal("could not write memory profile: ", err)
				} else {
					Log.Warning("Written Memory Profile")
				}
			} else {
				Log.Critical("memprofile flag not specified! not writing a profile")
			}
		case "profile":
			if config.Cpuprofile != "" {
				pprof.StopCPUProfile()
				cprof.Close() // error handling omitted for example
				Log.Warning("Written CPU Profile")
			} else {
				Log.Critical("cpuprofile flag not specified! not writing a cpuprofile")
			}
			//
			if config.Memprofile != "" {
				runtime.GC() // get up-to-date statistics
				if err := pprof.WriteHeapProfile(hprof); err != nil {
					log.Fatal("could not write memory profile: ", err)
				} else {
					Log.Warning("Written Memory Profile")
				}
			} else {
				Log.Critical("memprofile flag not specified! not writing a cpuprofile")
			}
		default:
			Log.Warning("Unknown command")
		}
	}
}

var MD5 string

func Hash() string {
	file, err := os.Open(os.Args[0])
	if err != nil {
		MD5 = "00000000000000000000000000000000"
	}
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		MD5 = "00000000000000000000000000000000"
	}
	//Get the 16 bytes hash
	hBytes := hash.Sum(nil)[:16]
	file.Close()
	MD5 = hex.EncodeToString(hBytes) //Convert bytes to string
	return MD5
}

//Shutdown - listens for sigterm and exits
func Shutdown() {
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	select {
	case <-shutdown:
		{
			Log.Warning("Starting shutdown")
			if !config.GConfig.DEBUGOPTS.NewServer {
				server.SetNetServerRun(false)
				server.SetRun(false)
			} else {
				nserver.GlobalServer.Shutdown()
			}
			if config.GConfig.Performance.EnableGCPlayer { //Don't send on an unintialised channel otherwise it will deadlock
				go func() { nserver.GCPShutdown <- true }()
			}
			server.ClientConnectionMutex.Lock()
			Log.Debug(server.ClientConnectionMap) //Check if any connections are still active in the map, there shouldn't be any left over
			DEBUG := true
			if DEBUG {
				utils.PrintDebugStats()
			}
			os.Exit(0)
		}
	}
}
