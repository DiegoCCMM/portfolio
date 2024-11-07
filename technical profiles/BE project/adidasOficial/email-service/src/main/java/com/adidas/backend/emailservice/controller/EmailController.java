package com.adidas.backend.emailservice.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/email-service")
public class EmailController {

    @PostMapping("/send")
    public String sendEmail(@RequestBody String email) {
        System.out.println("Correo recibido para el cliente con email: " + email);
        return "Correo recibido para el email: " + email;
    }
}