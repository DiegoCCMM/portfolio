

// constants won't change. They're used here to set pin numbers:
const int buttonPinR = 10;     // the number of the pushbutton pin
const int buttonPinY = 11;     // the number of the pushbutton pin
const int buttonPinB = 12;     // the number of the pushbutton pin
const int buttonPinG = 13;     // the number of the pushbutton pin

const int ledPinR =  3;      // the number of the LED pin
const int ledPinY =  4;      // the number of the LED pin
const int ledPinB =  5;      // the number of the LED pin
const int ledPinG =  6;      // the number of the LED pin

const int buzzerPin = 9;


int leds;

void setup() {
  // initialize the LED pin as an output:
  pinMode(ledPinR, OUTPUT);
  pinMode(ledPinY, OUTPUT);
  pinMode(ledPinB, OUTPUT);
  pinMode(ledPinG, OUTPUT);
  pinMode(buzzerPin, OUTPUT);

  randomSeed(analogRead(7));
  
}

//numero aleatorio entre 3 y 6
int createNumber() {
  return (random(3, 7));
}

//encender el led whichOne
void ledOn(const int& whichOne) {
  digitalWrite(whichOne, HIGH);
  delay(750);
  digitalWrite(whichOne, LOW);
}

//esperamos al botón
void waitButton(const int& buttonPin, const int& led, int& answered) {
  if (digitalRead(buttonPin) == LOW) //si es el correcto perfecto
  {
    ledOn(buttonPin - 7);
    answered = 1;

    if (buttonPin != led + 7)
    {
      gameState = 2;
    }
  }
}

//esperamos a que el jugador responda, pudiendo responder el boton correcto
//o uno incorrecto
void waitForPush(const int& led)
{
  int answered = 0;

  while (answered == 0) {

    if (digitalRead(buttonPinG) == LOW) //si es el correcto perfecto
    {
      ledOn(buttonPinG - 7 );
      answered = 1;

      if (buttonPinG != led + 7)
      {
        gameState = 2;
      }
    }
    if (digitalRead(buttonPinY) == LOW) //si es el correcto perfecto
    {
      ledOn(buttonPinY - 7);
      answered = 1;

      if (buttonPinY != led + 7)
      {

        gameState = 2;
      }
    }
    if (digitalRead(buttonPinR) == LOW) //si es el correcto perfecto
    {
      ledOn(buttonPinR - 7);
      answered = 1;

      if (buttonPinR != led + 7)
      {

        gameState = 2;
      }
    }
    if (digitalRead(buttonPinB) == LOW) //si es el correcto perfecto
    {
      ledOn(buttonPinB - 7);
      answered = 1;

      if (buttonPinB != led + 7)
      {

        gameState = 2;
      }
    }

    /*waitButton(buttonPinB,led, answered);
    waitButton(buttonPinG,led, answered);
    waitButton(buttonPinY,led, answered);
    waitButton(buttonPinR,led, answered);*/
    
  }
}

//sonido de victoria
void successState() {

  for (int i = 0; i < 200; i++)
  {
    digitalWrite(buzzerPin, HIGH);
    delay(1);//wait for 1ms
    digitalWrite(buzzerPin, LOW);
    delay(1);//wait for 1ms
  }
}

//sonido de derrota
void failState() {

  for (int i = 0; i < 100; i++)
  {
    digitalWrite(buzzerPin, HIGH);
    delay(2);//wait for 1ms
    digitalWrite(buzzerPin, LOW);
    delay(2);//wait for 1ms
  }
}

void loop() {
  gameState = 0;

  leds[turn] = createNumber();

  for (int i = 0; i <= turn; i++)
  {
    delay(850);
    ledOn(leds[i]);
  }

  gameState = 1;

  for (int i = 0; i <= turn && gameState != 2; i++)
  {
    waitForPush(leds[i]);
  }

  if (gameState != 2) { //next turn
    turn++;
    successState();
  } else { //se ha equivocado
    turn = 0;
    failState();
  }

  gameState = 0;

}
