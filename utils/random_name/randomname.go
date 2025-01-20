package randomname

import (
	"math/rand"
)

// GenerateName 生成随机昵称
func GenerateName() string {
	var name string
	selectedType := RandomType(rand.Intn(2))
	switch selectedType {
	case adjectiveAndPerson:
		name = adjectiveSlice[rand.Intn(adjectiveSliceCount)] + personSlice[rand.Intn(PersonSliceCount)]
	case personActSomething:
		name = personSlice[rand.Intn(PersonSliceCount)] + actSomethingSlice[rand.Intn(actSomethingSliceCount)]
	}
	return name
}
