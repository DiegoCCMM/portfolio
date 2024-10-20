package aima.core.environment.Fichas;

import java.util.LinkedHashSet;
import java.util.Set;

import aima.core.agent.Action;
import aima.core.environment.Canibales.CanibalesBoard;
import aima.core.search.framework.ActionsFunction;
import aima.core.search.framework.ResultFunction;

/**
 * @author Diego Caballé 738712
 */
public class FichasFunctionFactory {
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
			FichasBoard board = (FichasBoard) state;

			Set<Action> actions = new LinkedHashSet<Action>();

			if (board.canMove(FichasBoard.R2)) {
				actions.add(FichasBoard.R2);
			}
			if (board.canMove(FichasBoard.R1))  {
				actions.add(FichasBoard.R1);
			}
			if (board.canMove(FichasBoard.L1))  {
				actions.add(FichasBoard.L1);
			}
			if (board.canMove(FichasBoard.L2))  {
				actions.add(FichasBoard.L2);
			}
			if (board.canMove(FichasBoard.L3))  {
				actions.add(FichasBoard.L3);
			}
			
			if (board.canMove(FichasBoard.R3))  {
				actions.add(FichasBoard.R3);
			}

			return actions;
		}
	}

	private static class EPResultFunction implements ResultFunction {
		public Object result(Object s, Action a) {
			FichasBoard board = (FichasBoard) s;
			
			if (FichasBoard.R3.equals(a) && board.canMove(FichasBoard.R3)) {
				FichasBoard newBoard = new FichasBoard(board);
				newBoard.move(FichasBoard.R3);
				return newBoard;
			}else if (FichasBoard.R2.equals(a) && board.canMove(FichasBoard.R2)) {
				FichasBoard newBoard = new FichasBoard(board);
				newBoard.move(FichasBoard.R2);
				return newBoard;
			}else if (FichasBoard.R1.equals(a) && board.canMove(FichasBoard.R1)) {
				FichasBoard newBoard = new FichasBoard(board);
				newBoard.move(FichasBoard.R1);
				return newBoard;
			}else if (FichasBoard.L3.equals(a) && board.canMove(FichasBoard.L3)) {
				FichasBoard newBoard = new FichasBoard(board);
				newBoard.move(FichasBoard.L3);
				return newBoard;
			}else if (FichasBoard.L2.equals(a) && board.canMove(FichasBoard.L2)) {
				FichasBoard newBoard = new FichasBoard(board);
				newBoard.move(FichasBoard.L2);
				return newBoard;
			}else if (FichasBoard.L1.equals(a) && board.canMove(FichasBoard.L1)) {
				FichasBoard newBoard = new FichasBoard(board);
				newBoard.move(FichasBoard.L1);
				return newBoard;
			}

			// The Action is not understood or is a NoOp
			// the result will be the current state.
			return s;
		}
	}
}