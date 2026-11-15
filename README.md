PETNEXT – Pet Management System

Sistema de gestión de alimentación e higiene para mascotas desarrollado en Go (Golang).

1. Descripción del proyecto

PETNEXT es una aplicación de consola desarrollada en Go, diseñada para facilitar el registro y seguimiento de actividades relacionadas con el cuidado de mascotas, inicialmente enfocada en gatos.

El sistema permite administrar información de las mascotas, registrar sus actividades de alimentación y documentar la limpieza de sus areneros. Para conservar los registros utiliza archivos JSON, permitiendo consultar la información incluso después de cerrar y volver a ejecutar el programa.

El proyecto aplica conceptos de programación, estructuras de datos, validación de entradas, manejo de archivos y persistencia de información.

2. Objetivo general

Desarrollar un sistema de gestión mediante el lenguaje Go que permita organizar y consultar actividades de alimentación e higiene de mascotas, utilizando una estructura de programación que facilite el mantenimiento y la evolución del software.

3. Funcionalidades implementadas

Funcionalidad


* Registrar mascota

Registra nombre y edad de la mascota.


* Consultar mascotas

Muestra las mascotas activas registradas.


* Registrar alimentación

Guarda el tipo de alimento, cantidad en gramos y fecha.


* Registrar limpieza

Registra el arenero, la mascota relacionada, fecha y observaciones.


* Consultar historial

Muestra los registros de alimentación y limpieza.


* Salir

Finaliza la aplicación.

4. Tecnologías utilizadas

Lenguaje de programación: Go (Golang).

Interfaz: Consola.

Almacenamiento: Archivos JSON.

Control de versiones: Git y GitHub.

Editor recomendado: Visual Studio Code.

Paquetes de la biblioteca estándar

Paquete

Aplicación

fmt

Mostrar información en consola.

bufio

Leer datos ingresados por el usuario.

encoding/json

Convertir estructuras y almacenar información JSON.

os

Crear carpetas, leer y escribir archivos.

strconv

Convertir entradas de texto en números.

strings

Procesar cadenas de texto.

time

Registrar fecha y hora de las actividades.

No se requieren dependencias externas.

5. Estructura del proyecto

La versión actual utiliza una estructura sencilla, concentrando la lógica principal en main.go.

PETNEXT/
│
├── main.go
├── go.mod
├── README.md
├── .gitignore
│
└── data/
    ├── cats.json
    ├── feeding.json
    └── cleaning.json

Los archivos JSON se generan automáticamente cuando se guardan registros. Por ello, no es necesario crearlos manualmente.

6. Requisitos de instalación

Para ejecutar PETNEXT se necesita:

Go versión 1.22 o superior.

Git (opcional, para clonar el repositorio).

Visual Studio Code u otro editor compatible.



7. Almacenamiento de información

PETNEXT utiliza archivos JSON para conservar los registros generados durante su ejecución.

Archivo

Información almacenada

cats.json

Datos de las mascotas registradas.

feeding.json

Historial de alimentación.

cleaning.json

Historial de limpieza de areneros.

La aplicación crea la carpeta data automáticamente si no existe.

Nota: los registros son locales y se almacenan en el equipo donde se ejecuta el programa. No existe sincronización en la nube.

8. Aplicación de conceptos de programación

El desarrollo de PETNEXT permite aplicar los siguientes conceptos:

Estructuras (struct): representan las entidades Cat, Feeding y Cleaning.

Funciones: organizan las operaciones del sistema.

Condicionales: controlan las opciones del menú y las validaciones.

Ciclos: permiten mantener activo el menú principal.

Persistencia: conserva la información mediante archivos JSON.

Manejo de errores: contempla problemas de lectura y escritura de archivos.

Las estructuras representan los datos principales del dominio, mientras que las funciones organizan las operaciones correspondientes. La separación en paquetes y la incorporación de métodos asociados a las estructuras quedan como oportunidades de evolución de la arquitectura.

9. Alcance y futuras mejoras

Implementado en la versión actual

Registro de mascotas.

Consulta de mascotas activas.

Registro de alimentación.

Registro de limpieza de areneros.

Consulta de historiales.

Persistencia mediante JSON.

Mejoras propuestas

Separación de responsabilidades en paquetes independientes.

Administración de otras especies de mascotas.

Programación de horarios de alimentación.

Recordatorios de limpieza.

Reportes estadísticos.

Incorporación de pruebas automatizadas.

Interfaz gráfica.

Estas mejoras son propuestas para futuras versiones y no forman parte de las funcionalidades implementadas actualmente.

10. Autoría

Nombre del proyecto: PetNext
Tipo: Proyecto académico
Área: Programación y desarrollo de software
Lenguaje: Go (Golang)
Versión: 1.0

PetNext – Tecnología aplicada al cuidado de las mascotas.
