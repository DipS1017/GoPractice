package main
import "fmt"

func mains(){
  //very similar to js or c on looping 

  for x:=0;x<5;x++{
    fmt.Print(x );

  }

  //this is for loopping over array ,imperative way
  slice:=[]string{"dave","snla","sldkf","akdf","adkjf"}

  fmt.Println(slice);
  for i:=0;i<len(slice);i++{
    fmt.Print(slice[i]);

  }
  //in Python we get too loop a array with for in arr but in go we have to use range
  for index,value:=range slice{
    fmt.Println(index,value);

  }
  //in range we get 2 return values one is the index and the other is the value since we get 2 return it will create error if unused so in that case we can replace with _ to ignore the value

  for _,value:=range slice{
    fmt.Println(value);

  }

  //we can also loop over with maps as well

  //lets first just do maps
  menu:=map[string]float64{
    "pie":5.99,
    "cake":3.99,
    "ice cream":2.99,
    "milkshake":4.99,
  }
  fmt .Println(menu);
  fmt.Println(menu["ice cream"]);
  //lets now loop over the maps

  for k,v:=range menu{
    fmt.Printf("%v: $%v\n",k,v);

  }


}
