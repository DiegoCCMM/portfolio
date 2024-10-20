package aima.gui.demo.search;

import java.util.Iterator;
import java.util.List;
import java.util.Properties;

import aima.core.agent.Action;
import aima.core.environment.Fichas.FichasBoard;
import aima.core.environment.Fichas.FichasFunctionFactory;
import aima.core.environment.Fichas.FichasGoalTest;
import aima.core.search.framework.GraphSearch;
import aima.core.search.framework.Problem;
import aima.core.search.framework.QueueSearch;
import aima.core.search.framework.ResultFunction;
import aima.core.search.framework.Search;
import aima.core.search.framework.SearchAgent;
import aima.core.search.framework.TreeSearch;
import aima.core.search.informed.AStarSearch;
import aima.core.search.informed.GreedyBestFirstSearch;
import aima.core.search.local.SimulatedAnnealingSearch;
import aima.core.search.uninformed.DepthLimitedSearch;
import aima.core.search.uninformed.IterativeDeepeningSearch;
import aima.core.search.uninformed.UniformCostSearch;
import aima.core.search.uninformed.BreadthFirstSearch;
import aima.core.search.uninformed.DepthFirstSearch;


/**
 * @author Ravi Mohan
 * 
 */

public class FichasPract1 {
	static FichasBoard initial = new FichasBoard(new int[] { 1, 1, 1, 0, 2, 2, 2});

	public static void main(String[] args) {
		/*eightPuzzleDLSDemo();
		eightPuzzleIDLSDemo();
		eightPuzzleGreedyBestFirstDemo();
		eightPuzzleGreedyBestFirstManhattanDemo();
		eightPuzzleAStarDemo();
		eightPuzzleAStarManhattanDemo();
		eightPuzzleSimulatedAnnealingDemo();*/
		
		Search searchDLS11 = new DepthLimitedSearch(11);
		Search searchIDS = new IterativeDeepeningSearch();
		//Search searchBFSG = new GreedyBestFirstSearch(new GraphSearch(), new MisplacedTilleHeuristicFunction());
		//Search searchBFST = new GreedyBestFirstSearch(new TreeSearch(), new MisplacedTilleHeuristicFunction());
		Search searchUCSF = new UniformCostSearch();
		Search searchDFSG = new DepthFirstSearch(new GraphSearch());
		Search searchDFST = new DepthFirstSearch(new TreeSearch());
		Search searchUCST = new UniformCostSearch(new TreeSearch());
		
		Search BFS = new BreadthFirstSearch(new GraphSearch());
		
		Problem problem = new Problem(initial, FichasFunctionFactory
					.getActionsFunction(), FichasFunctionFactory
					.getResultFunction(), new FichasGoalTest());
		
		
		
		System.out.println("searchBFS\t");
		 eitghtPuzzleSearch(BFS, problem);
		
		
		System.out.println("searchDLS11\t");
		 eitghtPuzzleSearch(searchDLS11, problem);
		 
		 System.out.println("searchIDS\t");
		 eitghtPuzzleSearch(searchIDS, problem);
		 
//		 System.out.println("searchUCSF\t");
//		 eitghtPuzzleSearch(searchUCSF, problem); 
//		 System.out.println("searchDFSG\t");
//		 eitghtPuzzleSearch(searchDFSG, problem);
//		 System.out.println("searchDFST\t");
//	     eitghtPuzzleSearch(searchDFST, problem);  //no funciona
//	     System.out.println("searchUCST\t");
//		 eitghtPuzzleSearch(searchUCST, problem);
		
		
	}

	public static void executeActions(List<Action> actions, Problem problem) {
		Object initialState = problem.getInitialState();
		ResultFunction resultFunction = problem.getResultFunction();
		Object state = initialState;
		System.out.println("INITIAL STATE");
		System.out.println(state);
		for (Action action : actions) {
			System.out.println(action.toString());
			state = resultFunction.result(state, action);
			System.out.println(state);
			System.out.println("- - -");
		}
	}

	private static void visualization(SearchAgent agent, long time) {
		int depth, expandedNodes, queueSize, maxQueueSize;

		String pathcostM = agent.getInstrumentation().getProperty("pathCost");
		if (pathcostM != null)
			depth = (int) Float.parseFloat(pathcostM);
		else
			depth = 0;
		if (agent.getInstrumentation().getProperty("nodesExpanded") == null)
			expandedNodes = 0;
		else
			expandedNodes = (int) Float.parseFloat(agent.getInstrumentation().getProperty("nodesExpanded"));
		if (agent.getInstrumentation().getProperty("queueSize") == null)
			queueSize = 0;
		else
			queueSize = (int) Float.parseFloat(agent.getInstrumentation().getProperty("queueSize"));
		if (agent.getInstrumentation().getProperty("maxQueueSize") == null)
			maxQueueSize = 0;
		else
			maxQueueSize = (int) Float.parseFloat(agent.getInstrumentation().getProperty("maxQueueSize"));

		System.out.println(depth + "|\t" + expandedNodes + "|\t" + queueSize + "|\t" + maxQueueSize + 
				"|\t" + time);
	}

	private static void eitghtPuzzleSearch(Search search, Problem problem) {
		try {
			long startTime = System.nanoTime();
			SearchAgent agent = new SearchAgent(problem, search);
			long endTime = System.nanoTime();
			

			long duration = (endTime - startTime);
			System.out.println("depth|\texpandedNodes|\tqueueSize|\tmaxQueueSize|\ttime");
			visualization(agent, duration);
			
			executeActions(agent.getActions(), problem);

		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	private static void printInstrumentation(Properties properties) {
		Iterator<Object> keys = properties.keySet().iterator();
		while (keys.hasNext()) {
			String key = (String) keys.next();
			String property = properties.getProperty(key);
			System.out.println(key + " : " + property);
		}

	}

	private static void printActions(List<Action> actions) {
		for (int i = 0; i < actions.size(); i++) {
			String action = actions.get(i).toString();
			System.out.println(action);
		}
	}

}