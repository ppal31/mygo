package seeds

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
)

type Seed struct {
	db *sqlx.DB
}

// Execute will executes the given seeder method
func Execute(db *sqlx.DB, seedMethodNames ...string) error {
	s := Seed{db}

	seedType := reflect.TypeOf(s)
	var err error
	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			err = seed(s, method.Name)
			if err != nil {
				return err
			}
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		err = seed(s, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func seed(s Seed, seedMethodName string) error {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		return errors.New("not a valid seed name")
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
	return nil
}
