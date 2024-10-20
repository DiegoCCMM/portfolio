package aima.gui.demo.search;

import java.util.Iterator;
import java.util.List;
import java.util.Properties;

import aima.core.agent.Action;
import aima.core.environment.eightpuzzle.EightPuzzleBoard;
import aima.core.environment.eightpuzzle.EightPuzzleFunctionFactory;
import aima.core.environment.eightpuzzle.EightPuzzleGoalTest;
import aima.core.environment.eightpuzzle.ManhattanHeuristicFunction;
import aima.core.environment.eightpuzzle.MisplacedTilleHeuristicFunction;
import aima.core.search.framework.GraphSearch;
import aima.core.search.framework.Problem;
import aima.core.search.framework.QueueSearch;
import aima.core.search.framework.ResultFunction;
import aima.core.search.framework.Search;
import aima.core.search.framework.SearchAgent;
import aima.core.search.framework.TreeSearch;
import aima.core.search.informed.AStarSearch;
import aima.core.search.local.SimulatedAnnealingSearch;
import aima.core.search.uninformed.DepthLimitedSearch;
import aima.core.search.uninformed.BreadthFirstSearch;
import aima.core.search.uninformed.IterativeDeepeningSearch;
import aima.core.search.uninformed.UniformCostSearch;
import aima.core.search.uninformed.DepthFirstSearch;


/**
 * @author Ravi Mohan
 * 
 */

public class EightPuzzlePract1 {
	static EightPuzzleBoard boardWithThreeMoveSolution = new EightPuzzleBoard(new int[] { 1, 2, 5, 3, 4, 0, 6, 7, 8 });

	static EightPuzzleBoard random1 = new EightPuzzleBoard(new int[] { 1, 4, 2, 7, 5, 8, 3, 0, 6 });

	static EightPuzzleBoard extreme = new EightPuzzleBoard(new int[] { 0, 8, 7, 6, 5, 4, 3, 2, 1 });

	public static void main(String[] args) {
		/*eightPuzzleDLSDemo();
		eightPuzzleIDLSDemo();
		eightPuzzleGreedyBestFirstDemo();
		eightPuzzleGreedyBestFirstManhattanDemo();
		eightPuzzleAStarDemo();
		eightPuzzleAStarManhattanDemo();
		eightPuzzleSimulatedAnnealingDemo();*/
		
		Search searchDLS9 = new DepthLimitedSearch(9);
		Search searchDLS3 = new DepthLimitedSearch(3);
		Search searchIDS = new IterativeDeepeningSearch();
		Search searchBFSG = new BreadthFirstSearch(new GraphSearch());
		Search searchBFST = new BreadthFirstSearch(new TreeSearch());
		Search searchUCSF = new UniformCostSearch();
		Search searchDFSG = new DepthFirstSearch(new GraphSearch());
		Search searchDFST = new DepthFirstSearch(new TreeSearch());
		Search searchUCST = new UniformCostSearch(new TreeSearch());
		Search searchUCSG = new UniformCostSearch(new GraphSearch());
		
		Problem problem3 = new Problem(boardWithThreeMoveSolution, EightPuzzleFunctionFactory.getActionsFunction(),
				EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
		Problem problem9 = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
				EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
		Problem problem30 = new Problem(extreme, EightPuzzleFunctionFactory.getActionsFunction(),
				EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
		System.out.println("depth|\texpandedNodes|\tqueueSize|\tmaxQueueSize|\ttime");
		
		System.out.println("searchDLS9\t");
		 eitghtPuzzleSearch(searchDLS9, problem3);
		 System.out.println("searchDLS3\t");
		 eitghtPuzzleSearch(searchDLS3, problem3);
		 System.out.println("searchIDS\t");
		 eitghtPuzzleSearch(searchIDS, problem3);
		 System.out.println("searchBFSG\t");
		 eitghtPuzzleSearch(searchBFSG, problem3);
		 System.out.println("searchBFST\t");
		 eitghtPuzzleSearch(searchBFST, problem3);
		 System.out.println("searchUCSF\t");
		 eitghtPuzzleSearch(searchUCSF, problem3); 
		 System.out.println("searchDFSG\t");
		 eitghtPuzzleSearch(searchDFSG, problem3);
		 System.out.println("searchDFST\t");
	     //eitghtPuzzleSearch(searchDFST, problem3);  //no funciona
	     System.out.println("searchUCST\t");
		 eitghtPuzzleSearch(searchUCST, problem3);
		 System.out.println("searchUCSG\t");
		 eitghtPuzzleSearch(searchUCSG, problem3);
		 
		 System.out.println("searchDLS9\t");
		 eitghtPuzzleSearch(searchDLS9, problem9);
		 System.out.println("searchDLS3\t");
		 eitghtPuzzleSearch(searchDLS3, problem9);
		 System.out.println("searchIDS\t");
		 eitghtPuzzleSearch(searchIDS, problem9);
		 System.out.println("searchBFSG\t");
		 eitghtPuzzleSearch(searchBFSG, problem9);
		 System.out.println("searchBFST\t");
		 eitghtPuzzleSearch(searchBFST, problem9);	//no funciona
		 System.out.println("searchUCSF\t");
		 eitghtPuzzleSearch(searchUCSF, problem9); 
		 System.out.println("searchDFSG\t");
		 eitghtPuzzleSearch(searchDFSG, problem9);
		 System.out.println("searchDFST\t");
	     //eitghtPuzzleSearch(searchDFST, problem9); 	//no funciona
	     System.out.println("searchUCST\t");
		 eitghtPuzzleSearch(searchUCST, problem9);
		 System.out.println("searchUCSG\t");
		 eitghtPuzzleSearch(searchUCSG, problem9);
		 
		 System.out.println("searchDLS9\t");
		 eitghtPuzzleSearch(searchDLS9, problem30);
		 System.out.println("searchDLS3\t");
		 eitghtPuzzleSearch(searchDLS3, problem30);
		 System.out.println("searchIDS\t");
		 //eitghtPuzzleSearch(searchIDS, problem30);	//no funciona
		 System.out.println("searchBFSG\t");
		 eitghtPuzzleSearch(searchBFSG, problem30);
		 System.out.println("searchBFST\t");
		 //eitghtPuzzleSearch(searchBFST, problem30);	//no funciona
		 System.out.println("searchUCSF\t");
		 eitghtPuzzleSearch(searchUCSF, problem30); 
		 System.out.println("searchDFSG\t");
		 eitghtPuzzleSearch(searchDFSG, problem30);
		 System.out.println("searchDFST\t");
	     //eitghtPuzzleSearch(searchDFST, problem30); 	//no funciona
	     System.out.println("searchUCST\t");
		 //eitghtPuzzleSearch(searchUCST, problem30);   //no funciona
	     System.out.println("searchUCSG\t");
		 eitghtPuzzleSearch(searchUCSG, problem30);
		
		
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
			
			//executeActions(agent.getActions(), problem);
			

			long duration = (endTime - startTime);
			visualization(agent, duration);

		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	private static void eightPuzzleDLSDemo() {
		System.out.println("\nEightPuzzleDemo recursive DLS (9) -->");
		try {
			Problem problem = new Problem(boardWithThreeMoveSolution, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new DepthLimitedSearch(9);
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

	private static void eightPuzzleIDLSDemo() {
		System.out.println("\nEightPuzzleDemo Iterative DLS -->");
		try {
			Problem problem = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new IterativeDeepeningSearch();
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

	/*
	 * private static void eightPuzzleGreedyBestFirstDemo() { System.out.
	 * println("\nEightPuzzleDemo Greedy Best First Search (MisplacedTileHeursitic)-->"
	 * ); try { Problem problem = new Problem(boardWithThreeMoveSolution,
	 * EightPuzzleFunctionFactory.getActionsFunction(),
	 * EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
	 * Search search = new GreedyBestFirstSearch(new GraphSearch(), new
	 * MisplacedTilleHeuristicFunction()); SearchAgent agent = new
	 * SearchAgent(problem, search); printActions(agent.getActions());
	 * printInstrumentation(agent.getInstrumentation()); } catch (Exception e) {
	 * e.printStackTrace(); }
	 * 
	 * }
	 * 
	 * private static void eightPuzzleGreedyBestFirstManhattanDemo() { System.out.
	 * println("\nEightPuzzleDemo Greedy Best First Search (ManhattanHeursitic)-->"
	 * ); try { Problem problem = new Problem(boardWithThreeMoveSolution,
	 * EightPuzzleFunctionFactory.getActionsFunction(),
	 * EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
	 * Search search = new GreedyBestFirstSearch(new GraphSearch(), new
	 * ManhattanHeuristicFunction()); SearchAgent agent = new SearchAgent(problem,
	 * search); printActions(agent.getActions());
	 * printInstrumentation(agent.getInstrumentation()); } catch (Exception e) {
	 * e.printStackTrace(); }
	 * 
	 * }
	 */

	private static void eightPuzzleAStarDemo() {
		System.out.println("\nEightPuzzleDemo AStar Search (MisplacedTileHeursitic)-->");
		try {
			Problem problem = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new AStarSearch(new GraphSearch(), new MisplacedTilleHeuristicFunction());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

	private static void eightPuzzleSimulatedAnnealingDemo() {
		System.out.println("\nEightPuzzleDemo Simulated Annealing  Search -->");
		try {
			Problem problem = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			SimulatedAnnealingSearch search = new SimulatedAnnealingSearch(new ManhattanHeuristicFunction());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			System.out.println("Search Outcome=" + search.getOutcome());
			System.out.println("Final State=\n" + search.getLastSearchState());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	private static void eightPuzzleAStarManhattanDemo() {
		System.out.println("\nEightPuzzleDemo AStar Search (ManhattanHeursitic)-->");
		try {
			Problem problem = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new AStarSearch(new GraphSearch(), new ManhattanHeuristicFunction());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
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