package second

import "fmt"

type hooper struct {
    person
    nationality country
    offense uint8
    defense uint8
    overall uint8
}

type person struct {
    name string
    heightInCm uint8
    weightInKg uint8
    favoriteFood string
}

type country struct {
    continent string
    population uint32
}

func (h1 hooper) isBetterOverall(h2 hooper) bool {
    ok := h1.overall > h2.overall
    fmt.Printf("is %s better than %s?:\n%v\n", h1.name, h2.name, ok)
    return ok
}

func ThirdExercise() {
    melo := hooper{
            person: person{
                name: "Carmelo Anthony",
                heightInCm: 220,
                weightInKg: 100,
                favoriteFood: "nachos",
            },
            nationality: country{
                continent: "NA",
                population: 123456789,
            },
            offense: 92,
            defense: 63,
            overall: 81,
    }

    jordan := hooper{
        person: person{
            name: "Michael Jordan",
            heightInCm: 210,
            weightInKg: 90,
            favoriteFood: "winning",
        },
        nationality: country{
            continent: "NA",
            population: 123456789,
        },
        offense: 99,
        defense: 99,
        overall: 99,
    }

    fmt.Printf("%+v\n", melo)
    fmt.Printf("%s\n", melo.nationality.continent)
    
    jordan.isBetterOverall(melo)
    
}
