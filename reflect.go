package main  
import(  
   "fmt"  
   "reflect"  
)  
func main()  {  
   age := 27.5  
   fmt.Printf("%v\n" ,age)  
   fmt.Println(reflect.TypeOf(age))  
}  