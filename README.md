# TF_Concurrente
Repositorio que contiene todo el TF del curso de Programación Concurrente.

### Introducción, descripción del problema y motivación

La definición de un “sistema distribuido” es la siguiente: Una colección de sistemas independientes que aparecen sus usuarios como una única computadora, lo cual les permite trabajar de manera concurrente. Los beneficios de tener este sistema varían según de qué manera están organizados los sistemas destinados a mejorar, ya sea sistemas centralizados, o sistemas aislados. Sin embargo, las ventajas generales de tener un sistema distribuido es poder escalar de manera horizontal tus sistemas, es más tolerante a fallos, más eficientes al procesar data,  poder implementar paralelismo, entre otros. 

Estos beneficios mencionados anteriormente nos pueden ayudar a resolver el problema que se describe a continuación. Existen varios algoritmos basados en machine learning que pueden analizar varias cantidades de datos para luego procesarla según el propósito del algoritmo de machine learning. Sin embargo, pueden haber casos donde los datos por analizar pueden llegar a ser bastante extensos, por lo cual el algoritmo desarrollado puede llegar a tardar demasiado en analizarlo, cosa que bajaría su rendimiento. Por esto en este trabajo, proponemos implementar un sistema distribuido a la API que contiene el algoritmo de machine learning, para que el proceso de este sea distribuido entre diferentes ordenadores, y así mejorar la eficiencia del algoritmo.

La mayoría de la motivación detrás de este proyecto es el hecho de aprender cómo analizar la implementación y el procesamiento de un sistema distribuido mediante este proyecto. Poder mejorar el rendimiento de un algoritmo sin necesidad de modificar el algoritmo en sí, sino su entorno. Finalmente, poder aplicar todos los principios de concurrencia, distribución y paralelismo enseñados en el curso en este trabajo.


### Diseño (arquitectura o componentes, etc.)

Para la arquitectura de este trabajo se planteó utilizar una distribución de nodos de tipo peer 2 peer, en la cual los nodos no tienen una jerarquía definida y el sistema se encuentra descentralizado. Asimismo, de forma más específica, nuestra arquitectura está basada en el Token Ring Algorithm o arquitectura de tipo anillo, la cual según Tanenbaum y Van Steen: “A completely different approach to deterministically achieving mutual exclusion...” (2006).

Esta arquitectura plantea que, el anillo, está compuesto de una cierta cantidad de nodos, a los cuales se le asigna un proceso y cada posición del anillo tiene una posición numérica que puede ser una dirección de red. El orden de cada nodo como tal no es relevante, lo único relevante en esta arquitectura es que el nodo actual sepa cual es el nodo sucesor, tal como mencionan Tanenbaum y Van Steen.

Para nuestro caso, hemos planteado la utilización de 5 nodos en una arquitectura de tipo anillo, los cuales utilizarán la técnica SPMD (Simple programa, múltiple data) ya que cada nodo tendrá acceso al mismo algoritmo de ML, pero cada nodo procesa un segmento de data distinto al resto de nodos.

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
