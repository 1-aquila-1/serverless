package main

type Response struct {
	Url string `json:"url"`
}

type InputDTO struct {
	Nome     string `json:"nome"`
	Conteudo string `json:"conteudo"`
}
