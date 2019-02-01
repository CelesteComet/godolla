package godolla 

import (
  "strings"
  "strconv"
)

type godolla struct {}

func (g *godolla) Serial(n int) string {
  sb := strings.Builder{}
  sb.WriteString("(")
  for i := 1; i <= n; i++ {
    sb2 := strings.Builder{}
    sb2.WriteString("$")
    sb2.WriteString(strconv.Itoa(i))
    if i != n {
      sb2.WriteString(", ")
    }
    sb.WriteString(sb2.String())
  }
  sb.WriteString(")")
  return sb.String();
}
