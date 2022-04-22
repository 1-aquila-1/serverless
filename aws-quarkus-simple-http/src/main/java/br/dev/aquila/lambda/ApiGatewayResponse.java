package br.dev.aquila.lambda;

import java.util.Collections;
import java.util.Map;

import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyResponseEvent;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class ApiGatewayResponse extends APIGatewayProxyResponseEvent {

    public ApiGatewayResponse(int statusCode, String body, Map<String, String> headers) {
        setStatusCode(statusCode);
        setBody(body);
        setHeaders(headers);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static class Builder {

        private static final ObjectMapper objectMapper = new ObjectMapper();

        private int statusCode = 200;
        private Map<String, String> headers = Collections.emptyMap();
        private Object objectBody;

        public Builder setStatusCode(int statusCode) {
            this.statusCode = statusCode;
            return this;
        }

        public Builder setHeaders(Map<String, String> headers) {
            this.headers = headers;
            return this;
        }

        public Builder setObjectBody(Object objectBody) {
            this.objectBody = objectBody;
            return this;
        }

        public ApiGatewayResponse build() {
            String body = null;
            if (objectBody != null) {
                try {
                    body = objectMapper.writeValueAsString(objectBody);
                } catch (JsonProcessingException e) {
                    throw new RuntimeException(e);
                }
            }
            return new ApiGatewayResponse(statusCode, body, headers);
        }
    }
}
