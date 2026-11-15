PETNEXT – Pet Management System

Sistema de gestión de alimentación e higiene para mascotas desarrollado en Go (Golang).

1. Descripción del proyecto

PETNEXT es una aplicación de consola desarrollada en Go, diseñada para facilitar el registro y seguimiento de actividades relacionadas con el cuidado de mascotas, inicialmente enfocada en gatos.

El sistema permite administrar información de las mascotas, registrar sus actividades de alimentación y documentar la limpieza de sus areneros. Para conservar los registros utiliza archivos JSON, permitiendo consultar la información incluso después de cerrar y volver a ejecutar el programa.

El proyecto aplica conceptos de programación, estructuras de datos, validación de entradas, manejo de archivos y persistencia de información.

2. Objetivo general

Desarrollar un sistema de gestión mediante el lenguaje Go que permita organizar y consultar actividades de alimentación e higiene de mascotas, utilizando una estructura de programación que facilite el mantenimiento y la evolución del software.

3. Funcionalidades implementadas

Opción

Funcionalidad

Descripción

1

Registrar mascota

Registra nombre y edad de la mascota.

2

Consultar mascotas

Muestra las mascotas activas registradas.

3

Registrar alimentación

Guarda el tipo de alimento, cantidad en gramos y fecha.

4

Registrar limpieza

Registra el arenero, la mascota relacionada, fecha y observaciones.

5

Consultar historial

Muestra los registros de alimentación y limpieza.

6

Salir

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

7. Instalación y ejecución

Paso 1. Clonar el repositorio

git clone https://github.com/TU-USUARIO/PETNEXT.git

Reemplaza TU-USUARIO por el nombre de usuario correspondiente a tu cuenta de GitHub.

Paso 2. Acceder al proyecto

cd PETNEXT

Paso 3. Inicializar el módulo (si no existe go.mod)

go mod init petnext

Si el archivo go.mod ya existe, no es necesario ejecutar este comando.

Paso 4. Ejecutar el sistema

go run .

Paso 5. Compilar el proyecto

go build .

8. Almacenamiento de información

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

9. Organización lógica del código

Aunque la implementación actual se encuentra en un único archivo, sus responsabilidades están diferenciadas en funciones:

Gestión de mascotas: registro y consulta de gatos.

Gestión de alimentación: registro de alimentos, cantidades y fechas.

Gestión de limpieza: registro de actividades relacionadas con los areneros.

Historial: consulta de las actividades almacenadas.

Almacenamiento: lectura y escritura de archivos JSON.

10. Aplicación de conceptos de programación

El desarrollo de PETNEXT permite aplicar los siguientes conceptos:

Estructuras (struct): representan las entidades Cat, Feeding y Cleaning.

Funciones: organizan las operaciones del sistema.

Condicionales: controlan las opciones del menú y las validaciones.

Ciclos: permiten mantener activo el menú principal.

Persistencia: conserva la información mediante archivos JSON.

Manejo de errores: contempla problemas de lectura y escritura de archivos.

Las estructuras representan los datos principales del dominio, mientras que las funciones organizan las operaciones correspondientes. La separación en paquetes y la incorporación de métodos asociados a las estructuras quedan como oportunidades de evolución de la arquitectura.

11. Alcance y futuras mejoras

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

12. Autoría

Nombre del proyecto: PETNEXT
Tipo: Proyecto académico
Área: Programación y desarrollo de software
Lenguaje: Go (Golang)
Versión: 1.0

PETNEXT – Tecnología aplicada al cuidado de las mascotas.
