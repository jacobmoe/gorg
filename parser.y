%{

package gorg

import (
  "fmt"
)

%}

%union {
	item *lexItem
}

%start main
%token <item> i

%%

main: i i i i i i
  {
    print($1)
    print($2)
    print($3)
    print($4)
    print($5)
    print($6)
  }

%%

func print(b interface{}) {
  fmt.Println(b)
}
