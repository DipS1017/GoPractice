package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
  greeting:="hello this is for testing the std library";
  fmt.Println(greeting);
  //contains function returns a boolean value
  fmt.Println(strings.Contains(greeting,"testing" ));
  
  //replaceall function replaces all specified values in the string

  fmt.Println(strings.ReplaceAll(greeting,"testing","learning"));
  //split function splits the string into an array of strings

  fmt.Println(strings.Split(greeting," "));
  // there are other functions like tolower and toupper and also trim and trim space

  //lets go with now implementing a these function in a int 

  arr:=[]int{4,5,1,2,3,4,5,5};
//sort functions sorts the array in ascending order
  fmt.Println(arr);
  sort.Ints(arr);
  fmt.Println(arr);
  //similary we can also search the indices of the slice or array it will return the first hit and duplicates dont matter
  fmt.Println(sort.SearchInts(arr,5));
  //this sort also works in slices  and we can also search strings and get their index

}

