package main
import "fmt"

//in go we also use functions to define the code and the main function is the entry point of the code similar to js 
//function that isn't called isn't a error in go
func main(){
//callign a function is also similar 
  greet("dips");

  bye("dips");

  //we can also pass a function as a parameter to another function callback function basically
  cycleNames([]string{"dips","dave","david"},greet)

//when calling a func with return type we need to store it in vallue for sideeffect to be seen
  r1:=radius(5);
  fmt.Println(r1);


}
//Declaring a function is very similar to js with func keyword
func greet(u string){
  fmt.Printf("hello %v\n",u)
}
func bye(u string){
  fmt.Printf("bye %v\n",u)
}
func cycleNames(n[]string,f func(string)){
  for _,name:=range n{
    f(name);

  }

}
//now if we were to get a function with a return type i would have to define its type outside parameter
func radius(r int)int{
  return r*r; 

}

