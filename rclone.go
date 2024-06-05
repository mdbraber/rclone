// Sync files and directories to and from local and remote object stores
//
// Nick Craig-Wood <nick@craig-wood.com>
// Maarten den Braber <m@mdbraber.com>
package rclone

import (
	_ "github.com/rclone/rclone/backend/all" // import all backends
	"github.com/rclone/rclone/cmd"
	_ "github.com/rclone/rclone/cmd/all"    // import all commands
	_ "github.com/rclone/rclone/lib/plugin" // import plugins
)

// #cgo CFLAGS: -x objective-c -I/Users/mdbraber/src/ios_system
// #cgo LDFLAGS: -L/Users/mdbraber/src/ios_system -lios_system
// #include <ios_system/ios_system.h>
// #include <ios_error.h>
// #include <stdlib.h>
// #define THREAD_STDIN_FILENO fileno(thread_stdin)
// #define THREAD_STDOUT_FILENO fileno(thread_stdout)
// #define THREAD_STDERR_FILENO fileno(thread_stderr)
import "C"
import "os"
import "unsafe"
import "log"
import "sync"
import "runtime"
import "fmt"

var wg sync.WaitGroup
var code int
var doneCustomExit bool = false

func Main() {
	return
}

//export GoMain
func GoMain(argc C.int, argv **C.char) int {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()	
	
	wg = sync.WaitGroup{}

	// Set outputs
	os.Stdin = os.NewFile(uintptr(C.THREAD_STDIN_FILENO), "thread_stdin")
	os.Stdout = os.NewFile(uintptr(C.THREAD_STDOUT_FILENO), "thread_stdout")
	os.Stderr = os.NewFile(uintptr(C.THREAD_STDERR_FILENO), "thread_stderr")

	// Logging is initialized before we could change it, so (re)set logging to os.Stderr
	log.SetOutput(os.Stderr)
	
	fmt.Println("GoMain")	
	
	// Create custom exit function
	// This function is empty now, just preventing syscall.exit() to be called
	// Calling ios_exit inside the custom function does not seem to be needed?
	os.CustomExitFunc = customExit

	// Threat this function as the main function with arguments
	args := make([]string, int(argc))
	ptr := uintptr(unsafe.Pointer(argv))
	for i := 0; i < int(argc); i++ {
		args[i] = C.GoString(*(**C.char)(unsafe.Pointer(ptr + uintptr(i)*unsafe.Sizeof(argv))))
	}

	// Set the arguments to os.Args
	os.Args = args
	fmt.Fprintf(os.Stdout, "args: %s\n", args)   

	// Call the main function in a separate Goroutine and wait until it finishes    
    wg.Add(1)
    go func(wg *sync.WaitGroup) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()	
		defer fmt.Println("WaitGroup done")

		doneCustomExit = false
		fmt.Println("WaitGroup")
        
		cmd.Main()
		if (!doneCustomExit) { wg.Done() }

		return
    }(&wg)
    wg.Wait()
        
	fmt.Println("GoMain done")

	//C.ios_exit(C.int(code))
	return 0
}

func customExit(code int) {
	fmt.Println("customExit")
	
	code = code
	doneCustomExit = true	
	
	wg.Done()
	C.ios_exit(C.int(code))

	return
}
