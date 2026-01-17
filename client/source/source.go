// This file contains the source the client will send to the server and then be executed by the server.
package source

import "log"

func Function() {
	log.Println("This is working")
	function()
}

func function() {
	log.Println("Inside little function")
}
