package interfaces

import (
	"fmt"
)

type car struct {
    make string
    year uint32
    licenseRequired string
    numWheels uint8
}

type truck struct {
    year uint32
    longHaul bool
    commercial bool
}

type vehicle interface {
    getInfo()
}

func (c car) getInfo() {
    fmt.Printf(
        "This car is a %d %s requiring a %s license with %d wheels\n",
        c.year,
        c.make,
        c.licenseRequired,
        c.numWheels,
    )
}

func (t truck) getInfo() {
    fmt.Printf(
        "This truck was built in %d. Is it long haul?: %t. Is it commerical?: %t\n",
        t.year,
        t.longHaul,
        t.commercial,
    )
}

func printInfo(v vehicle) {
    v.getInfo()

//     if _, ok := v.(car); ok {
//         fmt.Println("this is a car")
//     }
//     if _, ok := v.(truck); ok {
//         fmt.Println("this is a truck")
//     }
    switch v.(type) {
    case car:
        fmt.Println("this is car")
    case truck:
        fmt.Println("this is a truck")
    default:
        fmt.Println("this neither truck nor car")
    }
}

func RunInterfacesDemo() {
    c := car{
        make: "Toyota",
        year: 1992,
        licenseRequired: "Class D",
        numWheels: 4,
    }

    t := truck{
        year: 1950,
        longHaul: true,
        commercial: true,
    }

    printInfo(c)
    printInfo(t)

}
