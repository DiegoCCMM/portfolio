
// Arduino pin numbers
const int X_pin = 1; // analog pin connected to X output
const int Y_pin = 0; // analog pin connected to Y output

const int ledPinR =  3;      // the number of the LED pin
const int ledPinY =  4;      // the number of the LED pin
const int ledPinB =  5;      // the number of the LED pin
const int ledPinG =  6;      // the number of the LED pin

const int buzzerPin = 9;

// variables will change:
int leds;
int wins = 0;

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
}


void ledOff(const int& whichOne) {
  digitalWrite(whichOne, LOW);
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

void winState(){
  for (int i = 0; i < 50; i++)
  {
    digitalWrite(buzzerPin, HIGH);
    delay(3);//wait for 1ms
    digitalWrite(buzzerPin, LOW);
    delay(3);//wait for 1ms
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

bool correcto(const double& dx, const double& dy){
  return((leds == 3 && dx == 1.0) || (leds == 4 && dy == 0.0) 
  || (leds == 5 && dy == 1.0) || (leds == 6 && dx == 0.0));
}

void loop() {
  
  leds = createNumber();

  ledOn(leds);

  delay(500);

  
  int x, y;
  double dx, dy;
  x = analogRead(X_pin);
  y = analogRead(Y_pin);
  
  dx = x / 1023.0;
  dy = y / 1023.0;

  ledOff(leds);

  
  bool correct = correcto(dx, dy);
  
  if (correct) { //next turn
    successState();
    delay(250);
    wins++;
    if(wins == 5)
    {
      winState();
      wins = 0;
    }
  } else { //se ha equivocado
    failState();
    delay(500);
  }

  
  
  
}