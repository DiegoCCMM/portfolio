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
    private static final HttpClient client = HttpClient.newHttpClient();


    public PriorityQueue<Client> getAllClients() {
        timer.schedule(new MyTask(), 0, 60*1000);
        try {
            ObjectMapper objectMapper = new ObjectMapper();
            AdiClubClient adiClubClient = objectMapper.readValue(client.send(request,
                    HttpResponse.BodyHandlers.ofString()).body(), AdiClubClient.class);
            hp.orderStreamedList(adiClubClient);
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
        return hp.getMaxHeap();
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
    
}
class MyTask extends TimerTask {
    public void run() {
        // TODO: pop from heap then call email
        System.out.println("Function called!");
    }
}
