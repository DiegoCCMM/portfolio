package aima.core.environment.Canibales;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import aima.core.agent.Action;
import aima.core.agent.impl.DynamicAction;
import aima.core.util.datastructure.XYLocation;

/**
 * @author Diego Caballé 738712
 */
public class CanibalesBoard {

	//Movimientos que empiezan en la izquierda
	public static Action IZQ_M2 = new DynamicAction("IZQ_M2");
	public static Action IZQ_C2 = new DynamicAction("IZQ_C2");
	public static Action IZQ_M1 = new DynamicAction("IZQ_M1");
	public static Action IZQ_C1 = new DynamicAction("IZQ_C1");
	public static Action IZQ_MC = new DynamicAction("IZQ_MC");

	//Movimientos que empiezan en la derecha 
	public static Action DER_M2 = new DynamicAction("DER_M2");
	public static Action DER_C2 = new DynamicAction("DER_C2");
	public static Action DER_M1 = new DynamicAction("DER_M1");
	public static Action DER_C1 = new DynamicAction("DER_C1");
	public static Action DER_MC = new DynamicAction("DER_MC");

	private int[] state;

	//
	// PUBLIC METHODS
	//

	public CanibalesBoard() {
		state = new int[] { 3, 3, 1, 0, 0};	//M,C,BARCA,M,C BARCA = 1 -> izquierda, 0 -> derecha
	}

	public CanibalesBoard(int[] state) {
		this.state = new int[state.length];
		System.arraycopy(state, 0, this.state, 0, state.length);
	}

	public CanibalesBoard(CanibalesBoard copyBoard) {
		this(copyBoard.getState());
	}

	public int[] getState() {
		return state;
	}

	public void move(Action where) {
		int Morig, Corig, Mdest, Cdest;
		
		if(leftToRightMove(where)) {
			Morig = state[0];
			Corig = state[1];
			Mdest = state[3];
			Cdest = state[4];
			
			Morig = Morig - numberOfM(where);
			Mdest = Mdest + numberOfM(where);
			Corig = Corig - numberOfC(where);
			Cdest = Cdest + numberOfC(where);
			
			state[0] = Morig;
			state[1] = Corig;
			state[3] = Mdest;
			state[4] = Cdest;
			
			state[2] = 0;
		}else {
			Mdest = state[0];
			Cdest = state[1];
			Morig = state[3];
			Corig = state[4];			
			
			Morig = Morig - numberOfM(where);
			Mdest = Mdest + numberOfM(where);
			Corig = Corig - numberOfC(where);
			Cdest = Cdest + numberOfC(where);
			
			state[0] = Mdest;
			state[1] = Cdest;
			state[3] = Morig;
			state[4] = Corig;
			
			state[2] = 1;
		}
	}
	
	public boolean peopleCanGetInAndGo(Action where) {
		int Morig, Corig, Mdest, Cdest;
		if(leftToRightMove(where)) {
			Morig = state[0];
			Corig = state[1];
			Mdest = state[3];
			Cdest = state[4];
			
		}else {
			Mdest = state[0];
			Cdest = state[1];
			Morig = state[3];
			Corig = state[4];
			
		}
		
		if(Morig >= numberOfM(where) && Corig >= numberOfC(where)) {
			int M = numberOfM(where);
			int C = numberOfC(where);
			
			if(C>= 1 && M == 0) {
				return (Mdest >= Cdest + C || Mdest == 0);
			}else if (M>= 1 && C == 0) {
				return ((Morig >= Corig + M || Morig == M) && Mdest >= Cdest - M);
			}else if( M == 1 & C == 1) {
				 return(Mdest >= Cdest);
			}
			
		}
		return false;
		
	}
	public int numberOfM(Action where) {
		if(where.equals(IZQ_M2) || where.equals(DER_M2)) {
			return 2;
		}else if(where.equals(IZQ_MC) || where.equals(DER_MC) ||
				where.equals(IZQ_M1) || where.equals(DER_M1)) {
			return 1;
		}else
			return 0;
		
		
	}
	public int numberOfC(Action where) {
		if(where.equals(IZQ_C2) || where.equals(DER_C2)) {
			return 2;
		}else if(where.equals(IZQ_MC) || where.equals(DER_MC) ||
				where.equals(IZQ_C1) || where.equals(DER_C1)) {
			return 1;
		}else 
			return 0;
		
	}
	
	public boolean boatIsThere(Action where) {
		boolean sol = true;
		int boat = getBoatPosition();
		
		if(boat == 0  && leftToRightMove(where))
		{
			sol = false;
		}else if(boat == 1 && rightToLeftMove(where))
		{
			sol = false;
		}
		
		return sol;
		
	}
	public boolean leftToRightMove(Action where) {
		
		if(where.equals(IZQ_M2) ||  
		where.equals(IZQ_C2) ||
		where.equals(IZQ_M1) ||
		where.equals(IZQ_C1) ||
		where.equals(IZQ_MC)) {
			return true;
		}else {return false;}
		
		
	}
	public boolean rightToLeftMove(Action where) {
		return !leftToRightMove(where);
	}
	
	@Override
	public boolean equals(Object obj) {
		if (this == obj)
			return true;
		if (obj == null)
			return false;
		if (getClass() != obj.getClass())
			return false;
		CanibalesBoard other = (CanibalesBoard) obj;
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
		String retVal="RIBERA-IZQ ";
		retVal+=mostrarMisioneros(state[0]);
		retVal+=mostrarCanibales(state[1]);
		if(state[2]==1) {
			retVal+=" BOTE ---RIO--- ";
		}
		else {
			retVal+="---RIO---  BOTE ";
		}
		retVal+=mostrarMisioneros(state[3]);
		retVal+=mostrarCanibales(state[4]);
		retVal+=" RIBERA-DCH";
		
		return retVal;
	}
	String mostrarCanibales(int n) {
		String d="";
		if(n==1) {
			d="C   ";
		}else if(n==2) {
			d="CC  ";
			
		}else if(n==3) {
			d="CCC ";
		}else {
			d="    ";
		}	
		return d;
	}
	String mostrarMisioneros(int n) {
		String d="";
		if(n==1) {
			d="M   ";
		}else if(n==2) {
			d="MM  ";
			
		}else if(n==3) {
			d="MMM ";
		}else {
			d="    ";
		}		
		return d;
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
	private int getBoatPosition() {
		return state[2];
	}
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