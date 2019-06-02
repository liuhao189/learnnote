package main

import "fmt";

func main() {
	var grade string ="B";
	var marks int = 90;
	switch marks{
	case 90:grade="A";
	case 80:grade="B";
	case 50,60,70:grade="C";
	default:grade="D";
	}

	switch{
	case grade == "A":
		fmt.Printf("优秀\n");
	case grade=="B":
		fmt.Printf("良好\n");
	case grade=="C":
		fmt.Printf("及格\n");
	case grade=="D":
		fmt.Printf("不及格\n");
	}
	fmt.Printf("Your grade is %s\n",grade);
}

// func print(c int){
// 	fmt.Printf("val is %d\n",c);
// }

// func main() {
// 	var a int =21;
// 	var b int =10;
// 	var c int;
// 	c=a+b;
// 	print(c);
// 	c=a-b;
// 	print(c);
// 	c=a*b;
// 	print(c);
// 	c=a/b;
// 	print(c);
// 	c=a%b;
// 	print(c);
// 	a++;
// 	print(a);
// 	a=21;
// 	a--;
// 	print(a);
// 	var a int=21;
// 	var b int=10;
// 	if(a==b){
// 		fmt.Printf("a == b is true\n")
// 	}else{
// 		fmt.Printf("a == b is false\n")
// 	}
// 	if(a<b){
// 		fmt.Printf("a < b is true \n")
// 	}else{
// 		fmt.Printf("a < b is false \n")
// 	}

// 	if(a>b){
// 		fmt.Printf("a > b is true\n")
// 	}else{
// 		fmt.Printf("a>b is false\n")
// 	}

// 	var a bool=true;
// 	var b bool=false;

// 	if(a&&b){
// 		fmt.Printf("a && b is true\n");
// 	}else{
// 		fmt.Printf("a && b is false!\n");
// 	}

// 	if(a || b){
// 		fmt.Printf("a || b is true!\n");
// 	}

// 	a=false;
// 	b=true;
// 	if(!(a&&b)){
// 		fmt.Printf("!(a&&b) is true!");
// 	}

// 	var a uint = 60;//0011 1100 
// 	var b uint = 13;//0000 1101
// 	var c uint = 0;// 0000 0000

// 	c= a &b;
// 	fmt.Printf("c is %d\n",c);//12
// 	c= a| b;
// 	fmt.Printf("c is %d\n",c);//61
// 	c=a ^b;
// 	fmt.Printf("c is %d\n",c);//49
// 	c=a<<2;
// 	fmt.Printf("c is %d\n",c)//240
// 	c=a>>2;
// 	fmt.Printf("c is %d\n",c);//15

// 	var a =4;
// 	var b int32;
// 	var c float32;
// 	var ptr *int;
// 	fmt.Printf("a type is %T\n",a);
// 	fmt.Printf("b type is %T\n",b);
// 	fmt.Printf("c type is %T\n",c);
// 	ptr=&a;
// 	fmt.Printf("a value is %d\n",a);
// 	fmt.Printf("*ptr is %d\n",*ptr);
// }