# TF_Concurrente
Repositorio que contiene todo el TF del curso de Programación Concurrente.

### Diseño (arquitectura o componentes, etc.)

Para la arquitectura de este trabajo se planteó utilizar una distribución de nodos de tipo peer 2 peer, en la cual los nodos no tienen una jerarquía definida y el sistema se encuentra descentralizado. Asimismo, de forma más específica, nuestra arquitectura está basada en el Token Ring Algorithm o arquitectura de tipo anillo, la cual según Tanenbaum y Van Steen: “A completely different approach to deterministically achieving mutual exclusion...” (2006).

Esta arquitectura plantea que, el anillo, está compuesto de una cierta cantidad de nodos, a los cuales se le asigna un proceso y cada posición del anillo tiene una posición numérica que puede ser una dirección de red. El orden de cada nodo como tal no es relevante, lo único relevante en esta arquitectura es que el nodo actual sepa cual es el nodo sucesor, tal como mencionan Tanenbaum y Van Steen.

Para nuestro caso, hemos planteado la utilización de 5 nodos en una arquitectura de tipo anillo, los cuales utilizarán la técnica SPMD (Simple programa, múltiple data) ya que cada nodo tendrá acceso al mismo algoritmo de ML, pero cada nodo procesa un segmento de data distinto al resto de nodos.
