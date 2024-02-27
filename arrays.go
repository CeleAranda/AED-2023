//Ejercicio 3.1 
Para cada una de las consignas siguientes, genere un algoritmo que permita solucionarla (3 algoritmos) usando un arreglo de 100 números enteros:
1-Almacenar 100 números.
2-Localizar el número de mayor valor y el de menor valor, informar sus valores y sus posiciones.
3-Contar y sumar todos los números pares.

ACCION ejercicio1 ES 
	AMBIENTE
		arr: arreglo [1..100] de enteros;
		i: entero;

	PROCESO
		i:=1..100;
		PARA i:=1 a 100 HACER:
			Escribir("Ingrese un valor para la posición -",i,":");
			Leer(arr[i]);
		Fin_para;
		
Fin_accion.


ACCION ejercicio2(arr: arreglo [1..100] de enteros) ES
	AMBIENTE
		mayor, menor: entero;
		i:= 1..100;
		pos_may: 1..100;
		pos_men: 1..100;

	PROCESO
		mayor := LV; //LV menor valor que genera la máquina
		menor := HV; //HV mayor valor que genera la máquina

		PARA i:=1 a 100 HACER:	
			SI (arr[i] > mayor) ENTONCES
				mayor:= arr[i];
				pos_may:= i;
			SINO
				SI (arr[i] < menor) ENTONCES
					menor:= arr[i];
					pos_men:= i;	
			Fin_si;
		Fin_para;

		//informo cuando recorrí todo el arreglo cuál elemento es mayor y menor

		Escribir("El elemento mayor es:",mayor,"y está en la posición -",pos_may,".");
		Escribir("El elemento menor es:",menor,"y está en la posición -",pos_men,".");

Fin_accion.


ACCION ejercicio3(arr: arreglo [1..100] de enteros) ES
	AMBIENTE
		cont, sum: entero;
		i:= 1..100;

	PROCESO
		cont:= 0;
		sum:= 0;

		PARA i:=1 a 100 HACER:
			SI (arr[i] MOD 2 = 0) ENTONCES:
				cont:= cont + 1;
				sum:= sum + arr[i];
			Fin_si;
		Fin_para;

		Escribir("La cantidad de números pares es:",cont,"y la suma de ellos es:",sum);
Fin_accion.
			


//Ejercicio3.2
Genere un único algoritmo que resuelva las 3 consignas del ejercicio anterior.

ACCION ejercicio3.2 ES
	AMBIENTE
		arr: arreglo [1..100] de entero;
		i: 1..100;
		pos_may: 1..100;
		pos_men: 1..100;
		cont, sum, mayor, menor: entero;
	
		PROCEDIMIENTO ingresar_datos():
			PARA i:=1 a 100 HACER
				Escribir("Ingrese un valor para la posición -",i,":");
				Leer(arr[i]);
			Fin_para;
		Fin_proced;


	PROCESO
		cont:= 0;
		sum:= 0;
		mayor:= LV;
		menor:= HV;

		ingresar_datos(); //cargo el arreglo

		PARA i:=1 a 100 HACER:
			SI(arr[i] MOD 2 = 0) ENTONCES:   //verifico si el elemento es par o no
				cont:= cont + 1;
				sum:= sum + arr[i];
			Fin_si;

			SI (arr[i] > mayor) ENTONCES:   //verifico cuál elemento es mayor y menor
				mayor:= arr[i];
				pos_may:= i;
			SINO
				SI (arr[i] < menor) ENTONCES:
					menor:= arr[i];
					pos_men:= i;
			Fin_si;
		Fin_para;

		Escribir("La cantidad de elementos pares es de:", cont,"y la suma de ellos es:", sum);
		Escribir("El elemento mayor es:",mayor,"y se encuentra en la posición -",pos_may);
		Escribir("El elemento menor es:",menor,"y se encuentra en la posición -",pos_men);
Fin_accion.



//Ejercicio 3.3
Dados un vector de reales a de n elementos, con n ∈ N0, y un número real x >= 0, calcular el valor y tal que:
y=i=1 n ∑(ai × xi)

ACCION ejercicio3.3 (arr: arreglo [0..n] de reales) ES
	AMBIENTE
		i: 0..n;
		x, result: real;

	PROCESO
		result:= 0;

		PARA i:= 1..n HACER:
			Escribir("Ingrese un valor real para el número x, debe ser mayor o igual a 0:");
			Leer(x);
			result:= result + (arr[i] * x**i);
		Fin_para; 
Fin_accion.



//Ejercicio 3.4
Considerando un arreglo de 50 números enteros, confeccione un algoritmo para resolver las siguientes consignas:

1) Modificar el arreglo dado, de modo que todos sus elementos sean múltiplos de 3.
2) Crear otro arreglo que contenga los números que no cumplieron la condición.
3) Informar cuántos números cumplieron la condición.

ACCION ejercicio 3.4 (arr: arreglo [1..50] de entero) ES
	AMBIENTE
		arr2: arr [1..50] de entero;
		cont: entero;
		i: 1..50;
	
	PROCESO
		cont:= 0;

		PARA i:=1 a 50 HACER:
			SI (arr[i] MOD 3 <> 0) ENTONCES:
				arr2[i]:= arr[i];   //guardo los que no son múltiplos en otro arreglo
				arr[i]:= arr[i] * 3;  //modifico el arreglo de entrada para que sus elementos sean múltiplos de 3
			SINO
				cont:= cont + 1;  //cuento los que si son múltiplos
			Fin_si;
		Fin_para;

		Escribir("La cantidad de números que sí son múltiplos de 3 es:", cont);
Fin_accion.



//Ejercicio 3.5
Dados 2 vectores:
A: arreglo [1 .. 30] de reales; B: arreglo [1 .. 30] de reales.
Ambos ordenados de forma creciente, escribir un algoritmo que realice la mezcla de ambos para obtener otro vector tambien ordenado de forma creciente:
C: arreglo [1 .. 60] de reales.

//se usa el ciclo excluyente porque puede que un arreglo no esté totalmente cargado, entonces si termina uno antes que el otro no se van a copiar todos lo elementos

ACCION ejercicio3.5 (A: arreglo [1 .. 30] de reales; B: arreglo [1 .. 30] de reales) ES
	AMBIENTE
		C: arreglo [1 .. 60] de reales;
		i: 1..60;
		j, k: 1..30;

	PROCESO
		MIENTRAS (j < 30 Y k < 30) HACER:  //el 30 "sería mi HV", o sea mientras no termine el arreglo A y B
			SI (A[j] < B[k]) ENTONCES:
				C[i]:= A[j];
				j:= j + 1;
			SINO
				C[i]:= B[k];
				k:= k + 1;
			Fin_si;
			i:= i + 1; //aumento el índice de C para que continúe el ciclo
		Fin_mientras;
		
		//acá copio los elementos que faltan si uno terimó antes
		MIENTRAS (j < 30):
			C[i]:= A[j];
			j:= j + 1;
			i:= i + 1;
		Fin_mientras;

		MIENTRAS (k < 30):
			C[i]:= B[k];
			k:= k + 1;
			i:= i + 1;
		Fin_mientras;

		PARA i:=1 a 60 HACER:
			Escribir("Arreglo C ordenado de manera creciente:");
			Escribir(C[i]);
		Fin_para;
Fin_accion.



//Ejercicio 3.6
Escribir un algoritmo que permita cargar un arreglo de N nombres, considerando que cada nombre debe tener entre 1 y 10 caracteres.

ACCION ejercicio 3.6 ES
	AMBIENTE
		arr: arrgelo (1..N, 1..10) de caracter;
		i: 1..N;
		j; 1..10;

	PROCESO
		PARA i:=1 a N HACER:
			Escribir("Ingrese el nombre",i);
			PARA j:=1 a 10 HACER:
				Escribir("Ingrese el caracter de la posición -",j,"del nombre:");
				Leer(arr[j]);
			Fin_para;
		Fin_para;

		PARA i:=1 a N HACER:
			Escribir("Arreglo cargado con nombres:");
			Escribir(arr[i]);
		Fin_para;
Fin_accion.



//Ejercicio 3.7 ¿¿¿???
Escribir un algoritmo que permita localizar un nombre en un arreglo de N nombres, ordenados alfabéticamente. Cada nombre puede tener, como máximo, 10 caracteres. Escriba por lo menos dos algoritmos que permitan solucionar el problema; especifique cuál de las formas considera más eficiente y por qué.

ACCION ejercicio3.7 ES
	AMBIENTE
		arr: arreglo [1..N, 1..10] de caracter;
		i: 1..N;
		j, k: 1..10;
		nom: arreglo [1..10] de caracter;

	PROCESO
		PARA k:=1 a 10 HACER:
			Escribir("Ingrese un nombre caracter a caracter, como máximo 10:");
			Leer(nom[k]);
		Fin_para;

		PARA i:=1 a N HACER:
			PARA j:=1 a 10 HACER:
Fin_accion.



//Ejercicio 3.8
Repita el ejercicio anterior, pero suponiendo que se precisa localizar todos los nombres que comienzan con una letra dada.



//Ejercicio 3.9
Se posee un arreglo de 200 libros con el siguiente formato:
NRO_LIBRO	TITULO	AUTOR	CANT_HOJAS

Ordenado por AUTOR y se presentan las siguientes premisas:
1-Se necesita saber que libros se poseen de “Nicklaus Wirth”.
2-Se necesita saber en qué posición se encuentra “Algoritmos + Estructuras de Datos=Programa”.
3-Se necesita saber cual es el libro de “Nicklaus Wirth” de mayor volumen.

ACCION ejercicio3.9 (arr: arreglo [1..200] de Libros) ES
	AMBIENTE
		Libros = Registro
			NRO_LIBRO: N(8);
			TITULO: AN(50);
			AUTOR: AN(50);
			CANT_HOJAS: N(9);
		Fin_registro;

		i: 1..200;
		mayor_vol: entero;
		titulo_may: caracter;

	PROCESO
		mayor_vol:= LV;

		PARA i:=1 a 200 HACER:
			SI (arr[i].AUTOR = "Nicklaus Wirth") ENTONCES:
				Escribir(arr[i]);  //muestro por pantalla los libros del autor ese

				SI (arr[i].CANT_HOJAS > mayor_vol) ENTONCES:
					mayor_vol:= arr[i].CANT_HOJAS;  //verifica cuál tiene mayor cantidad de hojas
					titulo_may:= arr[i].TITULO;
				Fin_si;
				
				SI (arr[i].TITULO = "Algoritmos + Estructuras de Datos=Programa") ENTONCES:
					Escribir("El libro: 'Algoritmos + Estructuras de Datos=Programa' se encuentra en la posición:",i);
				Fin_si;
			Fin_si;
		Fin_para;

		Escribir("El libro con mayor volumen es:", titulo_may);
Fin_accion.



//Ejercicio 3.10
Dado un arreglo de 50 elementos, cada uno de los cuales tiene los siguientes datos: Código de localidad y Lluvia caída en un año. Escribir un algoritmo que permita saber dada una localidad, cuanto llovió en el año. Aplicar el método más adecuado considerando que el arreglo esta ordenado por Código de localidad.

ACCION ejercicio 3.10 (arr: arreglo [1..50] de Lluvia) ES
	AMBIENTE
		Lluvia = Registro
			cdo_loc: N(8);
			cant_lluv: N(9,3);
		Fin_registro;

		loc: N(8);
		i: 1..50;
	
	PROCESO
		Escribir("Ingrese el código de una localidad:");
		Leer(loc);

		i:= 1;  //inicializo el índice antes del ciclo

		MIENTRAS (i < 50) Y (arr[i].cdo_loc <> loc) HACER:  //método de búsqueda: lineal con centinela 
			i:= i + 1;  //avanzo el arreglo
		Fin_mientras;
		
		SI (arr[i].cdo_loc = loc) ENTONCES:
			Escribir("En la localidad",loc,"llovió en el año:", arr[i].cant_lluv);
		SINO
			Escribir("No se encontró la localidad ingresada.")
		Fin_si;
Fin_accion.



//Ejercicio 3.11
Dado un arreglo de 100 elementos, que contiene la siguiente información sobre videos: Título de la película, Nombre del Director, Categoría de película, Cantidad de personas que la vieron, Alquilado (si/no); y está ordenado por el Título de la película, diseñe un algoritmo que, ingresando una categoría, liste todas las películas que pertenecen a dicha categoría.

ACCION ejercicio3.11 (arr: arreglo [1..100] de Pelis) ES
	AMBIENTE
		Pelis = Registro
			titulo: AN(40);
			nom_dir: AN(40);
			categ: AN(10);
			cant_pers: N(9);
			alquilado: {"si","no"};
		Fin_registro;

		cat: AN(10);
		i: 1..100;

	PROCESO	
		Escribir("Ingrese la categoria de película de la que desea obtener un listado:");
		Leer(cat);

		Escribir("Listado:")
		
		PARA i:=1 a 100 HACER:
			SI (arr[i].categ = cat) ENTONCES:
				Escribir(arr[i]);
			Fin_si;
		Fin_para;
Fin_accion.



//Ejercicio 3.12
A partir del arreglo de videos descrito en el ejercicio anterior, diseñe un algoritmo que permita atender un pedido de alquiler, para lo cual debe verificar si es posible o no y, cuando corresponda, actualizar la cantidad de personas que vieron dicha película.

ACCION ejercicio3.12 (arr: arreglo [1..100] de Pelis) ES
	AMBIENTE
		Pelis = Registro
			titulo: AN(40);
			nom_dir: AN(40);
			categ: AN(10);
			cant_pers: N(9);
			alquilado: {"si","no"};
		Fin_registro;

		peli_us: AN(40);
		i: 1..100;

	PROCESO
		Escribir("Ingrese el título de la película que desea alquilar:");
		Leer(peli_us);

		PARA i:=1 a 100 HACER:
			SI (arr[i].titulo = peli_us) ENTONCES:
				SI (arr[i].alquilado = "no") ENTONCES:
					Escribir("Es posible alquilar esta película.");
					arr[i].cant_pers:= arr[i].cant_pers + 1;
				SINO
					Escribir("No es posible alquilar esta película.");
				Fin_si;
			SINO
				Escribir("No se ha encontrado la película");
			Fin_si;
		Fin_para;
Fin_accion.



//Ejercicio 3.13
Se precisa ordenar un arreglo de N alumnos de mayor a menor, de acuerdo a la cantidad de materias aprobadas. Cada elemento del arreglo contiene Nro. de Legajo y Cantidad de materias aprobadas. Escriba por lo menos dos algoritmos que permitan solucionar el problema; especifique cuál de las formas considera más eficiente y por qué.

//se ordena con selección directa, por qué? ni idea, supongo que es porque es el método más rápido (creo)

ACCION ejercicio3.13 (arr: arreglo [1..N] de Alumnos) ES
	AMBIENTE
		Alumnos = Registro
			legajo: N(8);
			cant_mat: N(2);
		Fin_registro;

		x,max: entero;
		i,j: 1..N;

	PROCESO
		//ordena de mayor a menor
		PARA i:= 1 a N-1 HACER:
			x:= arr[i];
			max:= i  //max guarda la posicion

			PARA j:= i+1 a N HACER
				SI (x < arr[j]);
					max:= j;
					x:= arr[j];
				Fin_si;
			Fin_para;

			arr[max]:= arr[i];
			arr[i]:= x;
		Fin_para;
Fin_accion.



//Ejercicio 3.14
El mes que viene se realizará en Buenos Aires el desfile “Alta Moda 2018”, el cual reúne a los diseñadores más reconocidos del país. Para organizar las pasadas, se dispone de un arreglo por diseñador, con la siguiente información: nombre de el/la modelo y altura. La directora del evento necesita que se imprima cada lista ordenada de acuerdo a la altura de el/la modelo. Escribir un algoritmo que permita ingresar los datos de cada diseñador e imprimirlos de acuerdo a lo solicitado.

//no seee como ordenarrr aaa

ACCION ejercicio3.14 ES
	AMBIENTE
		arr: arreglo [1..N] de Diseñador;
		i:entero

		Diseñador = Registro
			nom_mod: AN(40);
			altura: N(2,3);
		Fin_registro;

		PROCEDIMIENTO carga_datos():
			Escribir("Ingrese el nombre el/la modelo:");
			Leer(arr[i].nom_mod);
			Escribir("Ingrese la altura:");
			Leer(arr[i].altura);
		Fin_proced;

	PROCESO
		PARA i:=1 a N HACER:
			carga_datos();

			

//Ejercicio 3.15
Se precisa ordenar un arreglo de enteros de menor a mayor, eliminando los números repetidos.

ACCION ejercicop3.15 ES
	AMBIENTE
		arr: arreglo de [1..N] de entero;
		i, j, x: entero;  //x actúa como auxiliar

		PROCEDIMIENTO carga() ES
			PARA i:=1 a N HACER:
				Esc("ingrese un numero:");
				Leer(Arr[i]);
			Fin_para;
		Fin_proced;

	PROCESO
		carga(); //10 4 56 56 8 8 9

		PARA i:=2 a N HACER:
			x:=arr[i]; //x = 4 (es el de la posicion 2)
			j:=i-1; //j = 1
			MIENTRAS (j > 0) y (x <= arr[j]) HACER: //arr[j] = 10 (es el de la posicion 1)
				arr[j+1]:= arr[j];  //arr[j+1] (en la posicion 2) = 10 
				j:= j-1;  //j = 0 (?)
			Fin_mientras;
			arr[j+1]:= x;  //arr[j+1] (posicion 1) = 4
		Fin_para;
		
		//4 8 8 9 10 56 56

		PARA i:=1 a N HACER:
			x:= arr[i];  //x = 4
			PARA j:=i+1 a N HACER: //j = 2
				SI (arr[j] = arr[i]) ENTONCES:  //verifico si el de la segunda posicion es igual a de la primera posicion 
					arr[j]:= null;
				Fin_si;
			Fin_para;
		Fin_para;
Fin_accion.



//Ejercicio3.16
Una empresa que comercializa una cierta cantidad de artículos diferentes desea confeccionar un ranking de ventas de los mismos, a partir de una secuencia ordenada por articulo que contiene:
|NRO_ARTICULO|	|TIPO|	|CANT_VENDIDA|
Escribir un algoritmo que emita el ranking deseado en orden decreciente por cantidad.
//met de ordenamiento -> seleccion directa, ordena de mayor a menor

ACCION ejercicio3.16 (arr: arreglo [1..N] de Articulos) ES
	AMBIENTE
		Articulos = Registro
			nro_art: N(8);
			tipo: AN(8);
			cant_vend: N(9);
		Fin_registro;

		i, j, x:entero;

	PROCESO
		PARA i:= 1 a N HACER:  
			x:= arr[i].cant_vend;  //guardo datos
			max:= i;  //guardo posicion 
			PARA j:= i+1 a N HACER:
				SI (x < arr[j].cant_vend) ENTONCES:  //arr[j].cant_vend es el elemento siguiente
					max:= j; //guardo la posicion del mayor 
					x:= arr[j].cant_vend; //guardo el dato mayor
				Fin_si;
			Fin_para;
			arr[max]:= arr[i].cant_vend;
			arr[i]:= x;
		Fin_para;

		PARA i:=1 a N HACER:
			Escribir("Ranking de ventas:");
			Escribir("El artículo:",arr[i].nro_art,"tiene el ranking:",i);
		Fin_para;
Fin_accion.



//Ejercicio3.17
Se precisa generar una secuencia con los datos de los 10 videos más vistos de una categoría, a partir de una secuencia de entrada de 200 registros que contiene el Título de la película, Nombre del Director, Categoría de película, Cantidad de personas que la vieron y que está ordenada por el Título de la película.

//puedo suponer que el arreglo de 200 registros está cargado y lo paso como parámetro (o no, y cargarlo dentro del programa).Defino 2 arreglos más, el primero (arreglo B) con la misma cantidad de elementos que el arreglo de entrada y el segundo (arreglo C) con 10 elementos (van a ser los videos más vistos). En B guardo todas las pelis de la categoria ingresada por el usuario, luego ordeno con intercambio directo (mayor a menor) y guardo los 10 primeros elementos en C

ACCION ejercicio3.17 (A: arreglo [1..200] de Videos) ES
	AMBIENTE
		Videos = Registro
			tit_peli: AN(30);
			nom_dir: AN(30);
			categ: AN(2);
			cant_pers: N(9);
		Fin_registro;
		
		B: arreglo [1..200] de Videos;
		C: arreglo [1..10] de Videos;

		i, j, x: entero;
		categ_us: AN(2);
		bandera: booleano;
		
		
	PROCESO
		Escribir("Ingrese la categoría de la que quiere listar los 10 videos más vistos:");
		Leer(categ_us);
		j:= 0;
		
		PARA i:=1 a 200 HACER:
			SI (A[i].categ = categ_us) ENTONCES
				j:= j + 1;
				B[j]:= A[i]; //copio el dato de A en B si las categorias son iguales
			Fin_si;
		Fin_para;

		bandera:= Falso;

		MIENTRAS NO bandera HACER: //intercambio directo
			bandera:= Verdadero;
			PARA i:=1 a j-1 HACER:
				SI (B[i].cant_pers < B[i+1].cant_pers) ENTONCES:
					x:= B[i]; //x es aux, guarda el dato de B[i] 
					B[i]:= B[i+1]; //a B[i] le asigno el dato de la posicion siguiente
					B[i+1]:= x; //y al siguiente le asigno el anterior, lo que guardé en x
					bandera:= Falso;
				Fin_si;
			Fin_para;
		Fin_mientras;

		PARA i:=1 a 10 HACER:
			C[i]:= B[i];  //guardo en el arreglo C los datos de B
		Fin_para;

		Escribir("Lista de los 10 videos más vistos de la categoría:", categ_us);
		PARA i:=1 a 10 HACER:
			Escribir(C[i]);
		Fin_para;
Fin_accion.



//Ejercicio3.18  ??¡?¡??¡??¡
Se precisa diseñar una agenda electrónica, donde se archivará el nombre, la dirección y el teléfono de hasta 50 personas. Diseñe un algoritmo que permita efectuar consultas, modificaciones, eliminaciones de los datos de una persona y agregados de nuevas personas (sólo será posible incorporar una persona si hay menos de 50 archivadas en la agenda). Los datos se hallan almacenados en una secuencia. Considere que la agenda siempre debe mantenerse ordenada alfabéticamente.

ACCION ejercicio3.18 Es  
	AMBIENTE
		persona = Registro
			nom: AN(30);
			direc: AN(40);
			tel: N(8);
		Fin_registro;

		arr: arreglo [1..50] de persona;
		i: 1..50;
Fin_accion



//Ejercicio3.19
Escribir un algoritmo que emita cuál es y dónde está ubicado el mayor elemento de cada fila de una matriz, e imprima un mensaje si todos los mayores se encuentran en la misma columna.

ACCION ejercicio3.19 (arr: arreglo[1..N, 1..N] de entero) ES
	AMBIENTE 
		i: 1..N; //filas
		j: 1..N; //columnas
		resmay_f: entero;
		resg_c, resg_c2: entero;
		bandera: booleano;

	PROCESO
		bandera: Verdadero;

		PARA i:=1 a N HACER:
			resmay_f:= arr[i, 1]; //recorro las filas

			PARA j:= 1 a 3 HACER:
				SI (resmay_f < arr[i,j]) ENTONCES //si mi resguardo de fila es mayor que los datos que siguen (en la misma fila)
					resmay_f:= arr[i,j];
					resg_c:= j //resguardo la columna
				Fin_si;
			Fin_para;

			Escribir("El mayor número encontrado en la fila",i,"y columna",resg_c,"es", resmay_f);

			SI (i=1) ENTONCES:
				resg_c2:= resg_c; 
			Fin_si;

			SI (i<>1) ENTONCES:
				si (resg_c <> resg_c2) ENTONCES: //si no entra acá es porque todos los números mayores están en la misma columna que el primero
					bandera:= Falso;
				Fin_si;
			Fin_si;

			SI bandera ENTONCES:
				Escribir("Los números mayores están en la misma columna.");
			SINO
			 	Escribir("Los números mayores no están en la misma columna.");
			Fin_si;
		Fin_para;
Fin_accion.

"""
	j j j
i	1 2 4 i=1
i	3 1 3 i=2
i	5 2 5 i=3
i	4 7 8 i=4

resmayor = 4 (primera fila) 
rc= 3 (tercer columna)

resmayor = 3 (segunda fila)
rc= 1 (primera columna)

resmayor = 5 (tercera fila )
rc= 1 (primera columna)

resmayor= 8 (cuarta fila)
rc= 3 (tercer columna)
"""



//Ejercicio3.20
Dadas dos matrices A y B, cuadradas, de 5 x 5, con números enteros, cargar una matriz C, de 5 x 5 teniendo en cuenta las siguientes condiciones: la diagonal principal completar con ceros, en las posiciones que están por encima de la diagonal principal, copiar los correspondientes elementos de la matriz A y en las posiciones que están por debajo de la diagonal principal, copiar los elementos correspondientes de la matriz B.

ACCION ejercicio3.20 (A: arreglo [1..5,1..5] de entero; B: arreglo [1..5,1..5] de entero) ES
	AMBIENTE
		C: arreglo [1..5,1..5] de entero;
		i,j: entero;
	
	PROCESO
		PARA i:=1 a 5 HACER:
			PARA j:=1 a 5 HACER:
				SI (i=j) ENTONCES:  //diagonal principal
					C[i,j]:= 0;  
				SINO
					SI (i<j) ENTONCES:  //por encima
						C[i,j]:= A[i,j];
					SINO
						C[i,j]:= B[i,j];  //por debajo
					Fin_si;
				Fin_si;
			Fin_para;
		Fin_para;
Fin_accion.

"""
    A            B           C
2 3 7 4 5    3 4 6 7 8   0 3 7 4 5
2 3 0 5 6    3 4 7 8 0   3 0 0 5 6
3 6 7 4 3    2 3 2 9 9   3 6 0 4 3
4 6 7 2 2    3 2 3 4 5   4 6 7 0 2
5 1 3 4 2    1 2 6 5 7   5 1 3 4 0

diagonal principal -> i=j
posiciones por encima de la diagonal -> i<j
posiciones por debajo de la diagonal -> i>j
"""



//Ejercicio3.21 
Dada una matriz de 6 x 6 de enteros, cuya última fila y columna contienen ceros, calcular la suma de cada fila y guardar en la última celda de la misma; y la suma de cada columna y guardar en la última celda de la misma. Calcular también el total general y guardar en la posición (6,6).

ACCION ejercicio3.21 (m: arreglo [1..6,1..6] de entero) ES
	AMBIENTE
		i,j: entero;
		totf, totc, totg: entero;

	PROCESO
		PARA i:=1 a 5 HACER:
			PARA j:=1 a 5 HACER:
				m[6,j]:= m[6,j] + m[i,j];  //acumula por columna en la última celda
				m[i,6]:= m[i,6] + m[i,j];  //acumula por filas en la última celda
			Fin_para;
		Fin_para;

		m[6,6]:= m[6,j] + m[i,6];

		Escribir("El total gener es:", m[6,6]);
Fin_accion.

"""
	j j j j j j 
i	2 3 7 4 5 0   
i	2 3 0 5 6 0   
i	3 6 7 4 3 0   
i	4 6 7 2 2 0   
i	5 1 3 4 2 0
i	0 0 0 0 0 0

//acumulando en la última celda de cada columna 
i= 2 , j= 1, m[2,1]=2
m[6,1]= 2+2

i= 2 , j= 2, m[2,2]=3
m[6,2]= 3+3

i= 2 , j= 3, m[2,3]=0
m[6,3]= 7+0

i= 2 , j= 4, m[2,4]=5
m[6,4]= 4+5

i= 2 , j= 5, m[2,5]=6
m[6,5]= 5+6
"""

//Ejercicio3.22
Dada una matriz cuadrada de 5 x 5 de números, sumar filas y columnas y guardar en una matriz de 2 x 5, de modo que la fila 1 contenga la suma de cada fila y la fila 2, la suma de cada columna.

ACCION ejercicio3.22 (m: arreglo [1..5,1..5] de real) ES
	AMBIENTE
		a: arreglo [1..2,1..5] de real;
		i,j: entero;
		k,l: entero;
	
	PROCESO
		PARA k:=1 a 2 HACER:
			PARA l:=1 a 5 HACER:
				a[k,l]:= 0;  //pongo a cero la matriz a
			Fin_para;
		Fin_para;

		PARA i:=1 a 5 HACER:
			PARA j:=1 a 5 HACER:
				a[1,j]:= a[1,j] + m[i,j]; //acumulo los numeros de la final en la fila 1 de a
				a[2,j]:= a[2,j] + m[i,j]; //acumulo los numeros de la columna en la fila 2 de a
			Fin_para;
		Fin_para;
Fin_accion.

"""
2 4 5 6 5    0 0 0 0 0  -> suma de cada fila 
4 2 1 3 4    0 0 0 0 0  -> suma de cada columna
5 4 3 4 2
1 2 3 2 3
2 3 4 1 2
"""



//Ejercicio3.23
En un sector de un hospital, donde se encuentran internados 50 pacientes, se toma la temperatura de cada paciente 4 veces al día durante una semana.
Lec	Dom Lun Mar	Mie Jue Vie Sab
1	XX	XX	XX	XX	XX	XX	XX  
2	XX	XX	XX	XX	XX	XX	XX 
3	XX	XX	XX	XX	XX	XX	XX 
4	XX	XX	XX	XX	XX	XX  XX  
Se dispone de un arreglo con la información recopilada de todos los pacientes.

Construir un algoritmo que:
Liste por pantalla las temperaturas máxima y mínima de cada paciente, indicando el día y lectura en la que ocurrieron.
Genere un nuevo arreglo que contenga la temperatura promedio por día de cada paciente.

ACCION ejercicio3.23 (dato: arreglo [1..4,1..7,1..50] de real) ES
	AMBIENTE
		i,j,k: entero;

		tem: arreglo [1..7,1..50] de real;
		l,m: entero;

		temp_max, temp_min: real;
		dia_max, dia_in: entero;
		lect_max, lect_min: entero;
		tot:entero;

	PROCESO
		PARA l:=1 a 7 HACER:
			PARA m:=1 a 50 HACER:
				tem[l,m]:= 0;  //pongo a 0 el arreglo de salida
			Fin_para;
		Fin_para;

		PARA k:=1 a 50 HACER:
			temp_max:= LV; 
			temp_min:= HV;

			PARA j:=1 a 7 HACER: //recorro por columnas para saber el promedio por dia
				tot:=0; //acumula las temperaturas, pongo en cero para la temperatura promedio de cada dia
				
				PARA i:=1 a 4 HACER:
					SI (dato[i,j,k] > temp_max) ENTONCES:
						temp_max:= dato[i,j,k];  //guardo la mayor temperatura de UN paciente
						dia_max:= j; //resguardo el dia
						lect_max:= i;  //resguardo el numero de lectura
					Fin_si;	
				
					SI(dato[i,j,k] < temp_min) ENTONCES:
						temp_min:= dato[i,j,k];  //guardo la menor temp
						dia_min:= j;  //resguardo el dia
						lect_min:= i;  //resguardo el numero de lectura
					Fin_si;

					tot:= tot + dato[i,j,k]; //guardo la temperatura de cada dia
				
				Fin_para;

				Escribir("La temperatura promedio del día",j,"es",tot/4);
				tem[j,k]:= tot/4; //guardo en el arreglo de salida la temperatura promedio del paciente k y dia j
	 		
			Fin_para;

			Escribir("Paciente:",k);
			Escribir("Temperatura máxima:",temp_max,", en la lectura:",lect_max,", el día:",dia_max);
			Escribir("Temperatura mínima:",temp_min,", en la lectura:",lect_min,", el día:",dia_min);
		
		Fin_para;
Fin_accion.



//Ejercicio3.24
Se cuenta con información acerca de los cajeros de un Supermercado, el cual se halla estructurado en 10 cajas registradoras. Dicha información está registrada en una secuencia que contiene Apellido y Nombre del empleado, número de caja que tiene asignada, importe facturado y horario de facturación; la secuencia está ordenada alfabéticamente por Apellido y Nombre. Se solicita una estadística de los importes facturados, discriminado por número de caja y franja de horas y además los montos totales, según el siguiente formato:

Cajas	8-10	10-12	12-16	16-18	18-20	Total x Cajas
1	     ...	 ...	 ...	 ...	 ...	 ...
..	     ...	 ...	 ...	 ...	 ...     ...
10	     ...	 ...	 ...	 ...	 ...	 ...
Total    ...     ...     ...     ...     ...     ... 
x horas	

ACCION ejercicio3.24 ES
	AMBIENTE
		Cajeros = Registro
			apelynom: AN(30);
			nro_caja: N(8);
			importe: N(8,2);
			hora: N(2,2);
		Fin_registro;

		Arch: archivo de Cajeros ordenado por apelynom;
		reg: Cajeros;

		arr: arreglo [1..11,1..6] de real;
		i,j: entero;
		horario:AN(6);

		PROCEDIMIENTO def_colum() ES:  //asigno a cada columna del arreglo una franja horaria
			SEGUN reg.hora HACER:
				"8-10": j:=1;
				"10-12": j:=2;
				"12-16": j:=3;
				"16-18": j:=4;
				"18-20": j:=5;
			Fin_segun;
		Fin_proced;

		PROCEDIMIENTO inicializar() ES:
			PARA i:=1 a 11 HACER:
				PARA j:=1 a 6 HACER:
					arr[i,j]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		inicializar();

		MIENTRAS NFDA(Arch) HACER:
			def_colum();
			i:= reg.nro_caja;
			arr[i,j]:= arr[i,j] + reg.importe;  //acumulo en la fila y columna correspondiente al horario
			arr[11,j]:= arr[11,j] + reg.importe;  //acumulo en la ultima celda de la columna
			arr[i,6]:= arr[i,6] + reg.importe;  //acumulo en la ultima celda de la fila
			arr[11,6]:= arr[11,6] + reg.importe;  //aculumo en la ultima casilla, total gral
			LEER(Arch,reg);
		Fin_mientras;
		CERRAR(Arch);

		PARA i:=1 a 10 HACER:
			Escribir("Caja:",i,", importe total:"arr[i,6]); //muestro el total por fila (por caja)
		Fin_para;

		PARA j:=1 a 5 HACER:
			SEGUN j HACER:
				1: horario:= "8-10";
				2: horario:= "10-12";
				3: horario:= "12-16";
				4: horario:= "16-18";
				5: horario:= "18-20";
			Fin_segun; 
			Escribir("Para el horario:",horario,"el importe total fue:",arr[11,j]);  //muestro el total por columna (franja horaria)
		Fin_para;

		Escribir("Total general:",arr[11,6]);
Fin_accion.



//Ejercicio3.25
Una Fábrica que posee 4 plantas de producción en nuestro país y compra materia prima a 3 proveedores distintos, desea realizar un control de los montos totales correspondientes a compras realizadas en cada planta durante el último año, discriminados por proveedor y por mes. Para ello dispone de un archivo con los datos de las facturas correspondientes. (Aclaración: el archivo no está ordenado por ningún criterio)

COMPRAS
|Nro_Factura| |Proveedor (A,B,C)| |Fecha dd/mm/aaaa| |Nro_Planta 1..4| |Importe|

ACCION ejercicio3.25 ES	
	AMBIENTE
		compras = Registro
			nro_fac: N(8);
			prov: {"A","B","C"};
			fecha = Registro
				año: N(4);
				mes: 1..12;
				dia: 1..31;
			Fin_registro;
			nro_planta: 1..4;
			importe: N(8,2);
		Fin_registro;

		Arch: archivo de compras;
		reg: compras;

		arr: arreglo [1..4,1..13,1..4] de real;  
		i,j,k: entero;

		PROCEDIMIENTO inicializar() ES:
			PARA k:=1 a 4 HACER:  //plantas
				PARA j:=1 a 13 HACER:  //ultima columna por total de mes
					PARA i:=1 a 4 HACER:  //ultima fila por total de proveedores
						arr[i,j,k]:= 0;
					Fin_para;
				Fin_para;
			Fin_para;
		Fin_proced;
		
	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		inicializar();

		MIENTRAS NFDA(Arch) HACER:
			j:= reg.fecha.mes;
			k:= nro_planta;

			SEGUN reg.prov HACER:  //asigno una fila a cada proveedor 
				"A": i:=1; 
				"B": i:=2;
				"C": i:=3;
			Fin_segun;

			arr[i,j,k]:= arr[i,j,k] + reg.importe;  //importe del porveedor por un mes
			arr[i,13,k]:= arr[i,13,k] + reg.importe;  //total por proveedor
			arr[4,j,k]:= arr[4,j,k] + reg.importe;  //total por mes
			arr[4,13,k]:= arr[4,13,k] + reg.importe;  //total de planta por año
			
			LEER(Arch,reg);
		Fin_mientras;
		CERRAR(Arch);

		PARA k:=1 a 4 HACER:
			PARA i:=1 a 3 HACER:
				Escribir("Para la planta",k,"el total del proveedor",i,"es:",arr[i,13,k]);  //totales por proveedor
			Fin_para;

			PARA j:=1 a 12 HACER:
				Escribir("Para la planta",k,"El total por el mes",j,"es:",arr[4,j,k]);  //totales por mes
			Fin_para;

			Escribir("El total por año de la planta",k,"es:",arr[4,13,k]);  //total por año
		Fin_para;
Fin_accion.



//Ejercicio3.26
Se dispone de un archivo secuencial de facturas de una empresa de energía eléctrica, correspondientes a un año, con el siguiente formato:
|Nro_Factura| |Nro_Usuario| |Zona| |Fecha| |Consumo|

Se desea obtener un cuadro estadístico que informe los consumos y los importes totales facturados por mes discriminados por zona, con el siguiente formato:

        Zona A	Zona B	Zona C	Zona D	Total por mes
Enero	  ...	 ...	 ...	 ...	   ...
Febrero	  ...	 ...	 ...	 ...	   ...
...	      ...	 ...	 ...	 ...	   ...
Diciem	  ...	 ...	 ...	 ...	   ...
Totalz	  ...	 ...	 ...	 ...    Total General

La ciudad está dividida en 4 Zonas (A,B,C,D), la tarifa por Kw. que cobra la empresa en cada una de ellas es: Zona A:0,05 , ZonaB:0,07 , Zona C:0,09 y Zona D: 0,13.

ACCION ejercicio3.26 ES
	AMBIENTE
		facturas = Registro
			nro_fac
			nro_us
			zona: {"A","B","C","D"};
			fecha = Registro
				año
				mes
				dia
			Fin_registro;
			consumo
		Fin_registro;

		Arch: archivo de facturas;
		reg: facturas;

		regis= Registro  
			consumo
			importe
		Fin_registro;

		arr: arreglo [1..13,1..5] de regis;  //en cada celda del arreglo tengo el registro (consumo e importe) de cada zona 
		i,j: entero;
		tarifa: real;

		PROCEDIMIENTO iniciar() ES:
			PARA i:=1 a 13 HACER:
				PARA j:=1 a 5 HACER:
					arr[i,j]:=0;
				Fin_para;
			Fin_para;
		Fin_proced;

		PROCEDIMIENTO convertir() ES:
			SEGUN reg.zona HACER:
				"A": tarifa:=0,05;
				"B": tarifa:=0,07;
				"C": tarifa:=0,09;
				"D": tarifa:=0,13;
			Fin_segun;
		Fin_proced;

	PROCESO
		ABRIR(Arch);
		LEER(Arch,reg);

		iniciar();

		MIENTRAS NFDA(Arch) HACER:
			i:= reg.fecha.mes;
			j:= reg.zona;

			convertir();

			arr[i,j].consumo:= arr[i,j].consumo + reg.consumo;  
			arr[13,j].consumo:= arr[13,j].consumo + reg.consumo; //acumulo en la ultima celda de la columna el consumo por zona
			arr[i,5].consumo:= arr[i,5].consumo + reg.consumo;  //acumulo en la ultima celda de la fila el consumo por mes
			arr[i,j].importe:= arr[i,j].importe + tarifa;
			arr[13,j].importe:= arr[13,j].importe + tarifa  //acumulo en la ultima celda de la columna el importe por zona
			arr[i,5].importe:= arr[i,5].importe + tarifa  //acumula en la ultima celda de la fila el importe por mes
		
			LEER(Arch,reg);
		Fin_mientras;
		CERRAR(Arch);

		PARA i:=1 a 12 HACER:
			PARA j:=1 a 4 HACER:
				Escribir("En el mes",i,"y la zona",j):
				Escribir("El consumo es:",arr[i,j].consumo);
				Escribir("El importe es:",arr[i,j].importe);
			Fin_para;
		Fin_para;

		PARA i:=1 a 12 HACER:
			Escribir("El consumo general del mes",i,"es:"arr[i,5].consumo);
			Escribir("El importe general del mes",i,"es:"arr[i,5].importe);
		Fin_para;

		PARA j:=1 a 4 HACER:
			Escribir("El consumo genral de la zona",j,"es:"arr[13,j].consumo);
			Escribir("El importe general de la zona",j,"es:",arr[13,j].importe);
		Fin_para;
Fin_accion.



//Ejercicio3.27
Se desea efectuar una estadística de ventas. Se cuenta para ello con una secuencia de las facturas emitidas, las cuales son identificadas por un Número; dicha secuencia contiene información relativa al cliente: su Número y Zona a la cual pertenece, como así también el Tipo de mercadería entregada, la Cantidad de unidades, el total gravado, el total exento de IVA y el valor del IVA. Construya un algoritmo que emita por zona, y dentro de la zona por tipo de mercadería el total de unidades vendidas, el total gravado, el total exento y el total de IVA, y además un total general con la misma información. Hay 9 zonas y 4 tipos de mercadería.

ACCION ejercicio3.27 ES
	AMBIENTE
		clientes = Registro
			nro_cli: N(8);
			zona: 1..9;
			tipo_merc: 1..4;
			cant_unid: N(8);
			tot_grav: N(8,2);
			tot_siniva: N(8,2);
			tot_coniva: N(8,2);
		Fin_registro;

		Arch: archivo de clientes ordenado por nro_cli;
		reg: clientes;

		totales = Registro
			cant_unid: N(8);
			tot_grav: N(8,2);
			tot_siniva: N(8,2);
			tot_coniva: N(8,2);
		Fin_registro;

		arr: arreglo [1..10,1..5] de totales;
		i,j: entero;

		PROCEDIMIENTO iniciar() ES:
			PARA i:=1 a 10 HACER:
				PARA j:=1 a 5 HACER:
					arr[i,j]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		iniciar();

		MIENTRAS NDFA(Arch) HACER:
			i:= reg.tipo_merc;
			j:= reg.zona;

			arr[i,j].cant_unid:= arr[i,j] + reg.cant_unid;
			arr[i,j].tot_grav:= arr[i,j].tot_grav + reg.tot_grav; 
			arr[i,j].tot_siniva:= arr[i,j].tot_siniva + reg.tot_siniva;
			arr[i,j].tot_coniva:= arr[i,j].tot_coniva + reg.tot_coniva;

			//acumulo los totales por tipo de zona (en la ultima celda de cada fila, columna->5)
			
			arr[i,5].cant_unid:= arr[i,5].cant_unid + reg.cant_unid;
			arr[i,5].tot_grav:= arr[i,5].tot_grav + reg.tot_grav;
			arr[i,5].tot_siniva:= arr[i,5].tot_siniva + reg.tot_siniva;
			arr[i,5].tot_coniva:= arr[i,5].tot_coniva + reg.tot_coniva;

			//acumulo los totales por mercadería (en la ultima celda de cada columna, fila->10)
			
			arr[10,j].cant_unid:= arr[10,j].cant_unid + reg.cant_unid;
			arr[10,j].tot_grav:= arr[10,j].tot_grav + reg.tot_grav;
			arr[10,j].tot_siniva:= arr[10,j].tot_siniva + reg.tot_siniva;
			arr[10,j].tot_coniva:= arr[10,j].tot_coniva + reg.tot_coniva;
			
			//acumulo los totales por zona y mercaderia (ultima celda de fila y columna)

			arr[10,5].cant_unid:= arr[10,5].cant_unid + reg.cant_unid;
			arr[10,5].tot_grav:= arr[10,5].tot_grav + reg.tot_grav;
			arr[10,5].tot_siniva:= arr[10,5].tot_siniva + reg.tot_siniva;
			arr[10,5].tot_coniva:= arr[10,5].tot_coniva + reg.tot_coniva;

			LEER(Arch,reg);
		Fin_mientras;
		CERRAR(Arch);

		PARA i:=1 a 9 HACER:
			PARA j:=1 a 4 HACER:
				Escribir("Para la zona:",i,"y el tipo de mercadería:",j);
				Escribir("Cantidad de unidades vendidas:",arr[i,j].cant_unid);
				Escribir("Total gravado:",arr[i,j].tot_grav);
				Escribir("Total excento del IVA:",arr[i,j].tot_siniva);
				Escribir("Total con IVA:",arr[i,j].tot_coniva);
			Fin_para;
		Fin_para;

		PARA i:=1 a 9 HACER:
			Escribir("Totales generales por la zona:",i);
			Escribir("Cantidad de unidades vendidas:",arr[i,5].cant_unid);
			Escribir("Total gravado:",arr[i,5].tot_grav);
			Escribir("Total excento del IVA:",arr[i,5].tot_siniva);
			Escribir("Total con IVA:",arr[i,5].tot_coniva);
		Fin_para;

		PARA j:=1 a 4 HACER:
			Escribir("Totales generales por la mercadería de tipo:",j);
			Escribir("Cantidad de unidades vendidas:",arr[10,j].cant_unid);
			Escribir("Total gravado:",arr[10,j].tot_grav);
			Escribir("Total excento del IVA:",arr[10,j].tot_siniva);
			Escribir("Total con IVA:",arr[10,j].tot_coniva);
		Fin_para;

		Escribir("Total general por zona y mercadería:");
		Escribir("Cantidad de unidades vendidas:",arr[10,5].cant_unid);
		Escribir("Total gravado:",arr[10,5].tot_grav);
		Escribir("Total excento del IVA:",arr[10,5].tot_siniva);
		Escribir("Total con IVA:",arr[10,5].tot_coniva);
Fin_accion. 



//Ejercicio3.28
La Municipalidad de Resistencia desea obtener una estadística de los valores (en $) de los terrenos de la ciudad, discriminados por zona y ubicación dentro de la manzana (en esquina, interna, etc.), y los totales por zona y ubicación. El valor de cada final de cada terreno es igual a:

Valor del Terreno = Superficie terreno (en M2) * valor del M2 * coeficiente de incremento

Para ello cuenta con la siguiente información:

1-Un archivo de los terrenos con los siguientes datos:
|Nro_Terreno| |Zona| |Ubicacion| |Superficie|
Zona: codificadas de A a F
Ubicación: codificada de 1 a 10

2-Un arreglo V que contiene los valores del M2 por zona.

ACCION ejercicio3.28 (v: arreglo [A..F] de real) ES  //este arreglo tiene los valores por metro cuadrado por zona
	AMBIENTE
		terrenos = Registro
			nro_terreno: N(8);
			zona: (A..F);
			ubic: (1..10);
			superf: N(8,2);
		Fin_registro;

		Arch: archivo de terrenos;
		reg: terrenos;

		arr: arreglo [1..11,A..G] de real;
		i,j: entero;

		coef: entero;

		PROCEDIMIENTO inicializar() ES:
			PARA i:=1 a 11 HACER:
				PARA j:=A a G HACER:
					arr[i,j]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		inicializar();

		MIENTRAS NFDA(Arch) HACER:
			i:= reg.ubicación;
			j:= reg.zona;

			valor_terreno:= reg.superf * v[j] * coef;  //v[j] es el valor que tira el arreglo de zona 
			arr[i,j]:= arr[i,j] + valor_terreno;
			arr[i,G]:= arr[i,G] + valor_terreno; //acumulo en la ultima casilla de la ultima columna (G), ubicacion
			arr[11,j]:= arr[11,j] + valor_terreno; //acumulo en la ultima casilla de la ultima fila (11), zona 
			arr[11,G]:= arr[11,G] + valor_terreno; 

			LEER(Arch,reg);
		Fin_mientras;
		CERRAR(Arch);

		PARA i:=1 a 10 HACER:
			PARA j:=A a F HACER:
				Escriba("Para la ubicación:",i,"y la zona:",j,"el valor del terreno es:",arr[i,j]);
			Fin_para;
		Fin_para;

		PARA i:=1 a 10 HACER:
			Escribir("Para la ubicación:",i,"el valor del terreno es:",arr[i,G]);
		Fin_para;

		PARA j:=A a F HACER:
			Escribir("Para la zona:",j,"el valor del terreno es:",arr[11,j]);
		Fin_para;

		Escribir("El total general de ubicación y zona es:",arr[11,G]);
Fin_accion.

			
		 
//Ejercicio3.31 ¿¿¿??¡?¡?¡?
Se lleva a cabo una encuesta a fin de anticipar los posibles porcentajes de votos que obtendrán cada uno de los 4 partidos políticos de mayor peso en las elecciones. Los datos solicitados a los encuestados son: partido al que Algoritmos y Estructuras de Datos 2019 votara, edad, sexo y partido al que votó en las elecciones anteriores. Las respuestas se tabulan de acuerdo a los siguientes criterios:

Partidos: P1, P2, P3, P4, Otro, En Blanco, Indeciso
Edad: 18-25, 26-35, 36-45, 46-55, 56-65, +65
Sexo: Femenino, Masculino

Diseñe un algoritmo que, utilizando un arreglo de 4 dimensiones (partido al que votara, edad, sexo , partido al que votó), permita responder a las siguientes consultas:
Cantidad de personas de cierta edad que votarán a un Partido dado.
Cantidad de personas de un sexo dado que votarán a un determinado Partido.
Cantidad de personas de cierta edad que votaron a un determinado Partido y actualmente votarán a otro Partido dado.
Cantidad de personas de un sexo dado que votaron a un determinado Partido y actualmente votarán a otro Partido dado.

ACCION ejerecicio3.31 ES
	AMBIENTE
		padron = Registro
			edad: N(2);
			sexo: {"M","F"};
			p_ant: {"P1","P2","P3","P4","O","B","I"};  //O=otro, B=blanco, I=Indeciso
			p_act: {"P1","P2","P3","P4","O","B","I"};
		Fin_registro;

		Arch: archivo de padron;
		reg: padron;

		arr: arreglo [1..7,1..6,1..2,1..7] de entero;
		i,j,k,l: entero;

		rta: {"S","N"};

		PROCEDIMIENTO inicializar() ES:
			PARA i:=1 a 7 HACER:
				PARA j:=1 a 6 HACER:
					PARA k:=1 a 2 HACER:
						PARA l:=1 a 7 HACER:
							arr[i,j,k,l]:= 0;
						Fin_para;
					Fin_para;
				Fin_para;
			Fin_para;
		Fin_proced;
		
		PROCEDIMIENTO ingreso_datos() ES:
			Escribir("Ingrese su edad:");
			Leer(reg.edad);

				SEGUN reg.edad HACER:  //en las columnas tengo las edades
					>=18 Y <=25: j:=1;
					>=26 Y <=35: j:=2;
					>=36 Y <=45: j:=3;
					>=46 Y <=55: j:=4;
					>=56 Y <=65: j:=5;
					>65: i:=6;
				Fin_segun;

			Escribir("Ingrese su sexo:");
			Leer(reg.sexo);

				SI reg.sexo="F" ENTONCES:
					k:= 1;
				SINO	
					k:= 2;
				Fin_si;

			Escribir("Ingrese el partido que votará: 'P1, P2, P3, P4, O(Otro), B(en blanco), I(indeciso)':");
			Leer(reg.p_act);

				SEGUN reg.p_act HACER:  //en las filas tengo los partidos que votó anteriormente
					"P1": i:= 1;
					"P2": i:= 2;
					"P3": i:= 3;
					"P4": i:= 4;
					"O": i:= 5;
					"B": i:= 6;
					"I": i:= 7;
				Fin_segun;

			Escribir("Ingrese el partido que votó anteriormente: 'P1, P2, P3, P4, O(Otro), B(en blanco), I(indeciso)':");
			Leer(reg.p_ant);

				SEGUN reg.p_ant HACER:  
					"P1": l:= 1;
					"P2": l:= 2;
					"P3": l:= 3;
					"P4": l:= 4;
					"O": l:= 5;
					"B": l:= 6;
					"I": l:= 7;
				Fin_segun;
		Fin_proced;

	PROCESO
		ABRIR S/(Arch);
		inicializar();
		rta:= V;

		Escribir("¿Desea ingresar datos? V/F");
		Leer(rta);

		MIENTRAS rta = "V" HACER:	
			ingresar_datos();
Fin_accion.



ACCION ejercicio_parcial(arr_libros: arreglo [1..100,1..5] de Libros) ES	
	AMBIENTE
		Libros = Registro
			tít: AN(50);
		Fin_registro;
		i,j:entero;  //el id del libro está dado por los índices 

		totales = Registro
			tot_disp_dig:
			tot_disp_nodig:
			tot_nodisp_dig:
			tot_nodisp_nodig:
		Fin_registro;

		arr: arreglo [1..6,1..3] de totales;
		k,l: entero

		Ejemplares = Registro
			id_ejem: N(8);
			id_libro: N(3);
			suc: 1..5;
			digital: {"si","no"};
			disponible: booleano;
		Fin_registro;

		Arch: archivo de Ejemplares ordenado por id_ejem e id_libro;
		reg: Ejemplares;

		PROCEDIMIENTO iniciar() ES
			PARA k:=1 a 6 HACER:
				PARA l:= 1 a 3 HACER:
					arr[k,l]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;

		mayor_suc, 

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		iniciar();

		MIENTRAS NFDA(Arch) HACER:
			k:= reg.suc;
			l:= reg.digital;

			SEGUN reg.suc HACER:
				"1": k:= 1;
				"2": k:= 2;
				"3": k:= 3;
				"4": k:= 4;
				"5": k:= 5;
			Fin_segun;

			SI (reg.digital = "si") ENTONCES:
				l:= 1;  //"si"
			SINO
				l:= 2;  //"no"
			Fin_si;

			SI (reg.disponible) ENTONCES:  //si disponible = verdadero
				SI (reg.digital = "si") ENTONCES:
					arr[k,l].tot_disp_dig:= arr[k,l].tot_disp_dig + 1;
					arr[6,l].tot_disp_dig:= arr[6,l].tot_disp_dig + 1;  //aculumo en la ultima fila (columna de disponibles digitales)
					arr[i,3].tot_disp_dig:= arr[i,3].tot_disp_dig + 1;  //acumulo en la ultima columna (fila de sucursal)
				SINO
					arr[k,l].tot_disp_nodig:= arr[k,l].tot_disp_nodig + 1;
					arr[6,l].tot_disp_nodig:= arr[6,l].tot_disp_nodig + 1;  //aculumo en la ultima fila (columna de disponibles no digitales)
					arr[i,3].tot_disp_nodig:= arr[i,3].tot_disp_nodig + 1;  //acumulo en la ultima columna (fila de sucursal)	
				Fin_si;
			SINO	
				SI (reg.digital = "si") ENTONCES:
					arr[k,l].tot_nodisp_dig:= arr[k,l].tot_nodisp_dig + 1;
					arr[6,l].tot_nodisp_dig:= arr[6,l].tot_nodisp_dig + 1;  //aculumo en la ultima fila (columna de no disponibles digitales)
					arr[i,3].tot_nodisp_dig:= arr[i,3].tot_nodisp_dig + 1;  //acumulo en la ultima columna (fila de sucursal)
				SINO
					arr[k,l].tot_nodisp_nodig:= arr[k,l].tot_nodisp_nodig + 1;
					arr[6,l].tot_nodisp_nodig:= arr[6,l].tot_nodisp_nodig + 1;  //aculumo en la ultima fila (columna de no disponibles no digitales)
					arr[i,3].tot_nodisp_nodig:= arr[i,3].tot_nodisp_nodig + 1;  //acumulo en la ultima columna (fila de sucursal)	
				Fin_si;
			Fin_si;
			LEER(Arch,reg);
		Fin_mientras;
		CERRAR(Arch);


			
//Ejercicio de parcial (recu diciembre 2023)
ACCION ejerc_parcial ES
	AMBIENTE 
		arr: arreglo [1..5,1..13] de paquetes;
		i,j: entero;

		paquetes = Registro
			cant_paq: N(8);
			nombre: AN(30);
		Fin_registro;

		clientes = Registro
			nro_cliente: N(8);
			apelynom: AN(30);
			dni: N(8);
			id_paq: 1..12;
			saldo: N(8,2);
			estado: {"saldado", "saldo favor", "deudor"};
			categ: {"simple", "plata", "oro", "diamante"};
			puntos: N(8);
			fecha_baja = Registro
				año:
				mes:
				dia:
			Fin_registro;
		Fin_registro;

		Arch: archivo de clientes ordenado por nro_cliente;
		reg: clientes;

		paquetes_turisticos = Registro
			id_paq: 1..12;
			nombre: AN(30);
			costo: N(8,2);
			destino: AN(30);
		Fin_registro;

		Arch2: archivo de paquetes_turisticos indexado por id_paq;
		reg2: paquetes_turisticos;

		tot_paqsald: entero;
		mayor, paq_may: entero;
		nom_may: AN(30);

		PROCEDIMIENTO inicializar() ES:
			PARA i:=1 a 5 HACER:
				PARA j:=1 a 13 HACER:
					arr[i,j]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);
		ABRIR E/(Arch2);
		LEER(Arch2,reg2);

		reg.id_paq:= reg2.id_paq;
		
		inicializar();

		mayor: LV;
	
		MIENTRAS NFDA(Arch) HACER:
			i:= reg.categ;
			j:= reg.id_paq;

			SEGUN reg.categ HACER:
				"simple": i:=1;
				"plata": i:=2;
				"oro": i:=3;
				"diamante": i:=4;
			Fin_segun;

			SEGUN reg.id_paq HACER:
				"1": j:=1;
				"2": j:=2;
				"3": j:=3;
				"4": j:=4;
				"5": j:=5;
				"6": j:=6;
				"7": j:=7;
				"8": j:=8;
				"9": j:=9;
				"10": j:=10;
				"11": j:=11;
				"12": j:=12;
			Fin_segun;

			SI (reg.estado = "saldado") O (reg.estado = "saldo favor") ENTONCES:
				arr[i,j].cant_paq:= arr[i,j].cant_paq + 1;
				arr[i,j].nombre:= reg2.nombre; 

