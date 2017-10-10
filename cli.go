package rockets

import (
    "fmt"
    "strings"
    "strconv"
    "database/sql"

    c "github.com/skilstak/go-colors"
    i "github.com/skilstak/go-input"
)

// Documentation prints out the documentation to the command
func Documentation(db *sql.DB) {
    fmt.Println(c.V + "Documentation [2017-10-08]")
    fmt.Println(c.C + "Likely to be outdated")
    fmt.Println()
    fmt.Println(c.C + "- h[elp]    Displays this documentation")
    fmt.Println(c.C + "- find rocket <id>    Finds rocket based on id")

    fmt.Println(c.C + "- change rocket manufacturer <manufacturer>   Finds rocket based on id")
    fmt.Println(c.C + "- add rocket <name> <height> <diameter> <manufacturer>")
    fmt.Println(c.C + "  L Adds rocket to database")
    fmt.Print(c.X)

    GetInput(db)
}

// Startup is the function that gets run to start the CLI
func Startup(db *sql.DB) {
    fmt.Print(c.Clear)
    fmt.Println(c.V + "Welcome to the rockets database CLI")
    fmt.Println(c.C + "Type 'h' or 'help' for help")
    fmt.Print(c.X + "\n")

    GetInput(db)
}

// GetInput is the main input loop
func GetInput(db *sql.DB) {
    var input string

    for {
        fmt.Print(c.Y + ">")
        input = i.Ask(c.B01 + " ")

        if input == "" {
            continue
        }

        break
    }

    if strings.HasPrefix(input, "find rocket") {
        FindRocket(db, input)
    }

    switch input {
    case "h":
        fallthrough
    case "help":
        Documentation(db)
    }
}

// FindRocket handles finding the rocket by ID
func FindRocket(db *sql.DB, input string) {
    rocketIDString := strings.Split(input, " ")[2]
    rocketID, err := strconv.Atoi(rocketIDString)

    if err != nil {
        panic(err)
    }

    foundRocket := FindRocketByID(db, rocketID)
    fmt.Print(c.C + foundRocket.String())
}
