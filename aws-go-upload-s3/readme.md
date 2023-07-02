<h1>Upload de imagem</h1>

Dependencia
```
github.com/aws/aws-lambda-go/events

github.com/aws/aws-sdk-go-v2/feature/s3/manager
github.com/aws/aws-sdk-go-v2/service/s3
```

Criar um arquivo `.env` na raiz do projeto com a variável
```
ARN-BUCKET=ARN_AWS_do_bucket
```

Definir a variável de `ambiente da lambda` na AWS

Documentação
```bash
# AWS
https://docs.aws.amazon.com/pt_br/code-library/latest/ug/go_2_s3_code_examples.html

# Variáveis de ambiente
https://docs.aws.amazon.com/pt_br/lambda/latest/dg/configuration-envvars.html#configuration-envvars-config

# Serverless
https://www.serverless.com/framework/docs/providers/aws/guide/variables/

# Serverless
https://github.com/neverendingqs/serverless-dotenv-plugin
```