# fullcycle-cepclima


Uso:

No arquivo .env adicione o seu secret do weather api.

Rode o comando:

```
docker-compose up --build
```

Teste local:

```
curl http://localhost:8080/weather/DIGITE_O_CEP_AQUI
```

Teste Google Cloud Run:

```
https://fullcycle-cepclima-e26j5gs6xa-uc.a.run.app/weather/DIGITE_O_CEP_AQUI
```


Rodar os testes:

```
go test ./handlers 
```