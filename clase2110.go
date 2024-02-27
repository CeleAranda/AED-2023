// clase 2110

//crear un nuevo NODO
NUEVO(Q)
LEER(*Q.DATO) //el usuario ingresa un dato 
*Q.DATO := valor //se le asigna el valor x

DISPONER(P) //libera el espaacio de memoria

//agregar NODO
PRIM := P;

NUEVO(Q);
PRIM := Q;
*Q.prox := NILL
PRIM := Q;

//Q.prox = puntero siguiente
//Q.ant = puntero anterior

//ELIMINAR
LEER (valor) 

q:= Prim
ant := NILL

	MIENTRAS (q <> nil) y (*q.dato <> valor) HACER
		ant:= q
		q:= *q.prox
	FIN MIENTRAS

	SI (q = nil) ENTONCES
		Escribir ('Error, el valor no existe')
	SINO
		SI (ant = nil) ENTONCES
			Prim = *q.prox
		SINO
			*ant.prox:= *q.prox //*se utiliza para moverse dentro del nodo
		FIN SI
		DISPONER(q)
	FIN SI






ACCION  ejercicio4_1 ES
	AMBIENTE
		k: entero;
		op: 1..3;
		NODO = Registro
			dato: entero;
			prox: puntero a NODO
		Fin_registro;
		p, ant, q, prim : puntero a NODO;

		PROCEDIMIENTO error ES		
				Escribir("Error");
		Fin_procedimiento;

		PROCEDIMIENTO eliminar ES	
			SI prim = NILL ENTONCES
				error();
			SINO	
				p := prim;
				ant := NILL;

				MIENTRAS p <> NILL y K > 1 HACER
					ant:= p;
					p:= *p.prox;
					k:= k-1;
				Fin_mientras;

				SI p=NILL ENTONCES
					error();
				SINO	
					SI p=prim ENTONCES 
						prim:= *p.prox;
					SINO
						*ant.prox:= *p.prox;
					Fin_si;
					DISPONER(p);

				Fin_si;
			Fin_si;
		Fin_procedimiento;

		PROCEDIMIENTO acceder ES
			p:= prim;
			SI p=prim ENTONCES
				error();
			SINO
				MIENTRAS p <> NILL y K > 1 HACER
					ant:= p;
					p:= *p.prox;
					k:= k-1; 
				Fin_mientras;
				SI p<> NILL ENTONCES	
					Escribir("El elemento en la posicion",k,"es", *p.dato);
				SINO
					error();
				Fin_si;
			Fin_si;
		Fin_procedimiento;

		PROCEDIMIENTO insertar() ES
			NUEVO(Q)
			Escribir("Ingrese u valor para la posicion K");
			LEER(*q.dato)

			SI prim = NILL ENTONCES
				prim := q;
				*q.prox := NILL
			SINO	
				p:= prim
				ant:= NILL
				MIENTRAS p <> NILL y K > 1 HACER
					ant:= p;
					p:= *p.prox;
					k:= k-1; 
				Fin_mientras;
				SI p = prim ENTONCES
					*q.prox:= p;
					prim:= q;
				SINO	
					*ant.prox:= q;
					*q.prox:= p;
				Fin_si;
			Fin_si;
		Fin_procedimiento;

		PROCESO
			Escribir("Ingrese el valor de la posicion K: ")
			LEER(k)

			Escribir("Seleccione una de las opciones: ")
			Escribir("1_Eliminar")
			Escribir("2_Acceder")
			Escribir("3_Insertar")
			LEER(op)

			SEGUN op HACER	
				1: eliminar();
				2: acceder();
				3: insertar();
			Fin_segun;
Fin_accion


ACCION ejercicio4.3 ES
	AMBIENTE
	NODO = Registro
		dato: entero;
		prox: puntero a NODO
	Fin_registro;
	p, ant, q, prim : puntero a NODO;

	cont_im, cont_par: entero;

	PROCESO	
		p:=prim;
		cont_im:= 0;
		cont_par:= 0;

		MIENTRAS p<> NILL HACER
			SI *p.dato MOD 2 = 0 ENTONCES
				cont_par:= cont_par + *p.dato
			SINO
				cont_im:= cont_im + *p.dato
			Fin_si;
			p:= *p.prox;
		Fin_mientras;

		Escribir("La sauma de números pares es:", cont_par);
		Escribir("La sauma de números pares es:", cont_im);
Fin_accion;

ACCION ejercicio4_2