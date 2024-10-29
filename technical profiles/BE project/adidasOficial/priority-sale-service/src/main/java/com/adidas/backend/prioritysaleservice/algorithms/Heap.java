package com.adidas.backend.prioritysaleservice.algorithms;

import com.adidas.backend.prioritysaleservice.models.Client;

import java.util.Comparator;
import java.util.PriorityQueue;

import org.springframework.stereotype.Component;

@Component
public class Heap {

    private Comparator<Client> cmp = Comparator.comparing(Client::isAdiClient).thenComparing(Client::compareTo);
    //it uses the current logic to create a MIN heap, for readability its reversed on use.
    private PriorityQueue<Client> maxHeap = new PriorityQueue<Client>(cmp.reversed());

    public PriorityQueue<Client> getMaxHeap() {
        return maxHeap;
    }

    public PriorityQueue <Client> orderStreamedList(Client c) {
        maxHeap.add(c);
        return maxHeap;
    }
}
