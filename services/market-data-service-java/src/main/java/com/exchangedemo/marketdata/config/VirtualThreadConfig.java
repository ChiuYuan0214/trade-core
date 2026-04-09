package com.exchangedemo.marketdata.config;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import org.apache.coyote.AbstractProtocol;
import org.apache.coyote.ProtocolHandler;
import org.springframework.boot.autoconfigure.task.TaskExecutionAutoConfiguration;
import org.springframework.boot.web.embedded.tomcat.TomcatProtocolHandlerCustomizer;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.task.AsyncTaskExecutor;
import org.springframework.core.task.support.TaskExecutorAdapter;

@Configuration
public class VirtualThreadConfig {

    @Bean(destroyMethod = "close")
    public ExecutorService marketDataVirtualThreadExecutor() {
        return Executors.newVirtualThreadPerTaskExecutor();
    }

    @Bean(name = TaskExecutionAutoConfiguration.APPLICATION_TASK_EXECUTOR_BEAN_NAME)
    public AsyncTaskExecutor applicationTaskExecutor(ExecutorService marketDataVirtualThreadExecutor) {
        return new TaskExecutorAdapter(marketDataVirtualThreadExecutor);
    }

    @Bean
    public TomcatProtocolHandlerCustomizer<ProtocolHandler> tomcatVirtualThreadCustomizer(
            ExecutorService marketDataVirtualThreadExecutor
    ) {
        return protocolHandler -> {
            if (protocolHandler instanceof AbstractProtocol<?> abstractProtocol) {
                abstractProtocol.setExecutor(marketDataVirtualThreadExecutor);
            }
        };
    }
}
