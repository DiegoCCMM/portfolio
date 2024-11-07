package com.adidas.backend.prioritysaleservice.mainLogic;

import com.adidas.backend.prioritysaleservice.models.*;
import com.adidas.backend.prioritysaleservice.algorithms.Heap;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.apache.http.client.utils.URIBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestParam;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.PriorityQueue;
import java.util.Timer;
import java.util.TimerTask;

@Component
public class PrioritySaleMainLogic {
    @Autowired
    private Heap hp = new Heap();

    private Timer timer = new Timer();
                                                        //adi-club-service
    private static final String adiClubService = "http://localhost:8080/adiclub";
    private static final String emailServiceUrl = "http://localhost:8083/email-service/send"; // URL del otro microservicio
    private static final HttpClient client = HttpClient.newHttpClient();

    public PrioritySaleMainLogic() {
        timer.schedule(new MyTask(), 0, 10 * 1000);
    }

    public PriorityQueue<Client> addClient(String emailAddress) {
        try {
            
            URIBuilder builder = new URIBuilder(adiClubService);
            builder.setParameter("emailAddress", emailAddress); 

            // Crear una nueva solicitud HTTP con la URL construida
            HttpRequest requestWithParams = HttpRequest.newBuilder()
            .uri(builder.build()) // Usar la URL construida
            .GET()
            .build();
                        
            HttpResponse<String> response = client.send(requestWithParams, HttpResponse.BodyHandlers.ofString());

            if (response.statusCode() == 200) {
                ObjectMapper objectMapper = new ObjectMapper();
                AdiClubClient newClient = objectMapper.readValue(response.body(), AdiClubClient.class);

                hp.orderStreamedList(newClient); 
                
                return hp.getMaxHeap();
            } else {
                System.err.println("Error: " + response.statusCode() + " - " + response.body());
            }
        } catch (IOException | InterruptedException | URISyntaxException e) {
            e.printStackTrace();
        }
        return hp.getMaxHeap(); 
    }
    

    class MyTask extends TimerTask {
        public void run() {
            System.out.println("Entered");
            
            Client winner = hp.pop();
            winner = new Client("diego@gmail.com");
            if (winner != null) {
                try {
                    System.out.println("Request");
                    HttpRequest request = HttpRequest.newBuilder()
                            .uri(new URI(emailServiceUrl))
                            .header("Content-Type", "application/json")
                            .POST(HttpRequest.BodyPublishers.ofString("\"" + winner.getEmail() + "\"")) // Enviar sólo el email como String
                            .build();

                    HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

                    if (response.statusCode() == 200) {
                        System.out.println("Email enviado a " + winner.getEmail());
                    } else {
                        System.err.println("Error al enviar email: " + response.statusCode() + " - " + response.body());
                    }
                } catch (IOException | InterruptedException | URISyntaxException e) {
                    e.printStackTrace();
                }
            } else {
                System.out.println("Heap vacío, no se encontraron usuarios para procesar.");
            }
        }
    }
}
