package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.io.IOException;

@Component
public class DirectoryWatcher {
    private final KafkaProducer kafkaProducer;

    @Autowired
    public DirectoryWatcher(KafkaProducer kafkaProducer) {
        this.kafkaProducer = kafkaProducer;
    }

    public void watchDirectory() {
        try{
            kafkaProducer.sendMessage();
        } catch (IOException e){
            e.printStackTrace();
        }

    }
}
