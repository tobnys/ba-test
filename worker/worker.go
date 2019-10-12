package worker

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
)

type Worker struct {
	ID       int
	Navigate func()
}

func InitWorker() selenium.WebDriver {
	const (
		// These paths will be different on your system.
		seleniumPath = "vendor/selenium-server-standalone-3.141.59.jar"
		chromedriver = "vendor/chromedriver"
		port         = 8080
	)
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),         // Start an X frame buffer for the browser to run in.
		selenium.ChromeDriver(chromedriver), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),          // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browsername": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	// FIX THIS
	if err := nil {
		panic(err)
	}

	if err := worker.Get("http://play.golang.org/?simple=1"); err != nil {
		panic(err)
	}
}
