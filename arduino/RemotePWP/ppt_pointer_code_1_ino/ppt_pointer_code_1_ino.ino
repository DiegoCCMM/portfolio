#include <IRremote.h>

int RECV_PIN = 2;

IRrecv irrecv(RECV_PIN);

decode_results results;
int key_pressed;

void setup()
{
  Serial.begin(9600);
  irrecv.enableIRIn(); // Start the receiver
}

void loop() {
  if (irrecv.decode(&results)) {
    
    key_pressed = results.value;
    Serial.println(key_pressed);
    
    irrecv.resume(); // Receive the next value
  }
  delay(100);
}