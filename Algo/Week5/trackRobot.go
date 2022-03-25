package week5

//This robot roams around a 2D grid. It starts at (0, 0) facing North. After each time it moves, the robot rotates 90 degrees clockwise. Given the amount the robot has moved each time, you have to calculate the robot's final position.
//
//To illustrate, if the robot is given the movements 20, 30, 10, 40 then it will move:
//
//20 steps North, now at (0, 20)
//30 steps East, now at (30, 20)
//10 steps South. now at (30, 10)
//40 steps West, now at (-10, 10)
//...and will end up at coordinates (-10, 10).

//Test.assertSimilar(trackRobot(20, 30, 10, 40), [-10, 10])
//Test.assertSimilar(trackRobot(10, -10, -10, 10), [-20, 20])
//Test.assertSimilar(trackRobot(),[0, 0])
//Test.assertSimilar(trackRobot(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), [6, 5])
//Test.assertSimilar(trackRobot(1, 0, 2, 0, 3, 0, 4, 0, 5, 0), [0, 3])
//Test.assertSimilar(trackRobot(0, 1, 0, 2, 0, 3, 0, 4, 0, 5), [3, 0])

func TrackRobot(args ...int) []int {
	var Xdeg, Ydeg int
	for ind, val := range args {

		ind = ind + 1
		if ind%2 == 0 {
			if ind%4 == 0 {
				Xdeg = Xdeg - val
			} else {
				Xdeg = Xdeg + val
			}

		} else {
			if ind%4 == 1 {
				Ydeg = Ydeg + val
			} else {
				Ydeg = Ydeg - val
			}

		}
	}
	return []int{Xdeg, Ydeg}
}
