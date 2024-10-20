package aima.gui.demo.search;

import java.math.BigDecimal;
import java.util.HashSet;
import java.util.Iterator;
import java.util.List;
import java.util.Properties;
import java.util.Set;

import aima.core.agent.Action;
import aima.core.environment.nqueens.AttackingPairsHeuristic;
import aima.core.environment.nqueens.NQueensBoard;
import aima.core.environment.nqueens.NQueensFitnessFunction;
import aima.core.environment.nqueens.NQueensFunctionFactory;
import aima.core.environment.nqueens.NQueensGoalTest;
import aima.core.search.framework.GraphSearch;
import aima.core.search.framework.Problem;
import aima.core.search.framework.Search;
import aima.core.search.framework.SearchAgent;
import aima.core.search.framework.TreeSearch;
import aima.core.search.local.GeneticAlgorithm;
import aima.core.search.local.HillClimbingSearch;
import aima.core.search.local.Individual;
import aima.core.search.local.Scheduler;
import aima.core.search.local.SimulatedAnnealingSearch;
import aima.core.search.local.HillClimbingSearch.SearchOutcome;
import aima.core.search.uninformed.BreadthFirstSearch;
import aima.core.search.uninformed.DepthFirstSearch;
import aima.core.search.uninformed.DepthLimitedSearch;
import aima.core.search.uninformed.IterativeDeepeningSearch;

/**
 * @author Ravi Mohan
 * 
 */

public class NQueensLocal {

	private static final int _boardSize = 8;

	public static void main(String[] args) {

		newNQueensDemo();
	}

	private static void newNQueensDemo() {

		/*
		 * nQueensWithDepthFirstSearch(); nQueensWithBreadthFirstSearch();
		 * nQueensWithRecursiveDLS(); nQueensWithIterativeDeepeningSearch();
		 */
		/*
		 * nQueensSimulatedAnnealingSearch(); nQueensHillClimbingSearch();
		 * nQueensGeneticAlgorithmSearch();
		 */
		System.out.println("Apartado 1 \n");
		nQueensHillClimbingSearch_Statistics(10000);
		System.out.println("Apartado 2 \n");
		nQueensRandomRestartHillClimbing();
		System.out.println("Apartado 3 \n");
		nQueensSimulatedAnnealing_Statistics(1000);
		System.out.println("Apartado 4 \n");
		nQueensHillSimulatedAnnealingRestart();
		System.out.println("Apartado 5 \n");
		nQueensGeneticAlgorithmSearch();
	}

	/*
	 * private static void nQueensWithRecursiveDLS() {
	 * System.out.println("\nNQueensDemo recursive DLS -->"); try { Problem problem
	 * = new Problem(new NQueensBoard(_boardSize),
	 * NQueensFunctionFactory.getIActionsFunction(),
	 * NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest()); Search
	 * search = new DepthLimitedSearch(_boardSize); SearchAgent agent = new
	 * SearchAgent(problem, search); printActions(agent.getActions());
	 * printInstrumentation(agent.getInstrumentation()); } catch (Exception e) {
	 * e.printStackTrace(); }
	 * 
	 * }
	 */
	private static void nQueensHillClimbingSearch_Statistics(int numExperiments) {

		double mediaNodosGeneradosError = 0, mediaNodosGeneradosAcierto = 0, aciertos = 0, errores = 0;

		Set<NQueensBoard> setGeneratedNQueens = GenerateInitialNQBoard.generateSetNqueensBoard(8, numExperiments);
		Iterator<NQueensBoard> iterador = setGeneratedNQueens.iterator();

		while (iterador.hasNext()) {
			Problem problem = new Problem(iterador.next(), NQueensFunctionFactory.getCActionsFunction(),
					NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest());
			HillClimbingSearch search = new HillClimbingSearch(new AttackingPairsHeuristic());
			try {
				SearchAgent agent = new SearchAgent(problem, search);

				// esta busqueda ha sido exitosa?
				if (search.getOutcome().toString().contentEquals("FAILURE")) {
					errores++;
					mediaNodosGeneradosError += agent.getActions().size();
				} else {// aciertos
					aciertos++;
					mediaNodosGeneradosAcierto += agent.getActions().size();

				}

			} catch (Exception e) {
				e.printStackTrace();
			}
			iterador.remove();
		}
		mediaNodosGeneradosError = mediaNodosGeneradosError / numExperiments;
		mediaNodosGeneradosAcierto = mediaNodosGeneradosAcierto / numExperiments;
		aciertos = (aciertos / numExperiments) * 100;
		errores = (errores / numExperiments) * 100;

		System.out.println("Fallos: " + errores);
		System.out.println("Coste medio fallos: " + mediaNodosGeneradosError);
		System.out.println("Exitos: " + aciertos);
		System.out.println("Coste medio Exitos: " + mediaNodosGeneradosAcierto);

	}

	private static void nQueensRandomRestartHillClimbing() {
		boolean encontrado = false;
		double mediaNodosGeneradosError = 0, mediaNodosGeneradosAcierto = 0, aciertos = 0, errores = 0, numIntentos = 0;

		while (!encontrado) {
			numIntentos++;
			Problem problem = new Problem(GenerateInitialNQBoard.generateRandomNqueensBoard(8),
					NQueensFunctionFactory.getCActionsFunction(), NQueensFunctionFactory.getResultFunction(),
					new NQueensGoalTest());
			HillClimbingSearch search = new HillClimbingSearch(new AttackingPairsHeuristic());
			try {
				SearchAgent agent = new SearchAgent(problem, search);

				// esta busqueda ha sido exitosa?
				if (search.getOutcome().toString().contentEquals("FAILURE")) {
					errores++;
					mediaNodosGeneradosError += agent.getActions().size();
				} else {// aciertos
					encontrado = true;
					aciertos++;
					mediaNodosGeneradosAcierto += agent.getActions().size();

					System.out.println("Search Outcome=SOLUTION_FOUND");
					System.out.println("Final state= \n" + search.getLastSearchState());

				}
			} catch (Exception e) {
				e.printStackTrace();
			}

		}

		System.out.println("Intentos: " + numIntentos);
		System.out.println("Fallos: " + errores);
		System.out.println("Coste medio fallos: " + mediaNodosGeneradosError);
		System.out.println("Exitos: " + aciertos);
		System.out.println("Coste medio Exitos: " + mediaNodosGeneradosAcierto);

	}

	private static void nQueensSimulatedAnnealing_Statistics(int numExperiments) {

		double mediaNodosGeneradosError = 0, mediaNodosGeneradosAcierto = 0, aciertos = 0, errores = 0;
		int k = 500;
		double lambda = 0.1;
		int limit = 1800;
		Set<NQueensBoard> setGeneratedNQueens = GenerateInitialNQBoard.generateSetNqueensBoard(8, numExperiments);
		Iterator<NQueensBoard> iterador = setGeneratedNQueens.iterator();
		Scheduler sched = new Scheduler(k, lambda, limit);
		System.out.println("Parámetros Scheduler: Scheduler ( " + k + ", " + lambda + ", " + limit + " );");
		while (iterador.hasNext()) {

			Problem problem = new Problem(iterador.next(), NQueensFunctionFactory.getCActionsFunction(),
					NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest());
			SimulatedAnnealingSearch search = new SimulatedAnnealingSearch(new AttackingPairsHeuristic(), sched);

			try {
				SearchAgent agent = new SearchAgent(problem, search);

				// esta busqueda ha sido exitosa?
				if (search.getOutcome().toString().contentEquals("FAILURE")) {
					errores++;
					mediaNodosGeneradosError += agent.getActions().size();
				} else {// aciertos
					aciertos++;
					mediaNodosGeneradosAcierto += agent.getActions().size();

				}

			} catch (Exception e) {
				e.printStackTrace();
			}
			iterador.remove();
		}
		mediaNodosGeneradosError = mediaNodosGeneradosError / numExperiments;
		mediaNodosGeneradosAcierto = mediaNodosGeneradosAcierto / numExperiments;
		aciertos = (aciertos / numExperiments) * 100;
		errores = (errores / numExperiments) * 100;

		System.out.println("Fallos: " + errores);
		System.out.println("Coste medio fallos: " + mediaNodosGeneradosError);
		System.out.println("Exitos: " + aciertos);
		System.out.println("Coste medio Exitos: " + mediaNodosGeneradosAcierto);
	}

	private static void nQueensHillSimulatedAnnealingRestart() {
		boolean encontrado = false;
		double mediaNodosGeneradosError = 0, mediaNodosGeneradosAcierto = 0, aciertos = 0, errores = 0, numIntentos = 0;

		int k = 500;
		double lambda = 0.1;
		int limit = 1800;

		Scheduler sched = new Scheduler(k, lambda, limit);

		while (!encontrado) {
			numIntentos++;
			Problem problem = new Problem(GenerateInitialNQBoard.generateRandomNqueensBoard(8),
					NQueensFunctionFactory.getCActionsFunction(), NQueensFunctionFactory.getResultFunction(),
					new NQueensGoalTest());
			SimulatedAnnealingSearch search = new SimulatedAnnealingSearch(new AttackingPairsHeuristic(), sched);
			try {
				SearchAgent agent = new SearchAgent(problem, search);

				// esta busqueda ha sido exitosa?
				if (search.getOutcome().toString().contentEquals("FAILURE")) {
					errores++;
					mediaNodosGeneradosError += agent.getActions().size();
				} else {// aciertos
					encontrado = true;
					aciertos++;
					mediaNodosGeneradosAcierto += agent.getActions().size();

					System.out.println("Search Outcome=SOLUTION_FOUND");
					System.out.println("Final state= \n" + search.getLastSearchState());

				}
			} catch (Exception e) {
				e.printStackTrace();
			}

		}
	
		System.out.println("Intentos: " + numIntentos);
		System.out.println("Fallos: " + errores);
		System.out.println("Coste medio fallos: " + mediaNodosGeneradosError);
		System.out.println("Exitos: " + aciertos);
		System.out.println("Coste medio Exitos: " + mediaNodosGeneradosAcierto);
	}

	/*
	 * private static void nQueensWithBreadthFirstSearch() { try {
	 * System.out.println("\nNQueensDemo BFS -->"); Problem problem = new
	 * Problem(new NQueensBoard(_boardSize),
	 * NQueensFunctionFactory.getIActionsFunction(),
	 * NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest()); Search
	 * search = new BreadthFirstSearch(new TreeSearch()); SearchAgent agent2 = new
	 * SearchAgent(problem, search); printActions(agent2.getActions());
	 * printInstrumentation(agent2.getInstrumentation()); } catch (Exception e1) {
	 * 
	 * e1.printStackTrace(); } }
	 */
	/*
	 * private static void nQueensWithDepthFirstSearch() {
	 * System.out.println("\nNQueensDemo DFS -->"); try { Problem problem = new
	 * Problem(new NQueensBoard(_boardSize),
	 * NQueensFunctionFactory.getIActionsFunction(),
	 * NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest()); Search
	 * search = new DepthFirstSearch(new GraphSearch()); SearchAgent agent = new
	 * SearchAgent(problem, search); printActions(agent.getActions());
	 * printInstrumentation(agent.getInstrumentation()); } catch (Exception e) {
	 * e.printStackTrace(); } }
	 * 
	 * private static void nQueensWithIterativeDeepeningSearch() {
	 * System.out.println("\nNQueensDemo Iterative DS  -->"); try { Problem problem
	 * = new Problem(new NQueensBoard(_boardSize),
	 * NQueensFunctionFactory.getIActionsFunction(),
	 * NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest()); Search
	 * search = new IterativeDeepeningSearch(); SearchAgent agent = new
	 * SearchAgent(problem, search);
	 * 
	 * System.out.println(); printActions(agent.getActions());
	 * printInstrumentation(agent.getInstrumentation()); } catch (Exception e) {
	 * e.printStackTrace(); } }
	 */
	/*private static void nQueensSimulatedAnnealingSearch() {
		System.out.println("\nNQueensDemo Simulated Annealing  -->");
		try {
			Problem problem = new Problem(new NQueensBoard(_boardSize), NQueensFunctionFactory.getIActionsFunction(),
					NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest());
			SimulatedAnnealingSearch search = new SimulatedAnnealingSearch(new AttackingPairsHeuristic());
			SearchAgent agent = new SearchAgent(problem, search);

			System.out.println();
			printActions(agent.getActions());
			System.out.println("Search Outcome=" + search.getOutcome());
			System.out.println("Final State=\n" + search.getLastSearchState());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	private static void nQueensHillClimbingSearch() {
		System.out.println("\nNQueensDemo HillClimbing  -->");
		try {
			Problem problem = new Problem(new NQueensBoard(_boardSize), NQueensFunctionFactory.getIActionsFunction(),
					NQueensFunctionFactory.getResultFunction(), new NQueensGoalTest());
			HillClimbingSearch search = new HillClimbingSearch(new AttackingPairsHeuristic());
			SearchAgent agent = new SearchAgent(problem, search);

			System.out.println();
			printActions(agent.getActions());
			System.out.println("Search Outcome=" + search.getOutcome());
			System.out.println("Final State=\n" + search.getLastSearchState());
			printInstrumentation(agent.getInstrumentation());
		} catch (Exception e) {
			e.printStackTrace();
		}
	}*/
	
	
	
	public static void nQueensGeneticAlgorithmSearch() {
		System.out.println("\nNQueensDemo GeneticAlgorithm  -->");
		int poblacion = 64;
		double mutacion = 0.08;
		System.out.println("Parámetros iniciales: " + poblacion + ", " + mutacion);
		try {
			NQueensFitnessFunction fitnessFunction = new NQueensFitnessFunction();
			// Generate an initial population
			Set<Individual<Integer>> population = new HashSet<Individual<Integer>>();
			for (int i = 0; i < poblacion; i++) {
				population.add(fitnessFunction.generateRandomIndividual(_boardSize));
			}

			GeneticAlgorithm<Integer> ga = new GeneticAlgorithm<Integer>(_boardSize,
					fitnessFunction.getFiniteAlphabetForBoardOfSize(_boardSize), mutacion);

			// Run for a set amount of time
			Individual<Integer> bestIndividual = ga.geneticAlgorithm(population, fitnessFunction, fitnessFunction,
					1000L);

			System.out.println(
					"Max Time (1 second) Best Individual=\n" + fitnessFunction.getBoardForIndividual(bestIndividual));
			System.out.println("Board Size      = " + _boardSize);
			System.out.println("# Board Layouts = " + (new BigDecimal(_boardSize)).pow(_boardSize));
			System.out.println("Fitness         = " + fitnessFunction.getValue(bestIndividual));
			System.out.println("Is Goal         = " + fitnessFunction.isGoalState(bestIndividual));
			System.out.println("Population Size = " + ga.getPopulationSize());
			System.out.println("Itertions       = " + ga.getIterations());
			System.out.println("Took            = " + ga.getTimeInMilliseconds() + "ms.");

			/*// Run till goal is achieved
			bestIndividual = ga.geneticAlgorithm(population, fitnessFunction, fitnessFunction, 0L);

			System.out.println("");
			System.out.println("Goal Test Best Individual=\n" + fitnessFunction.getBoardForIndividual(bestIndividual));
			System.out.println("Board Size      = " + _boardSize);
			System.out.println("# Board Layouts = " + (new BigDecimal(_boardSize)).pow(_boardSize));
			System.out.println("Fitness         = " + fitnessFunction.getValue(bestIndividual));
			System.out.println("Is Goal         = " + fitnessFunction.isGoalState(bestIndividual));
			System.out.println("Population Size = " + ga.getPopulationSize());
			System.out.println("Itertions       = " + ga.getIterations());
			System.out.println("Took            = " + ga.getTimeInMilliseconds() + "ms.");*/

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