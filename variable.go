package main
import "fmt"
func main(){
var a string="hello this is a test for string"; // this is where the type is declared
var b ="this is a text"; // this is where the type is inferred both works and the space aren't registered in a new line in print unless \n is used or println is used
  c:="this is also a text"// this is the shorthand way : is used to declare and assign this will only be used when declaring and not the 2nd time 

  //also remember that the type of the variable cannot be changed once it is declared and the shorthand version cant be used outside the function
  fmt.Println(a,b,c);
  var d int = 10; //unlike JS var is a keyword in go and u can define the type as well
  var e=25;
  f:=30;
  fmt.Println(d,e,f);
  //bits and memory are also very imp to define 
  //int8 int16 int32 int64
  //similary there is unsigned version which means it cannt be negative and since we aren't using minuses we can go more that 127 range
  //uint8 uint16 uint32 uint64
var g int8=127;
var ab int16=231;
var ac int32=2400;
  fmt.Println(g,ab,ac);
  //similarly for int there is a float one too which is float 32 and float 64 mostly using float 64
  //string are inserted in go similar to c and c++ , in go we can use %v to print the value of the variable regarless of string ot int , but specifically we can use %s for string and %d for int
  //we can also puts ""quotes in the string to print the value of the variable using %q doesn't work for int
  fmt.Printf("the value of the variable is %q %v\n",a,a);
  //we can also use %T to print the type of the variable
  fmt.Printf("the value of the variable is %v and the type is %T\n",a,a);
  //similar to print f we can also use sprint f to store the value in a variable but this needs a var to store the value
  //arrays in go  
  //arrays are defined in go as [size]type{values} and the size is fixed and cannot be changed and also the the arrays are filled with 0 if not defined
  //u can change values of array with its index arr1[2]=5 , has to be in limit of out of bounds

  var arr [5]int=[5]int{21,2,34};
  arr[4]=3;
  //the array can be infered autmatically and declaring array on right
  var arr1=[3]string{"hello","world","this is a test"};
  fmt.Println(arr,arr1);

  //slices are arrays under the hood in golang and its similar to arrays as JS dev like me know it
  //its the same as arrays but the size is not defined and can be changed and also the slices are passed by reference

  //it also follows similar to arrays but we can append values as limit isn't defined BUT it returns a new slice and doesn't change the original slice but not a new var 
  var slice3=[]int{1,2,3,4,5,6,7,8,};
  var slice4=append(slice3,9);
  fmt.Println(slice3,slice4);
  //we also have ranges like in python and JS
  //it includes the first value and excludes the last value
  range1:=slice3[1:3];
  fmt.Println(range1);

    
 



}
