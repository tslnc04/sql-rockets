package rockets

import (
    "fmt"
    "strings"
    "strconv"
    "database/sql"
    "os"

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
    fmt.Println(c.C + "- change rocket manufacturer <id> <manufacturer>   Finds rocket based on id")
    fmt.Println(c.C + "- e[x]it   Exit the program")
    fmt.Println(c.C + "- clear   Clear the screen")
    fmt.Println(c.C + "- upsert rocket <id> <name> <height> <diameter> <manufacturer>")
    fmt.Println(c.C + "  L Adds rocket unless it exist, where it will update all rockets")
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
    } else if strings.HasPrefix(input, "change rocket manufacturer") {
        ChangeManufacturer(db, input)
    } else if strings.HasPrefix(input, "add rocket") {
        AddRocketCLI(db, input)
    } else if strings.HasPrefix(input, "upsert rocket") {
        UpsertRocket(db, input)
    }

    switch input {
    case "h", "help":
        Documentation(db)
    case "clear":
        fmt.Println(c.Clear)
    case "exit", "x":
        fmt.Println(c.C + "Bye")
        os.Exit(0)
    }

    GetInput(db)
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

    GetInput(db)
}

// ChangeManufacturer is the CLI handler for changing the rocket manufacturer
func ChangeManufacturer(db *sql.DB, input string) {
    splitInput := strings.Split(input, " ")

    rocketIDString := splitInput[3]
    rocketID, err := strconv.Atoi(rocketIDString)

    if err != nil {
        panic(err)
    }

    rocketManufacturer := strings.Join(splitInput[4:], " ")
    success := ChangeRocketManufacturer(db, rocketID, rocketManufacturer)

    if success {
        fmt.Println(c.C + "Manufacturer successfully updated.")
    }

    GetInput(db)
}

// AddRocketCLI is the CLI handler for adding a rocket to the databse
// Note: CLI is added at the end due to AddRocket already existing, it should really be added to all
func AddRocketCLI(db *sql.DB, input string) {
    splitInput := strings.Split(input, " ")

    rocketName := splitInput[2]
    rocketHeight, err := strconv.ParseFloat(splitInput[3], 32)

    if err != nil {
        panic(err)
    }

    rocketDiameter, err := strconv.ParseFloat(splitInput[4], 32)

    if err != nil {
        panic(err)
    }

    rocketManufacturer := strings.Join(splitInput[5:], " ")
    success := AddRocket(db, rocketName, float32(rocketHeight), float32(rocketDiameter), rocketManufacturer)

    if success {
        fmt.Println(c.C + "Rocket added successfully.")
    }

    GetInput(db)
}

// UpsertRocket is the handler for performing an upsert
func UpsertRocket(db *sql.DB, input string) {
    splitInput := strings.Split(input, " ")

    rocketIDString := splitInput[2]
    rocketID, err := strconv.Atoi(rocketIDString)

    if err != nil {
        panic(err)
    }

    rocketName := splitInput[3]
    rocketHeight, err := strconv.ParseFloat(splitInput[4], 32)

    if err != nil {
        panic(err)
    }

    rocketDiameter, err := strconv.ParseFloat(splitInput[5], 32)

    if err != nil {
        panic(err)
    }

    rocketManufacturer := strings.Join(splitInput[6:], " ")
    success := AddOrUpdateRocket(db, rocketID, rocketName, float32(rocketHeight), float32(rocketDiameter), rocketManufacturer)

    if success {
        fmt.Println(c.C + "Rocket upserted successfully.")
    }

    GetInput(db)
}
