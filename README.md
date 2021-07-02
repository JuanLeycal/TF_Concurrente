# TF_Concurrente
Repositorio que contiene todo el TF del curso de Programación Concurrente.

### Introducción, descripción del problema y motivación

La definición de un “sistema distribuido” es la siguiente: Una colección de sistemas independientes que aparecen sus usuarios como una única computadora, lo cual les permite trabajar de manera concurrente. Los beneficios de tener este sistema varían según de qué manera están organizados los sistemas destinados a mejorar, ya sea sistemas centralizados, o sistemas aislados. Sin embargo, las ventajas generales de tener un sistema distribuido es poder escalar de manera horizontal tus sistemas, es más tolerante a fallos, más eficientes al procesar data,  poder implementar paralelismo, entre otros. 

Estos beneficios mencionados anteriormente nos pueden ayudar a resolver el problema que se describe a continuación. Existen varios algoritmos basados en machine learning que pueden analizar varias cantidades de datos para luego procesarla según el propósito del algoritmo de machine learning. Sin embargo, pueden haber casos donde los datos por analizar pueden llegar a ser bastante extensos, por lo cual el algoritmo desarrollado puede llegar a tardar demasiado en analizarlo, cosa que bajaría su rendimiento. Por esto en este trabajo, proponemos implementar un sistema distribuido a la API que contiene el algoritmo de machine learning, para que el proceso de este sea distribuido entre diferentes ordenadores, y así mejorar la eficiencia del algoritmo.

La mayoría de la motivación detrás de este proyecto es el hecho de aprender cómo analizar la implementación y el procesamiento de un sistema distribuido mediante este proyecto. Poder mejorar el rendimiento de un algoritmo sin necesidad de modificar el algoritmo en sí, sino su entorno. Finalmente, poder aplicar todos los principios de concurrencia, distribución y paralelismo enseñados en el curso en este trabajo.

### Objetivos

Como objetivo del presente trabajo, se tienen los logros del curso planteados en el sílabo del curso, junto a las competencias ABET planteadas por la división de Excelencia Académica de la UPC.

Primero, entre los logros del curso tenemos los siguientes:

- Manejo de la información: Capacidad de identificar la información necesaria, así como de buscarla, seleccionarla, evaluarla y usarla éticamente, con la finalidad de resolver un problema.

- Trabajo en Equipos Multidisciplinarios: Funcionar eficazmente como miembro o líder de un equipo que participa en actividades apropiadas para la disciplina del programa. La capacidad de funcionar efectivamente en un equipo cuyos miembros juntos proporcionan liderazgo, crean un entorno de colaboración e inclusivo, establecen objetivos, planifican tareas y cumplen objetivos.

Asimismo, tenemos la competencia ABET 5, el cual es definido por la capacidad de funcionar efectivamente en un equipo cuyos miembros juntos proporcionan liderazgo, crean un entorno de colaboración e inclusivo, establecen objetivos, planifican tareas y cumplen objetivos. Dicho item consta de 3 criterios:

- Participa en equipos multidisciplinarios con eficacia, eficiencia y objetividad, en el marco de un  proyecto en soluciones de tecnologías de la información.

- Conoce al menos un sector empresarial o dominio de aplicación de soluciones de tecnologías de la información.

- Conocimientos de nuevos métodos de colaboración y comunicación.

En conclusión, es necesario ser eficaz al transimitir información de manera oral y escrita, como tambien ser productivos en ambientes colaborativos, pudiendo transmitir mensajes utilizando medios de comunicación remotos y resolver problemas de manera ágil conociendo los sectores empresariales en las que debemos aplicar nuestra solución.

### Diseño (arquitectura o componentes, etc.)

Para la arquitectura de este trabajo se planteó utilizar una distribución de nodos de tipo peer 2 peer, en la cual los nodos no tienen una jerarquía definida y el sistema se encuentra descentralizado. Asimismo, de forma más específica, nuestra arquitectura está basada en el Token Ring Algorithm o arquitectura de tipo anillo, la cual según Tanenbaum y Van Steen: “A completely different approach to deterministically achieving mutual exclusion...” (2006).

Esta arquitectura plantea que, el anillo, está compuesto de una cierta cantidad de nodos, a los cuales se le asigna un proceso y cada posición del anillo tiene una posición numérica que puede ser una dirección de red. El orden de cada nodo como tal no es relevante, lo único relevante en esta arquitectura es que el nodo actual sepa cual es el nodo sucesor, tal como mencionan Tanenbaum y Van Steen.

Para nuestro caso, hemos planteado la utilización de 5 nodos en una arquitectura de tipo anillo, los cuales utilizarán la técnica SPMD (Simple programa, múltiple data) ya que cada nodo tendrá acceso al mismo algoritmo de ML, pero cada nodo procesa un segmento de data distinto al resto de nodos.

Finalmente, se cuenta con un archivo “init.go”, el cual es el encargado de inicializar el algoritmo, inicializar los nodos y es el que contiene la lógica del API para poder recibir los diversos “requests” por parte del user y enviarle “responses” en formato JSON.
A continuación, se adjunta un diagrama que contiene la arquitectura planteada para el presente proyecto:

Tal como se puede apreciar en la imagen, se tiene un nodo API, el cual se encarga de los requests y responses hacia el Front-End y de comunicarse con la red de nodos en arquitectura de tipo anillo. Asimismo, esta red está compuesta de 5 nodos que se encargan de procesar el algoritmo de entrenamiento y enviar la data procesada al nodo siguiente para que realice el mismo proceso. Finalmente, cuando se complete el número de iteraciones del entrenamiento de la data, el último nodo envía los clusters hacia el REST API, el cual se encargará de retornarla al Front-End en formato JSON y este último lo mostrará al usuario.

### Desarrollo

Para el desarrollo de este trabajo se está utilizando la herramienta VMWare Workstation para la creación de 5 máquinas virtuales las cuales representan los 5 nodos explicados previamente en la sección de arquitectura. Cada una de estas máquinas virtuales contiene una imagen del sistema operativo Debian 10 y han sido configuradas de tal modo que cada una posee una dirección IP única, distinta a la de la máquina en la cual se ejecutan todos los nodos.

Para la transferencia de archivos desde nuestra máquina local hacia cada uno de los nodos, se está utilizando la herramienta MobaXTerm. Para la correcta utilización de esta, fue necesario habilitar el acceso SSH al root en el archivo “/etc/ssh/sshd_config” para poder realizar una conexión entre la herramienta y la máquina virtual. A continuación, se procedió a crear 5 sesiones con sus respectivas credenciales asociadas a cada máquina virtual, con el objetivo de facilitar la transferencia de archivos. Gracias a esta configuración, se pueden compartir los archivos necesarios para la correcta ejecución del algoritmo a cada máquina virtual.

Para el desarrollo del algoritmo, se está utilizando el lenguaje de programación GO, el cual nos permite ejecutar procesos de forma concurrente en cada máquina virtual. Asimismo, para la conexión entre máquinas virtuales, se está utilizando el paquete “net” el cual contiene las funciones Dial y Listen, las cuales nos permiten comunicar las máquinas virtuales a través de su dirección IP y utilizando el protocolo TCP.

Con respecto al algoritmo se está desarrollando el algoritmo K Means, en el cual la data recolectada de https://www.datosabiertos.gob.pe/ es asignada a un cluster determinado por los valores cuantitativos de sus columnas. Finalmente, con respecto al algoritmo K-means se ha implementado un algoritmo concurrente en el cual en cada iteración se calculan los centroides utilizando los parámetros provenientes del dataset. 

Tal como mencionan Morissette y Chartier k-means clustering, pertenece a las técnicas de agrupamiento basadas en partición, las cuales a su vez se basan en la relocación iterativa de los data points entre clusters. (2013). En nuestro caso, cada iteración se realiza de forma distribuida en un nodo perteneciente a la arquitectura. De esta forma nos aseguramos que el procesamiento del algoritmo utilice todos los recursos disponibles provenientes de las 5 máquinas virtuales pertenecientes a la red. Al finalizar las iteraciones, se envía la data al nuestro nodo API, el cual lo retornará al frontend en formato JSON.

### Conclusiones

Gracias a este trabajo podemos concluir que la programación distribuida nos permite mejorar la eficiencia de muchos algoritmos. En la actualidad, posiciones como data science y data analysis están en demanda y paradigmas como el big data están presentes cada vez en más empresas e instituciones. Según Youssra Riahi y Sara Riahi:
“Analytics tools are used when a company needs to do a forecasting and wants to know what will happen in the future, while BI tools help to transform those forecasts into common language.”

Asimismo, mencionan que el big data es considerado el sucesor de la inteligencia de negocios, por lo cual es un paradigma relevante durante nuestra carrera profesional.
Saber cómo plantear una arquitectura distribuida nos permitirá brindar soluciones más eficientes utilizando algoritmos que requieran mucho poder computacional ya que los podremos distribuir en distintas máquinas y se logrará una mejor utilización de los recursos disponibles para el procesamiento de la data.

### Bibliografía

- Van Steen, M., & S. Tanenbaum, A. (2016). A brief introduction to distributed systems. CrossMark. https://doi.org/10.1007/s00607-016-0508-7
- Tanenbaum, A., & Van Steen, M. (2006). Distributed Systems Principles and Paradigms (2.a ed.). Pearson.
- Morissette, L., & Chartier, S. (2013). The k-means clustering technique: General considerations and implementation in Mathematica. Tutorials in Quantitative Methods for Psychology, 9(1), 15–24. https://doi.org/10.20982/tqmp.09.1.p015
- Riahi, Y., & Riahi, S. (2018). Big Data and Big Data Analytics: Concepts, Types and Technologies. International Journal of Research and Engineering, 5(9), 524–528. https://doi.org/10.21276/ijre.2018.5.9.5

### Links

- Link al repositorio del backend:
https://github.com/JuanLeycal/TF_Concurrente

- Link al video exposición:
https://www.youtube.com/watch?v=QoMiT3ZnTP4

- Link al dataset (RAW):
https://raw.githubusercontent.com/JuanLeycal/TF_Concurrente/develop/DatasetSelectivo.csv
