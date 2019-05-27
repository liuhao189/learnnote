package main

import "fmt";

var k=999;

func main(){
	var i=666;
	// var i=555;
	k:=777;
	// j:=888;
	fmt.Println(&i,i,k);
}
