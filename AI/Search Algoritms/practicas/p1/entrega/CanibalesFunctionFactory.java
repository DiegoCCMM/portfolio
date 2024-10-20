package aima.core.environment.Canibales;

import java.util.LinkedHashSet;
import java.util.Set;

import aima.core.agent.Action;
import aima.core.search.framework.ActionsFunction;
import aima.core.search.framework.ResultFunction;

/**
 * @author Diego Caballé 738712
 */
public class CanibalesFunctionFactory {
	private static ActionsFunction _actionsFunction = null;
	private static ResultFunction _resultFunction = null;

	public static ActionsFunction getActionsFunction() {
		if (null == _actionsFunction) {
			_actionsFunction = new EPActionsFunction();
		}
		return _actionsFunction;
	}

	public static ResultFunction getResultFunction() {
		if (null == _resultFunction) {
			_resultFunction = new EPResultFunction();
		}
		return _resultFunction;
	}

	private static class EPActionsFunction implements ActionsFunction {
		public Set<Action> actions(Object state) {
			CanibalesBoard board = (CanibalesBoard) state;

			Set<Action> actions = new LinkedHashSet<Action>();

			if (board.peopleCanGetInAndGo(CanibalesBoard.DER_M2)) {
				actions.add(CanibalesBoard.DER_M2);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.DER_M1))  {
				actions.add(CanibalesBoard.DER_M1);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.DER_C1))  {
				actions.add(CanibalesBoard.DER_C1);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.DER_C2))  {
				actions.add(CanibalesBoard.DER_C2);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.DER_MC))  {
				actions.add(CanibalesBoard.DER_MC);
			}
			
			if (board.peopleCanGetInAndGo(CanibalesBoard.IZQ_M2))  {
				actions.add(CanibalesBoard.IZQ_M2);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.IZQ_M1 ))  {
				actions.add(CanibalesBoard.IZQ_M1);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.IZQ_C1))  {
				actions.add(CanibalesBoard.IZQ_C1);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.IZQ_C2))  {
				actions.add(CanibalesBoard.IZQ_C2);
			}
			if (board.peopleCanGetInAndGo(CanibalesBoard.IZQ_MC))  {
				actions.add(CanibalesBoard.IZQ_MC);
			}

			return actions;
		}
	}

	private static class EPResultFunction implements ResultFunction {
		public Object result(Object s, Action a) {
			CanibalesBoard board = (CanibalesBoard) s;
			
			if (CanibalesBoard.DER_M2.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.DER_M2) && 
					board.boatIsThere(CanibalesBoard.DER_M2)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.DER_M2);
				return newBoard;
			}else if (CanibalesBoard.DER_C2.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.DER_C2) && 
					board.boatIsThere(CanibalesBoard.DER_C2)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.DER_C2);
				return newBoard;
			}else if (CanibalesBoard.DER_M1.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.DER_M1) && 
					board.boatIsThere(CanibalesBoard.DER_M1)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.DER_M1);
				return newBoard;
			}else if (CanibalesBoard.DER_C1.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.DER_C1) && 
					board.boatIsThere(CanibalesBoard.DER_C1)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.DER_C1);
				return newBoard;
			}else if (CanibalesBoard.DER_MC.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.DER_MC) && 
					board.boatIsThere(CanibalesBoard.DER_MC)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.DER_MC);
				return newBoard;
			}else if (CanibalesBoard.IZQ_M2.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.IZQ_M2) && 
					board.boatIsThere(CanibalesBoard.IZQ_M2)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.IZQ_M2);
				return newBoard;
			} else if (CanibalesBoard.IZQ_C2.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.IZQ_C2) && 
					board.boatIsThere(CanibalesBoard.IZQ_C2)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.IZQ_C2);
				return newBoard;
			} else if (CanibalesBoard.IZQ_M1.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.IZQ_M1) && 
					board.boatIsThere(CanibalesBoard.IZQ_M1)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.IZQ_M1);
				return newBoard;
			} else if (CanibalesBoard.IZQ_C1.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.IZQ_C1) && 
					board.boatIsThere(CanibalesBoard.IZQ_C1)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.IZQ_C1);
				return newBoard;
			} else if (CanibalesBoard.IZQ_MC.equals(a)
					&& board.peopleCanGetInAndGo(CanibalesBoard.IZQ_MC) && 
					board.boatIsThere(CanibalesBoard.IZQ_MC)) {
				CanibalesBoard newBoard = new CanibalesBoard(board);
				newBoard.move(CanibalesBoard.IZQ_MC);
				return newBoard;
			}

			// The Action is not understood or is a NoOp
			// the result will be the current state.
			return s;
		}
	}
}