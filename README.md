
# Clasificador de Bases de Datos

Este proyecto contiene el desarrollo de una aplicacion que permite clasificar bases de datos MySQL.

La clasificacion se basa en obtener los nombres de esquemas, tablas y columnas de la base de datos a analizar, para luego clasificar los nombres de columnas y nombres de tablas en base a un conjunto de Tipos de Informacion definidos como un par (Nombre de Tipo de Informacion: Patron Regex). Dicho patron regex se plantea como una lista de palabras que en caso de que el nombre de columna o tabla pertenezca a dicha lista, entonces se considera que pertenece o es del Nombre de Tipo de Informacion correspondiente.

La aplicacion se basa en 3 servicios:
1. API desarrollada en Golang y Gin.
2. Base de Datos del Clasificador.
3. Base de Datos MySQL de Ejemplo a clasificar.

Ademas se incluye un servicio extra PHPMyAdmin para administrar Base de Datos del Clasificador.

## Tabla de Contenidos

- [Requisitos](#requisitos)
- [Ejecucion de la API](#ejecución-de-la-api)
- [Descripcion API Endpoints](#descripcion-api-endpoints)
- [Estructura de Directorios](#estructura-de-directorios)
- [Development Tools](#development-tools)

## Requisitos

- **Docker**: Version 20.10 o superior
- **Docker Compose**: Version 1.29 o superior

## Ejecucion de la API

1. **Configurar las variables de entorno**:

   Asegúrate de tener un archivo \`.env\` en el directorio /src del proyecto.

   Se incluye un archivo para facilitar la ejecucion de la aplicacion

2. **Construir y ejecutar los contenedores**:

    ``` bash
    docker-compose up --build
    ```

3. **Acceder API**:

    La API estará disponible en `http://localhost:8081`.

## Descripcion API Endpoints

Se lista cada endpoint y responsabilidad:

1. GET:  /api/v1/ping               --> Health check  

2. POST: /api/v1/database           --> Cargar una configuracion base de datos MySQL a clasificar, devolviendo un id asociado a esta
   . Ejemplo de Solicitud:
   ```
   {
        "host": "db_courses",
        "port": "3306",
        "username": "root",
        "password": "courses123"
   }
   ```  

3. POST: /api/v1/database/scan/:id  --> Realizar un escaneo a una base de datos basado en el id asociado a la misma  

4. GET:  /api/v1/database/scan/:id  --> Devolver los resultados de un escaneo de base de datos basado en el id asociado a la misma  

5. POST: /api/v1/database/infotype/column --> Cargar un nuevo tipo de informacion para clasificacion de columnas
   . Ejemplo de Solicitud:
   ```
    {
        "name": "EMAIL",
        "description": "",
        "regex": "\\b(email)\\b|\\b(EMAIL)\\b|\\b(emailAddress)\\b"
    }
   ```  

6. GET: /api/v1/database/infotype/column --> Listar tipos de informacion de clasifiacion de columnas disponibles  

7. POST: /api/v1/database/infotype/table --> Cargar un nuevo tipo de informacion para clasificacion de tablas  
   . Ejemplo de Solicitud:
   ```
    {
        "name": "ADMINISTRATIVE",
        "description": "",
        "regex": "\\b(registrations)\\b|\\b(Registrations)\\b|\\b(payments)\\b"
    }
   ```  

8. GET: /api/v1/database/infotype/column --> Listar tipos de informacion de clasifiacion de tablas disponibles

## Estructura de Directorios

A continuación se muestra la estructura básica de directorios del proyecto:

``` 
├── src/                        # Codigo fuente de la app
|    ├── config/                # Configuraciones de la aplicación
|    ├── controllers/           # Controladores para manejar las rutas
|    ├── dto/                   # Esquemas de body JSON para solicitudes y respuestas HTTP
|    ├── models/                # Definición de modelos de la base de datos
|    ├── repositories/          # Logica de interaccion con bases de datos
|    ├── routes/                # Definición de rutas
|    ├── services/              # Lógica de negocio y servicios
|    ├── go.mod                 # Archivo de dependencias de Go
|    ├── go.sum                 # Resumen de versiones de las dependencias
|    ├── main.go                # Punto de entrada de la aplicación
|    └── .env                   # Variables de entorno (Este archivo por lo general no se dispone en el repositorio)
├── sql_scripts                 # SQL scripts para creacion e inicializacion de base de datos
├── Dockerfile                  # Dockerfile para crear la imagen del contenedor de la API
└── docker-compose.yml          # Archivo de Docker Compose para orquestar todos los contenedores
``` 

## Development Tools

### PHPMyAdmin
En caso de querer administrar la base de datos de clasificador:

1. Acceder a PHPMyAdmin URL: `http://localhost:8080`.
2. Logearse con credenciales: 
   `user: "root"`
   `password: "clasif456"`