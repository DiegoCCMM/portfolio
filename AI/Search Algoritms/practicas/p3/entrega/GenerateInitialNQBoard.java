package aima.gui.demo.search;

import java.util.HashSet;
//import java.util.Iterator;
import java.util.Random;
import java.util.Set;

import aima.core.environment.nqueens.NQueensBoard;
import aima.core.util.datastructure.XYLocation;

/**
 * @author JAB
 * 
 */

public class GenerateInitialNQBoard {
	
	
	public static Set<NQueensBoard> generateSetNqueensBoard(int boardSize, int populationSize){
		Set<NQueensBoard> setGeneratedNQueens = new HashSet<NQueensBoard>();

		while(setGeneratedNQueens.size()< populationSize){
			setGeneratedNQueens.add(generateRandomNqueensBoard(boardSize)); 
		}
		
		return setGeneratedNQueens;
	}


	//Colocar aleatoriamente una reina en una columna
	public static NQueensBoard generateRandomNqueensBoard(int boardSize) {
		
		NQueensBoard board = new NQueensBoard(boardSize);
		
		for(int i = 0; i < boardSize ; i++) {
			board.addQueenAt(new XYLocation(i, new Random().nextInt(boardSize)));
		}
		return board;
	}
	
	//devuelve un estado aleatorio de profundidad depth desde 
	// {0, 1, 2, 3, 4, 5, 6, 7, 8} 
	/*public static EightPuzzleBoard random(int depth) {
		 Set<Object> statesGenerated = new HashSet<Object>();
		 //System.out.println("Generating random state");

		Random r = new Random();
		EightPuzzleBoard board = new EightPuzzleBoard(new int[] { 0, 1, 2, 3,
				4, 5, 6, 7, 8 });
		EightPuzzleBoard newBoard = new EightPuzzleBoard(board);
		
		statesGenerated.add(new EightPuzzleBoard(board));
		int pasos = 1;
		while (pasos<=depth) {

			newBoard = new EightPuzzleBoard(board);
			int th = r.nextInt(4);
			if (th == 0) {
				newBoard.moveGapUp();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapUp();
					 pasos++;
				}
			}
			if (th == 1) {
				newBoard.moveGapDown();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapDown();
					 pasos++;
				}
			}
			if (th == 2) {
				newBoard.moveGapLeft();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapLeft();
					 pasos++;
				}			
			}
			if (th == 3) {
				newBoard.moveGapRight();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapRight();
					 pasos++;
				}				
			}
			
			
		}
		

		return(board);

	}
	
	//devuelve un estado aleatorio de profundidad depth desde 
	//initial State. Ojo, no hay garantia de que no haya una solucion mas corta.
	public static EightPuzzleBoard random(int depth, EightPuzzleBoard initialState){
		 Set<Object> statesGenerated = new HashSet<Object>();

		Random r = new Random();
		EightPuzzleBoard board = new EightPuzzleBoard(initialState);
		EightPuzzleBoard newBoard = new EightPuzzleBoard(board);
		
		statesGenerated.add(new EightPuzzleBoard(board));
		int pasos = 1;
		int limit = 100000; //OJO podria no encontrar solucion de esa profundidad. Ponemos un limite de iteraciones
		int cont = 1;
		while (pasos<=depth && cont < limit) {
			newBoard = new EightPuzzleBoard(board);
			int th = r.nextInt(4);
			if (th == 0) {
				newBoard.moveGapUp();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapUp();
					 pasos++;
				}
			}
			if (th == 1) {
				newBoard.moveGapDown();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapDown();
					 pasos++;
				}
			}
			if (th == 2) {
				newBoard.moveGapLeft();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapLeft();
					 pasos++;
				}			
			}
			if (th == 3) {
				newBoard.moveGapRight();
				if   (!statesGenerated.contains(newBoard)){
					 statesGenerated.add(new EightPuzzleBoard(newBoard));
					 board.moveGapRight();
					 pasos++;
				}				
			}
			
		cont++;	
		}	

		return(board);
	}
*/
}	