package aima.core.environment.Fichas;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import aima.core.agent.Action;
import aima.core.agent.impl.DynamicAction;
import aima.core.util.datastructure.XYLocation;

/**
 * @author Diego Caballé 738712
 */
public class FichasBoard {

	public static Action L1 = new DynamicAction("L1");		//no salta fichas

	public static Action R1 = new DynamicAction("R1");
	
	public static Action L2 = new DynamicAction("L2");		//salta una ficha

	public static Action R2 = new DynamicAction("R2");
	
	public static Action L3 = new DynamicAction("L3");		//salta dos fichas

	public static Action R3 = new DynamicAction("R3");
	

	private int[] state;

	//
	// PUBLIC METHODS
	//

	public FichasBoard() {
		state = new int[] { 1, 1, 1, 0, 2, 2, 2};	//M,C,BARCA,M,C BARCA = 1 -> izquierda, 0 -> derecha
	}

	public FichasBoard(int[] state) {
		this.state = new int[state.length];
		System.arraycopy(state, 0, this.state, 0, state.length);
	}

	public FichasBoard(FichasBoard copyBoard) {
		this(copyBoard.getState());
	}

	public int[] getState() {
		return state;
	}

	public void move(Action where) {
		int gap = getGap();
		int move = getMove(where);
		
		state[gap] = state[gap + move];
		
		state[gap + move] = 0;
		
	}
	
	public boolean canMove(Action where) {
		int gap = getGap();
		int move = getMove(where);
		
		if (gap + move > 6 || gap + move < 0)
		{
			return false;
		}else {
			return true;
		}
		
	}
	
	private int getMove(Action where) {
		if(where.equals(L1)) {
			return(-1);
		}else if(where.equals(L2)) {
			return(-2);
		}else if(where.equals(L3)) {
			return(-3);
		}else if(where.equals(R1)) {
			return(1);
		}else if(where.equals(R2)) {
			return(2);
		}else if(where.equals(R3)) {
			return(3);
		}
		return(0);
	}
	
	private int getGap() {
		int i = 0;
		while (i < 7)
		{
			if(state[i] == 0) {
				return i;
			}
			i++;
		}
		return 0;
	}
	
	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		FichasBoard other = (FichasBoard) obj;
		if (!Arrays.equals(state, other.state))
			return false;
		return true;
	}

	@Override
	public int hashCode() {
		final int prime = 31;
		int result = 1;
		result = prime * result + Arrays.hashCode(state);
		return result;
	}

	@Override
	public String toString() {
		String retVal="|";
		for(int i = 0; i < 7; i++) {
			if(state[i] == 1 )
			{
				retVal += "B|";
			}else if(state[i] == 2) {
				retVal += "V|";
			}else if(state [i] == 0) {
				retVal += " |";
			}
		}
		
		return retVal;
	}

//	//
//	// PRIVATE METHODS
//	//
//
//	/**
//	 * Note: The graphic representation maps x values on row numbers (x-axis in
//	 * vertical direction).
//	 */
//	private int getXCoord(int absPos) {
//		return absPos / 3;
//	}
//
//	/**
//	 * Note: The graphic representation maps y values on column numbers (y-axis
//	 * in horizontal direction).
//	 */
//	private int getYCoord(int absPos) {
//		return absPos % 3;
//	}
//
//	private int getAbsPosition(int x, int y) {
//		return x * 3 + y;
//	}
//
//	private int getValueAt(int x, int y) {
//		// refactor this use either case or a div/mod soln
//		return state[getAbsPosition(x, y)];
//	}
//
//
//	private int getPositionOf(int val) {
//		int retVal = -1;
//		for (int i = 0; i < 9; i++) {
//			if (state[i] == val) {
//				retVal = i;
//			}
//		}
//		return retVal;
//	}
//
//	private void setValue(int x, int y, int val) {
//		int absPos = getAbsPosition(x, y);
//		state[absPos] = val;
//
//	}
	
}