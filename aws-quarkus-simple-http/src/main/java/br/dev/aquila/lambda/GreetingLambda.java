package br.dev.aquila.lambda;

import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.APIGatewayProxyRequestEvent;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

public class GreetingLambda implements RequestHandler<APIGatewayProxyRequestEvent, ApiGatewayResponse> {

    @Override
    public ApiGatewayResponse handleRequest(APIGatewayProxyRequestEvent input, Context context) {
       
        try {
            JsonNode body = new ObjectMapper().readTree(input.getBody());
            var response = new Response().setMessage("Hello, " + body.get("name").asText());
            return ApiGatewayResponse.builder()
                    .setStatusCode(200)
                    .setObjectBody(response)
                    .build();
        } catch (Exception e) {
            e.printStackTrace();
        }
        return ApiGatewayResponse.builder()
                .setStatusCode(400)
                .setObjectBody(new Response().setMessage("O atributo name não foi enviado ou ocorreu um erro inesperado."))
                .build();
    }

}
