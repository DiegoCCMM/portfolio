package com.adidas.backend.publicservice.service;

import com.adidas.backend.publicservice.kafkapublisher.KafkaConfig;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Import;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.SendResult;
import org.springframework.stereotype.Service;
import org.springframework.util.concurrent.ListenableFuture;

import java.util.concurrent.CompletableFuture;

@Service
@Slf4j
@Import(KafkaConfig.class)
public class KafkaProducerService {
    @Autowired
    KafkaTemplate<String, String> publicServiceKafkaTemplate;
    public void send(String topicName, String value) {
        CompletableFuture<SendResult<String, String>> future = publicServiceKafkaTemplate.send(topicName, value).completable();
        future.whenComplete((sendResult, exception) -> {
            if (exception != null) {
                future.completeExceptionally(exception);
            } else {
                future.complete(sendResult);
            }
            log.info(String.format("Task status send to Kafka topic : %s, Object: ", topicName)+ value);
        });
    }
}