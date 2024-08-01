package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.kafka.core.KafkaTemplate;

@SpringBootApplication
public class DemoApplication {

    private static DirectoryWatcher directoryWatcher;

    @Autowired
    public DemoApplication(DirectoryWatcher directoryWatcher) {
        this.directoryWatcher = directoryWatcher;
    }

    public static void main(String[] args) {
        SpringApplication.run(DemoApplication.class, args);

        // 在构造函数中注入的 directoryWatcher 对象可以直接使用
        directoryWatcher.watchDirectory();
    }

}

