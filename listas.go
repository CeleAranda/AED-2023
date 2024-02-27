//ejercicio 4.1
Diseñar un algoritmo para acceder, eliminar o insertar el k-ésimo elemento de una lista (siendo k una posición dada). Si la lista está vacía o si el valor de k esta fuera del rango del índice de la lista, invocar al procedimiento de ERROR. En cualquier otro caso, efectuar el procedimiento solicitado.

ACCION ejercicio41 (prim: puntero a NODO;) ES
	AMBIENTE
		NODO = Registro
			dato: entero;
			pos: entero;
			prox: puntero a NODO;
		Fin_registro;

		p,q,ant: puntero a NODO; //lista simplemente enlazada

		rta, cont, k: entero;
		

	PROCESO
		Escribir("¿Qué desea realizar?");
		Escribir("|1-Insertar| |2-Eliminar| |3-Acceder|");
		Leer(rta);

		Escribir("Ingrese en qué posición desear realizar la operación:");
		Leer(k);

		cont:=0;
		
		SI (rta = 3) ENTONCES: //ACCESOOOO
			p:= prim;
			MIENTRAS (p <> NIL) Y (cont < *p.pos) HACER:
				p:= *p.prox; //avanzo
				cont:= cont + 1; //incremento el contador para controlar con la cantiad de elementos
			Fin_mientras;

			SI (p = NIL) ENTONCES: //lista vacia
				error();
			SINO
				SI (k < cont) ENTONCES: //al no estar vacia la lista tengo que corroborar que la posicion ingresada no pase a los elementos que conté de la lista
					Escribir("El elemento de la posición",k,"es:",*p.dato);
				SINO
					error(); //la posicion ingresada es mayor al tamaño de la lista
				Fin_si;
			Fin_si;
		SINO
			SI (rta = 1) ENTONCES: //INSERCIONNNN ORDENADA
				NUEVO(q);
				SI (prim = NIL) ENTONCES:
					error() //lista vacia
				SINO
					p:= prim; //con p apunto donde apunta prim
					ant:= NIL; //resguardo nulo
					
					MIENTRAS (p <> NIL) Y (*p.pos < *q,pos) HACER: //mientras no desborde y la posicion de p sea menor, cuando es mayor no entra al ciclo, entonces ahi se debe insertar
						ant:= p; //resguardo p (anterior), funciona como las secuencias cuando hay que resguardar el caracter/entero anterior
						p:= *p.prox;
					Fin_mientras;

					SI (prim = NIL) ENTONCES:
						prim:= q;
						*q.prox:= p;
					SINO
						SI (*p.pos > *q.pos) ENTONCES: //desborde
							error();
						SINO
							*ant.prox:= q; //a lo que resguardé lo apunto a q (nuevo)
							*q.prox:= p; //al proximo de q (nuevo) lo apunto a p
						Fin_si;
					Fin_si;
			SINO
				//eliminar el nodo
				p:= prim;
				MIENTRAS (p <> NIL) Y (*p.pos <> k) HACER: //mientras p sea diferente de nil y diferente de la posicion ingresada, resguardar y avanzar
					ant:= p;
					p:= *p.prox;
				Fin_mientras;

				SI (p = NIL) ENTONCES:
					error(); //nada que eliminar, lista vacia
				SINO
					*ant.prox:= *p.prox; //voy al anterior (resguardo) y al prox de anterior lo apunto al prox de p (elimino el enlace entre ant y p)
					DISPONER(p);
				Fin_si;
			Fin_si;
		Fin_si;
Fin_accion.


//ejercicio 4.2
Se dispone de una lista simplemente encadenada de números enteros, diseñar un algoritmo que a partir de ella genere otra lista conteniendo los nodos cuyos datos terminan en cero; dichos elementos, deberán ser eliminados de la lista original. Se asume que la lista está cargada, y que el algoritmo recibe como parámetro de entrada la dirección del primer elemento.

ACCION ejercicio42 (prim1,prim2: puntero a NODO) ES
	AMBIENTE
		NODO = Registro
			dato: entero;
			prox: puntero a NODO;
		Fin_registro;

		p,q,ant: puntero a NODO;
		

	PROCESO
		p:= prim1 //no se por que puta tengo que apuntar con p a prim 
		prim2:= NIL; //vacia(?)

		MIENTRAS (p <> NIL) HACER:
			NUEVO(q); //pido lugar en memoria   
			*q.dato:= *p.dato; //copio el dato de p en q. Entiendo que creo un nuevo nodo y ahí empiezo a apuntar con los punteros de la lista2
			
			SI (p MOD 10 = 0) ENTONCES:
				SI (prim2 = NIL) ENTONCES: //creo que esto ocurre una vez, porque dps ya se carga la lista ES CON LISTA VACIAAA
					*q.prox:= p; //al proximo de q lo apunto a p
					prim2:= q; //mi primer elemento de la lista2 es q
				SINO 
					p:= prim2; //con p apunto al mismo que apunta prim2, o sea el primer elemento CARGO EL NODO ANTES DEL PRIMER ELEMENTO
					ant:= p; //resguardo mi p con ant

					MIENTRAS (p <> NIL) Y (*p.dato < *q.dato) HACER: //creo que esto es centinela 
						ant:= p; //resguardo mi p (anterior) y apunto al siguiente
						p:= *P.prox; //a p lo apunto a su siguiente (nodo siguiente celeste no es taN dificil)
					Fin_mientras;

					SI (p = prim) ENTONCES:
						prim:= q;  //cargo el primer elemento en la lista 2, creo
						*q.prox:= NIL;
					SINO
						*ant.prox:= q; //sino lo cargo al nodo intermedio, creo
						*q.prox:= p;
					Fin_si;
				Fin_si;

				p:= prim;

				SI (prim = NIL) ENTONCES:  //ELIMINO EL NODO
					prim:= *p.prox; //*p.prox tiene NIL creo
				SINO
					*ant.prox:= *p.prox;
				Fin_si;
				DISPONER(p); //libero la memoria
			SINO
				p:= *p.prox; //"avanzo"
			Fin_si;
		Fin_mientras;
Fin_accion.

//ejercicio 4.3
Dada una lista simplemente encadenada de números diseñar un algoritmo que calcule en forma independiente: La suma de los números impares, y la suma de los números pares.

ACCION ejercicio43(prim:puntero a NODO;) ES
	AMBIENTE 
		NODO = Registro
			num: real;
			prox: puntero a NODO;
		Fin_registro;

		p,ant: puntero a NODO;

		contp, conti: entero;

	PROCESO
		contp:= 0;
		conti:= 0;
		p:= prim; //apunto al primer nodo con p

		MIENTRAS (p <> NIL) HACER: //TODO va dentro de un mientras porque tengo que controlar que no desborde 
			SI (*p.dato MOD 2 = 0) ENTONCES:
				contp:= contp + *p.dato; //acumulo los valores 
			SINO
				conti:= conti + *p.dato;
			Fin_si;
			p:= *p.prox; //voy al siguiente nodo (avanzo)
		Fin_mientras;

		Escribir("La suma de números pares es:", contp);
		Escribir("La suma de números impares es:", conti);
Fin_accion.

//ejercicio 4.4
Se dispone de una lista simplemente encadenada cuyos registros están ordenados en forma ascendente por una clave de tipo entero; diseñar un algoritmo que invierta el o rden de la lista. //creo que tengo que apilar :p

ACCION ejercicio44 (prim: puntero a NODO;) ES
	AMBIENTE
		NODO = Registro
			clave: entero;
			prox: puntero a NODO;
		Fin_registro;

		p, q, ant, prim2: puntero a NODO;
	
	PROCESO
		prim2:= NIL; //lista nueva
		p:=prim; 

		MIENTRAS (p <> NIL) HACER:
			NUEVO(q); //asigno un lugar de memoria
			*q.calve:= *p.clave; //guardo el dato

			SI (prim2 = NIL) ENTONCES:
				*p.prox:= NIL; //primer elemento
			SINO
				*q.prox:= prim2; //apilo 
			Fin_si;

			prim2:= q;
			p:= *p.prox; //avanzo al siguiente con p
		Fin_mientras;

		Escribir("Lista invertida:");

		p:= prim2;

		MIENTRAS (p <> NIL) HACER: //muestro por pantalla la lista invertida
			Escribir(p);
			p:= *p.prox;
		Fin_mientras;
Fin_accion.


// #include <stdio.h>
// #include <stdlib.h>

// typedef struct nodo {
//   int numero;
//   struct nodo* prox;
// } nodo;

// void imprimir(nodo* inicio) {

//     nodo* actual = inicio;
//     int posicion = 0;

//     while (actual != NULL) {

//         printf("Nodo encontrado con valor \"%d\" en posicion: %d\n", actual->numero, posicion);

//         posicion++;
//         actual = actual->prox;

//     }
// }

// int main () {
//   nodo* prim = malloc(sizeof(struct nodo));
//   prim->numero = 0;

//   int listaSize = 5;
//   nodo* ultimo = prim;
  
//   for(int i = 1; i < listaSize; i++) {
//       nodo* nuevo = malloc(sizeof(struct nodo));
//       nuevo->numero = i;
//       ultimo->prox = nuevo;
//       ultimo = nuevo;
//   }

// imprimir(prim);

//   nodo* p = prim;
//   nodo* prim2 = NULL;
      
//   while (p != NULL) {
//       nodo* q = malloc(sizeof(struct nodo));
//       q->numero = p->numero;
//       q->prox = prim2; 
//       prim2 = q;
//       q = p;
//       p = p->prox;
//       free(q);
//   }
//   imprimir(prim2);
// }


//ejercicio4.5
Dada una lista simplemente encadenada que contiene datos de todas las provincias de la República Argentina: nombre, capital, cantidad total de habitantes y cantidad de analfabetos, y está ordenada en forma decreciente por número de habitantes analfabetos, generar otras tres listas que contengan el nombre, la capital y el porcentaje de analfabetos de las Provincias que respondan a las siguientes restricciones.

L1: <= 10 % analfabetos
L2: 16 a 25 % analfabetos
L3: => 26 % analfabetos

ACCION ejercicio45 (prim: puntero a NODO1;) ES
	AMBIENTE
		NODO1 = Registro
			nom:
			cap:
			cant_hab:
			cant_anal:		
			prox: puntero a NODO1	
		Fin_registro;

		p: puntero a NODO1;
		q,prim1,prim2,prim3: puntero a NODO2;

		NODO2 = Registro
			nom:
			cap:
			por:
			prox: puntero a NODO2;
		Fin_registro;

		porc: real;

	PROCESO
		p:=prim;
		prim1:= NIL;
		prim2:= NIL;
		prim3:= NIL;

		MIENTRAS (p <> NIL) HACER:
			porc:= *p.cant_anal DIV *p.cant_hab;  
			NUEVO(q);

			SI (porc <= 0,1) ENTONCES: //lista1
				*q.nom:= *p.nom;
				*q.cap:= *p.cap;
				*q.por:= porc;
				
				SI (prim1 = NIL) ENTONCES: //primer elemento, con lista vacía
					*p.prox:= NIL;
				SINO	
					*q.prox:= prim1; //apilar
				Fin_si;
			SINO 
				SI (porc >= 1,6) Y (porc <= 2,5) ENTONCES: //lista2
					*q.nom:= *p.nom;
					*q.cap:= *p.cap;
					*q.por:= porc;

					SI (prim2 = NIL) ENTONCES:
						*p.prox:= NIL;
					SINO
						*q.prox:= prim2;
					Fin_si;
				SINO
					*q.nom:= *p.nom; //lista3
					*q.cap:= *p.cap;
					*q.por:= porc;

					SI (prim3 = NIL) ENTONCES:
						*p.prox:= NIL;
					SINO	
						*q.prox:= prim3;
					Fin_si;
				Fin_si;
			Fin_si;
			*p.prox:= p;
		Fin_mientras;
Fin_accion.


//ejercicio4.6
En el restaurante ÑOQUIS se está pensando en una solución informática para el soporte de datos del nuevo sistema de atención a clientes. Se han decidido por LISTAS por su dinamismo en cuanto a la cantidad de elementos. Diseñe un algoritmo que realice las siguientes funciones:

Añadir cliente al ser atendido (lista simple ordenada por Nombre del Cliente).
Registrar su consumo (Acumular el Total Consumido en valores de montos).
Realizar el cobro (emitir ticket con Nombre, Fecha, Número de Mesa y Total).
Eliminar del listado de atención.
La información almacenada debe mantenerse ordenada por Nombre del cliente.

