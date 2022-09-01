package robot_test

import (
	"fmt"

	"github.com/DazFather/parrbot/robot"
)

func ExampleKeepActiveSessions() {
	robot.Config.KeepActiveSessions()
	fmt.Println(robot.Config.DeleteSessionTimer) // Output: 0s
}

func ExampleSetAPIToken() {
	var (
		wrongToken string = "123456"
		myToken    string = "123:ABC"
		err        error
	)

	err = robot.Config.SetAPIToken(wrongToken)
	fmt.Println(err)

	err = robot.Config.SetAPIToken(myToken)
	fmt.Println(err)

	// Output:
	// Wrong format for TOKEN value
	// <nil>
}
