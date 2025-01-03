package com.adidas.backend.prioritysaleservice.services;

import com.adidas.backend.prioritysaleservice.mainLogic.PrioritySaleMainLogic;
import com.adidas.backend.prioritysaleservice.models.Client;

import java.util.PriorityQueue;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping("/priority-sale")
public class PrioritySaleService {
    @Autowired
    PrioritySaleMainLogic clientMainLogic;
    
    @GetMapping("/addClient")
    public  PriorityQueue<Client> addClient(@RequestParam String emailAddress){
        return clientMainLogic.addClient(emailAddress);
    }
}
