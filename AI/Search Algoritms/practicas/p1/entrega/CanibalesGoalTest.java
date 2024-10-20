package aima.core.environment.Canibales;

import aima.core.search.framework.GoalTest;

/**
 * @author Ravi Mohan
 * 
 */
public class CanibalesGoalTest implements GoalTest {
	CanibalesBoard goal = new CanibalesBoard(new int[] { 0, 0, 0, 3, 3});

	public boolean isGoalState(Object state) {
		CanibalesBoard board = (CanibalesBoard) state;
		return board.equals(goal);
	}
}