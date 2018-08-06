package main


/*
#include <unistd.h>


	void shitboss(){
		write(1,"boss is shit\n",13);
	}
 */
import "C"


func main(){
	C.shitboss()
}