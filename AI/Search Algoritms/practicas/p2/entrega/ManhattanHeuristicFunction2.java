package aima.core.environment.eightpuzzle;

import aima.core.search.framework.HeuristicFunction;
import aima.core.util.datastructure.XYLocation;

/**
 * @author Ravi Mohan
 * 
 */
public class ManhattanHeuristicFunction2 implements HeuristicFunction {

	public double h(Object state) {
		EightPuzzleBoard board = (EightPuzzleBoard) state;
		EightPuzzleBoard goalboard = EightPuzzleGoalTest.goal;
		int retVal = 0;
		for (int i = 1; i < 9; i++) {
			XYLocation loc = board.getLocationOf(i);
			XYLocation locGoal = goalboard.getLocationOf(i);
			retVal += evaluateManhattanDistanceOf(locGoal, loc);
		}
		return retVal;

	}

	public int evaluateManhattanDistanceOf(XYLocation locGoal, XYLocation loc) {
		int retVal = -1;
		int xpos = loc.getXCoOrdinate();
		int ypos = loc.getYCoOrdinate();
		int xGoalpos = locGoal.getXCoOrdinate();
		int yGoalpos = locGoal.getYCoOrdinate();
		
		retVal = Math.abs(xpos - xGoalpos) + Math.abs(ypos - yGoalpos);
		
		return retVal;
	}
}