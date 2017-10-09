package rockets

import (
    "fmt"
    c "github.com/skilstak/go-colors"
    i "github.com/skilstak/go-input"
)

func Documentation() {
    fmt.Println(c.V + "Documentation [2017-10-08]")
    fmt.Println(c.C + "Likely to be outdated")
    fmt.Println()
    fmt.Println(c.C + "- h[elp]    Displays this documentation")
    fmt.Println(c.C + "- find rocket <id>    Finds rocket based on id")

    fmt.Println(c.C + "- change rocket manufacturer <manufacturer>   Finds rocket based on id")
    fmt.Println(c.C + "- add rocket <name> <height> <diameter> <manufacturer>")
    fmt.Println(c.C + "  L Adds rocket to database")
    fmt.Print(c.X)

    GetInput()
}

func Startup() {
    fmt.Print(c.Clear)
    fmt.Println(c.V + "Welcome to the rockets database CLI")
    fmt.Println(c.C + "Type 'h' or 'help' for help")
    fmt.Print(c.X + "\n")

    GetInput()
}

func GetInput() {
    var input string

    for {
        fmt.Print(c.Y + ">")
        input = i.Ask(c.B01 + " ")

        if input == "" {
            continue
        }

        break
    }

    switch input {
    case "h":
        fallthrough
    case "help":
        Documentation()
    }
}
