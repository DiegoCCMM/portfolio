package aima.core.environment.Fichas;

import aima.core.search.framework.GoalTest;

/**
 * @author Ravi Mohan
 * 
 */
public class FichasGoalTest implements GoalTest {
	FichasBoard goal = new FichasBoard(new int[] { 2, 2, 2, 0, 1, 1, 1});

	public boolean isGoalState(Object state) {
		FichasBoard board = (FichasBoard) state;
		return board.equals(goal);
	}
}