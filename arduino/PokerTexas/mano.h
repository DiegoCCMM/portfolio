#ifndef MANO_H
#define MANO_H


enum PALO {DIAMANTES, TREBOLES, CORAZONES, PICAS};	//pares rojos, impares negros
enum FIGURA {AS = 1, JOTA = 10, REINA = 11, REY = 12};	//pares rojos, impares negros

class mano {

private:
	//las cartas se entienden como un palo y una figura.
	struct carta {
		PALO palo;
		FIGURA figura;
	};	
	//la mano del jugador son 5 cartas.
    carta cartas[5];

public:
	
	carta nuevaCarta(){
		//TODO genera una nueva carta que no sea igual
	}
	
	void repartir(){
		//TODO repartir genera aleatoriamente las 5 primeras cartas
		//estas cartas no se pueden repetir, con lo que se tendrá que generar una a una. 
	}
	
	void descartarYRepartir(int seleccionadas[5]){
		//TODO a partir de  
	}
	
};

#endif
