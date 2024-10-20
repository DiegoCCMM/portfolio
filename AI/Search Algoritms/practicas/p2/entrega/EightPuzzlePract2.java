package aima.gui.demo.search;

import java.util.Iterator;
import java.util.List;
import java.util.Properties;

import aima.core.agent.Action;
import aima.core.environment.eightpuzzle.EightPuzzleBoard;
import aima.core.environment.eightpuzzle.EightPuzzleFunctionFactory;
import aima.core.environment.eightpuzzle.EightPuzzleGoalTest;
import aima.core.environment.eightpuzzle.ManhattanHeuristicFunction2;
import aima.core.environment.eightpuzzle.MisplacedTilleHeuristicFunction2;
import aima.core.search.framework.GraphSearch;
import aima.core.search.framework.Problem;
import aima.core.search.framework.Search;
import aima.core.search.framework.SearchAgent;
import aima.core.search.informed.AStarSearch;
import aima.core.search.informed.GreedyBestFirstSearch;
import aima.core.search.local.SimulatedAnnealingSearch;
import aima.core.search.uninformed.BreadthFirstSearch;
import aima.core.search.uninformed.DepthLimitedSearch;
import aima.core.search.uninformed.IterativeDeepeningSearch;
import aima.core.util.math.Biseccion;

/**
 * @author Ravi Mohan
 * 
 */

public class EightPuzzlePract2 {
	static EightPuzzleBoard boardWithThreeMoveSolution = new EightPuzzleBoard(new int[] { 1, 2, 5, 3, 4, 0, 6, 7, 8 });;

	static EightPuzzleBoard random1 = new EightPuzzleBoard(new int[] { 1, 4, 2, 7, 5, 8, 3, 0, 6 });

	static EightPuzzleBoard extreme = new EightPuzzleBoard(new int[] { 0, 8, 7, 6, 5, 4, 3, 2, 1 });

	public static void main(String[] args) {

		double matrizBFS[][] = new double[2][30];
		double matrizIDS[][] = new double[2][30];
		double matrizMT[][] = new double[2][30];
		double matrizMH[][] = new double[2][30];

		// eightPuzzleDLSDemo();
		// eightPuzzleIDLSDemo();
		// eightPuzzleGreedyBestFirstDemo();
		// eightPuzzleGreedyBestFirstManhattanDemo();
		cabeceraExperimentos();
		matrizBFS = experimentos("BFS", 25, 100);
		matrizIDS = experimentos("IDS", 25, 100);
		matrizMT = experimentos("ASMT", 25, 100);
		matrizMH = experimentos("ASMH", 25, 100);
		for (int i = 2; i < 25; i++) {
			System.out.format("|| %3d||", i);
			System.out.format(" %6d  |", (int) matrizBFS[0][i]);
			if ((int) matrizIDS[0][i] != 0) {
				System.out.format(" %6d  |", (int) matrizIDS[0][i]);
			}else {
				System.out.format(" ---     |");
			}
			System.out.format(" %6d  |", (int) matrizMT[0][i]);
			System.out.format(" %6d  ||", (int) matrizMH[0][i]);
			System.out.format(" %6f  |", matrizBFS[1][i]);
			if ((int) matrizIDS[1][i] != 0) {
				System.out.format(" %6f  |", matrizIDS[1][i]);
			}else {
				System.out.format(" ---       |");
			}
			System.out.format(" %6f  |", matrizMT[1][i]);
			System.out.format(" %6f  ||", matrizMH[1][i]);
			System.out.println();
		}
		// eightPuzzleAStarManhattanDemo();
		// eightPuzzleSimulatedAnnealingDemo();
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

	private static void eightPuzzleGreedyBestFirstDemo() {
		System.out.println("\nEightPuzzleDemo Greedy Best First Search (MisplacedTileHeursitic)-->");
		try {
			Problem problem = new Problem(boardWithThreeMoveSolution, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new GreedyBestFirstSearch(new GraphSearch(), new MisplacedTilleHeuristicFunction2());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

	private static void eightPuzzleGreedyBestFirstManhattanDemo() {
		System.out.println("\nEightPuzzleDemo Greedy Best First Search (ManhattanHeursitic)-->");
		try {
			Problem problem = new Problem(boardWithThreeMoveSolution, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new GreedyBestFirstSearch(new GraphSearch(), new ManhattanHeuristicFunction2());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

	private static void eightPuzzleAStarDemo() {
		System.out.println("\nEightPuzzleDemo AStar Search (MisplacedTileHeursitic)-->");
		try {
			Problem problem = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			Search search = new AStarSearch(new GraphSearch(), new MisplacedTilleHeuristicFunction2());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}

	}

	private static void cabeceraExperimentos() {
		for (int i = 1; i < 91; i++)
			System.out.print("-");
		System.out.println();
		System.out.format("||    || %20s%17s || %19s%19s||\n", "Nodos Generados", " ", "b*", " ");
		for (int i = 1; i < 91; i++)
			System.out.print("-");
		System.out.println();
		System.out.format("|| %3s||", "d");
		System.out.format(" %6s  |", "BFS");
		System.out.format(" %6s  |", "IDS");
		System.out.format(" %6s  |", "A*h(1)");
		System.out.format(" %6s  ||", "A*h(2)");
		System.out.format(" %6s  |", "BFS");
		System.out.format(" %6s  |", "IDS");
		System.out.format(" %6s  |", "A*h(1)");
		System.out.format(" %6s  ||\n", "A*h(2)");
		for (int i = 1; i < 91; i++)
			System.out.print("-");
		System.out.println();
		for (int i = 1; i < 91; i++)
			System.out.print("-");
		System.out.println();
	}


	private static double[][] experimentos(String searchName, int profundidadMAX, int numExp) {
		Search search[] = { new BreadthFirstSearch(), new IterativeDeepeningSearch(),
				new AStarSearch(new GraphSearch(), new MisplacedTilleHeuristicFunction2()),
				new AStarSearch(new GraphSearch(), new ManhattanHeuristicFunction2()) };
		double resultadosNG[] = new double[profundidadMAX];
		double resultadosB[] = new double[profundidadMAX];
		double[][] resultados = new double[2][profundidadMAX];

		SearchAgent SA = null;
		
		double mediaNodosGenerados = 0;
		double mediaNodosExpandidos = 0;
		int depth = 0;

		if (searchName.compareTo("IDS") == 0 && profundidadMAX > 10) {
			profundidadMAX = 11;
		}
		for (int nivel = 2; nivel < profundidadMAX; nivel++) {
			for (int experimento = 1; experimento <= numExp; experimento=experimento) {

				EightPuzzleBoard initialState = 
						new EightPuzzleBoard(GenerateInitialEightPuzzleBoard.randomIni());

				EightPuzzleBoard finalState = GenerateInitialEightPuzzleBoard.random(nivel,
						new EightPuzzleBoard(initialState));

				EightPuzzleGoalTest finalGState = new EightPuzzleGoalTest();
				EightPuzzleGoalTest.setGoalState(new EightPuzzleBoard(finalState));
				
				SA = agenteDeBusqueda(searchName, initialState, finalGState, search);
				
				if (SA.getInstrumentation().getProperty("pathCost") != null) {
					depth = (int) Float.parseFloat(SA.getInstrumentation().getProperty("pathCost"));
				} else {
					depth = 0;
				}
				
				if (depth == nivel) {
					mediaNodosGenerados += (double) Float
							.parseFloat(SA.getInstrumentation().getProperty("nodesGenerated"));
					mediaNodosExpandidos += (double) Float
							.parseFloat(SA.getInstrumentation().getProperty("nodesExpanded"));
					experimento++;

				}
					   
					 
				

			}

			resultadosNG[nivel] = Math.round(mediaNodosGenerados / numExp);
			Biseccion bf = new Biseccion();
			bf.setDepth(nivel);
			bf.setGeneratedNodes((int) resultadosNG[nivel]);
			resultadosB[nivel] = bf.metodoDeBiseccion(1.0000000000001, 4, 1E-12);

			resultados[0][nivel] = resultadosNG[nivel];
			resultados[1][nivel] = resultadosB[nivel];
			mediaNodosGenerados = 0;
			mediaNodosExpandidos = 0;
		}

		return resultados;

	}
	
	private static SearchAgent agenteDeBusqueda(String searchName, EightPuzzleBoard initialState,
			EightPuzzleGoalTest finalGState, Search[] search) {
		SearchAgent SA = null;
		if (searchName.compareTo("BFS") == 0) {
			SA = eitghtPuzzleSearch(search[0], initialState, finalGState);
		} else if (searchName.compareTo("IDS") == 0) {
			SA = eitghtPuzzleSearch(search[1], initialState, finalGState);
		} else if (searchName.compareTo("ASMT") == 0) {
			SA = eitghtPuzzleSearch(search[2], initialState, finalGState);
		} else if (searchName.compareTo("ASMH") == 0) {
			SA = eitghtPuzzleSearch(search[3], initialState, finalGState);
		}
		return SA;
	}

	private static SearchAgent eitghtPuzzleSearch(Search search, EightPuzzleBoard board,
			EightPuzzleGoalTest goalBoard) {

		try {

			Problem problem = new Problem(board, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), goalBoard);
			SearchAgent agent = new SearchAgent(problem, search);

			return (agent);

		} catch (Exception e) {
			e.printStackTrace();
		}
		return null;

	}

	private static void eightPuzzleSimulatedAnnealingDemo() {
		System.out.println("\nEightPuzzleDemo Simulated Annealing  Search -->");
		try {
			Problem problem = new Problem(random1, EightPuzzleFunctionFactory.getActionsFunction(),
					EightPuzzleFunctionFactory.getResultFunction(), new EightPuzzleGoalTest());
			SimulatedAnnealingSearch search = new SimulatedAnnealingSearch(new ManhattanHeuristicFunction2());
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
			Search search = new AStarSearch(new GraphSearch(), new ManhattanHeuristicFunction2());
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