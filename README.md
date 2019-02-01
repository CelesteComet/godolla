# GoDolla
GoDolla is a Go package designed to make writing Postgresql queries that involve variables less painful.


import "github.com/CelesteComet/godolla"

Just initialize a GoDolla object, and call the Serial method passing in an integer.


~~~go
import gd "github.com/CelesteComet/godolla"

func main() {
  gd := gd{}  
  log.Println(gd.Serial(3)) // produces the following string ($1, $2, $3)
}
~~~
