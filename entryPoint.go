package gogame

import (
	"flag"
	"os"
	"runtime/pprof"

	"github.com/oyberntzen/gogame/ggcore"
	"github.com/oyberntzen/gogame/ggdebug"
)

var profile = flag.Bool("profile", false, "set to true to enable profiling")

func CreateApplication(clientApp ClientApplication) {
	flag.Parse()
	if *profile {
		ggcore.CoreInfo("Starting profiling")

		file, err := os.Create("cpuprofile.pprof")
		ggcore.CoreCheckError(err)
		defer file.Close()

		ggcore.CoreCheckError(pprof.StartCPUProfile(file))
		defer pprof.StopCPUProfile()
	}

	ggdebug.BeginSession("startup", "startup.json")
	app := newCoreApplication(clientApp)
	ggdebug.EndSession()

	ggdebug.BeginSession("loop", "loop.json")
	app.run()
	ggdebug.EndSession()
}
