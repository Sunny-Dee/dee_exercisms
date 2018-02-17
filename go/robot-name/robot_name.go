package robotname

import (
	"math/rand"
	"strconv"
)

const base = 10
const max = 999
const min = 100

// Robot holds robot information such as name
type Robot struct {
	RobotName string
}

// Name returns the robot name. Generates one if name is empty.
func (r *Robot) Name() string {
	if r.RobotName == "" {
		r.generateName()
	}
	return r.RobotName
}

// Reset robot name to empty string.
func (r *Robot) Reset() {
	r.RobotName = ""
}

func (r *Robot) generateName() {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r.RobotName = string(letters[rand.Int()%(len(letters)-1)]) +
		string(letters[rand.Int()%(len(letters)-1)]) +
		strconv.FormatInt(int64(rand.Intn(max-min)+min), base)
}
