1. Iniciar Proyecto

go mod init github.com/greetings

Package for install:
1. go get github.com/joho/godotenv
2. go get -u github.com/gin-gonic/gin
3. go get go.mongodb.org/mongo-driver/mongo


Buenas practicas:
Poner el tiempo en UTC
```
player.CreationTime = time.Now().UTC()
```

RAMAS CREADA HASTA EL MOMENTO PARA CREAR LA ARQUITECTURA HEXAGONAL:
1. conect/db-mongo
2. handler/como-se-conforma
3. init/gin
4. init/project-go
5. main
6. mongodb/capa-ports-arq-hexagonal
7. ports/separando-service-and-folder-ports
8. ultima-capa-conexion-db