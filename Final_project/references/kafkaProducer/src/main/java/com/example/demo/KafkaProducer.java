package com.example.demo;

import jakarta.annotation.Resource;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;
import java.io.BufferedReader;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

@Component
public class KafkaProducer {
    @Resource
    private KafkaTemplate<String, String> kafkaTemplate;

    private static String filePath = "D:/dxy/iso/data_hz.csv";
    // 示例代码
    public void sendMessage() throws IOException{
        ScheduledExecutorService executorService = Executors.newSingleThreadScheduledExecutor();
        BufferedReader reader = Files.newBufferedReader(Paths.get(filePath));

        Runnable task = new Runnable() {
            @Override
            public void run() {
                try {
                    String line = reader.readLine();
                    if (line != null) {
                        kafkaTemplate.send("weibo",line);
                    } else {
                        // 文件结束，关闭资源并停止定时器
                        reader.close();
                        executorService.shutdown();
                    }
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        };

        // 每隔0.5秒执行一次任务
        executorService.scheduleAtFixedRate(task, 0, 1000, TimeUnit.MILLISECONDS);
    }
}

