package com.adidas.backend.publicservice.controller;

import org.apache.http.client.utils.URIBuilder;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;
import java.net.URISyntaxException;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

@RestController
@RequestMapping(value = "/public")
public class PublicRestController {
                                                            //
    private static final String prioritySaleService = "http://priority-sale-service:8080/priority-sale/addClient";
    private static final HttpClient client = HttpClient.newHttpClient();

    @PostMapping("/addClient")
    public ResponseEntity<String> addClient(@RequestBody String emailAddress) {
        try {

            URIBuilder builder = new URIBuilder(prioritySaleService);
            builder.setParameter("emailAddress", emailAddress); 

            // Crear una nueva solicitud HTTP con la URL construida
            HttpRequest requestWithParams = HttpRequest.newBuilder()
            .uri(builder.build()) // Usar la URL construida
            .GET()
            .build();
            
            HttpResponse<String> response = client.send(requestWithParams, HttpResponse.BodyHandlers.ofString());
            
            return ResponseEntity.ok(response.body());
        } catch (IOException | URISyntaxException | InterruptedException e) {
            e.printStackTrace();
            return ResponseEntity.status(500).body("Error en la comunicación con el servicio");
        }
    }

    @GetMapping("/dummy")
    public ResponseEntity<String> getDummyEndpointResponse() {
        String response = "Hello, this is a dummy response from public service";
        return ResponseEntity.ok(response);
    }
}
