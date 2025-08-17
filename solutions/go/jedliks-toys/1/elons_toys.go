package elon

import "fmt"

// TODO: define the 'Drive()' method
func (self *Car) Drive() {
	if (self.battery < self.batteryDrain) {
        return
    }
    
    self.distance += self.speed
    self.battery -= self.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (self Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %v meters", self.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (self Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %v%%", self.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (self Car) CanFinish(trackDistance int) bool {
    remainingDistance := trackDistance - self.distance
    remainingBatteryCycles := self.battery / self.batteryDrain
    delta := remainingDistance - (self.speed * remainingBatteryCycles)
    
    return delta <= 0
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
