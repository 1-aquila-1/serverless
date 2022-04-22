package br.dev.aquila.lambda;

public class Response {

    private String message;

    public String getMessage() {
        return message;
    }

    public Response setMessage(String message) {
        this.message = message;
        return this;
    }
}
