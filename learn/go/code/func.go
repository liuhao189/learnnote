package main 

import "fmt";
import "math";

func max(num1,num2 int) int{
	if(num1>num2){
		return num1;
	}else{
		return num2;
	}
}

func swap(x,y string) (string,string){
	return y,x;
}

func getSequence() func() int{
	i:=0;
	return func() int{
		i++;
		return i;
	}
}

type Circle struct{
	radius float64;
}

func (c Circle) getArea() float64{
	return 3.14 * c.radius * c.radius;
}

func main(){
	var maxNum=max(10,15);
	fmt.Printf("Max Num is %d\n",maxNum);
	a,b:=swap("Google","Bing");
	fmt.Println(a);
	fmt.Println(b);
	getSquareRoot := func(x float64) float64{
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9));

	nextNumber:=getSequence();
	fmt.Println(nextNumber());
	fmt.Println(nextNumber());
	fmt.Println(nextNumber());

	nextNumber1:=getSequence();
	fmt.Println(nextNumber1());
	fmt.Println(nextNumber1());

	fmt.Println("****************")

	var c1 Circle;
	c1.radius=10.00;
	fmt.Printf("Area is %f.\n",c1.getArea());
}