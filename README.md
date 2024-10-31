Repositorio de Algoritmia en Go
Este repositorio contiene ejemplos de algoritmos y estructuras de datos implementados en Go (Golang). Está diseñado para ayudar a los desarrolladores a familiarizarse con conceptos básicos de algoritmia y proporcionar pruebas unitarias para asegurar la funcionalidad de los algoritmos.

Estructura del Repositorio
La estructura del repositorio es la siguiente:

``
├── cmd
├── config
├── internal
│   ├── application
│   │   ├── src
│   │   │   ├── algorithms
│   │   │   │   ├── basic_problems
│   │   │   │   ├── recursion
│   │   │   │   │   ├── binary_search_recursive
│   │   │   │   │   ├── factorial
│   │   │   │   │   ├── fibonacci
│   │   │   │   │   └── summation
│   │   │   │   ├── searching
│   │   │   │   │   └── binary_search
│   │   │   │   └── simple_linked_list_applications
│   │   │   └── data_structures
│   │   │       ├── circular_linked_list
│   │   │       ├── double_linked_list
│   │   │       ├── linked_list
│   │   │       ├── lista_ligada
│   │   │       ├── queue
│   │   │       │   ├── applications
│   │   │       │   └── implementations
│   │   │       │       ├── array
│   │   │       │       └── linkedlist
│   │   │       └── stack
│   │   │           ├── applications
│   │   │           │   └── parenthesis_validation
│   │   │           ├── array
│   │   │           ├── double_linked_list
│   │   │           └── linkedlist
│   │   └── test
│   │       └── application
│   │           ├── algorithms
│   │           │   ├── basic_problems
│   │           │   ├── recursion
│   │           │   │   ├── binary_search_test
│   │           │   │   ├── factorial_test
│   │           │   │   ├── fibonacci_test
│   │           │   │   └── summation_test
│   │           │   ├── searching
│   │           │   │   └── binary_search_test
│   │           │   └── simple_linked_list_applications
│   │           └── data_structures
│   │               ├── queue
│   │               ├── simple_linked_list
│   │               └── stack
│   │                   └── applications
│   ├── domain
│   └── infrastructure
└── pkg
└── middleware
``
Descripción de las Carpetas
cmd: Aquí se pueden encontrar los comandos para ejecutar la aplicación o pruebas.
config: Archivos de configuración necesarios para la aplicación.
internal: Contiene la lógica de la aplicación, dividida en subcarpetas para facilitar la organización:
application/src: Implementaciones de algoritmos y estructuras de datos.
application/test: Pruebas unitarias asociadas a los algoritmos y estructuras de datos.
domain: Modelos de dominio que representan las entidades de negocio.
infrastructure: Implementaciones de infraestructura, como la conexión a bases de datos.
pkg: Paquetes reutilizables, en este caso, middleware que se puede utilizar en la aplicación.
Ejemplos de Algoritmos
Recursión
Binary Search (Recursivo): Implementación del algoritmo de búsqueda binaria utilizando recursión.
Factorial: Cálculo del factorial de un número utilizando recursión.
Fibonacci: Generación de la secuencia de Fibonacci a través de una función recursiva.
Summation: Suma de una serie de números utilizando un enfoque recursivo.
Estructuras de Datos
Listas Enlazadas: Implementaciones de listas enlazadas simples, dobles y circulares.
Colas: Implementaciones de colas utilizando arreglos y listas enlazadas, junto con aplicaciones prácticas.
Pilas: Implementación de pilas y su uso en la validación de paréntesis.
Pruebas Unitarias
Cada algoritmo y estructura de datos incluye pruebas unitarias para garantizar su correcto funcionamiento. Estas pruebas se encuentran en la carpeta internal/application/test.

Cómo Ejecutar las Pruebas
Para ejecutar las pruebas, asegúrate de tener Go instalado y utiliza el siguiente comando en la raíz del repositorio:

bash
Copiar código
go test ./internal/application/test/...
Contribuciones
Si deseas contribuir a este proyecto, no dudes en abrir un Pull Request o crear un problema en el repositorio. Las contribuciones son bienvenidas.

Licencia
Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

Si necesitas más detalles o cambios específicos, ¡hazmelo saber!