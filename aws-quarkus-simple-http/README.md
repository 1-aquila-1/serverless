# aws-quarkus-simple-http-api

Este projeto utiliza Quarkus e Serverless Framework.

Website:

https://quarkus.io/

https://www.serverless.com/framework/docs/providers/aws


## COMANDO

Execute os camandos abaixo para geração do .zip e para fazer o deploy na AWS:
```shell script
gradle clean build

sls deploy
```

> **_OBS:_** Precisa do AWS CLI instalado e configurado para executar o comando `sls deploy`

Instalação

https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

Configuração das credenciais: 
   - AWS Access Key ID
   - AWS Secret Access Key

https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html


> **_NOTE:_**  URL e request de exemplo.

```
POST https://sua.url.da.amazonaws.com/dev/hello

{
   "name": "Will"
}
```