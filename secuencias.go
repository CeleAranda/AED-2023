//secuencia de datos
//almacenamiento en memoria externa (disco) -> garantiza persistencia de los datos 

//Ejercicio 1
Dada una secuencia de letras del alfabeto que finaliza con una marca '*', contar cuantas letras "A" hay en la secuencia.

ACCION ejercicio1 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v: caracter;
		conta: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);
		conta:=0;

		MIENTRAS (v <> "*") HACER:
			SI (v = "a") ENTONCES:
				conta:= conta + 1;
			Fin_si;
			AVZ(sec,v);
		Fin_mientras;

		Escribir("Cantidad de letras 'a':",conta);
	Fin_accion.

//Ejercicio 2
Dada una secuencia de letras del alfabeto que finaliza con la letra "Z", contar cuantas consonantes hay en la secuencia.

ACCION ejercicio2 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v: caracter;
		cont: entero;
		vocales= {"a","e","i","o","u"}; //los conjuntos se definen como constantes, porque es todo el conjunto

	PROCESO
		ARR(sec);
		AVZ(sec,v);
		cont:=0;

		REPETIR 
			SI (v NO EN vocales) ENTONCES:
				cont:= cont + 1;
			Fin_si
			AVZ(sec,v);
		HASTA QUE (v = "z")

		CERRAR(sec);

		Escribir("Cantidad de consonantes:",cont);
	Fin_accion.


		// MIENTRAS (v <> "z") HACER:
		// 	SI (v NO en vocales) ENTONCES:
		// 		cont:= cont + 1;
		// 	Fin_si;
		// 	AVZ(sec,v);
		// Fin_mientras;

//Ejercicio 3
Se dispone de una secuencia de caracteres y se desea obtener una secuencia de salida que resulte de copiar la secuencia de entrada, descartando el caracter "$".

ACCION ejercicio3 ES
	AMBIENTE
		sec, sec_sal: secuencia de caracteres;
		v: caracter;

	PROCESO
		ARR(sec);
		AVZ(sec,v);
		CREAR(sec);

		MIENTRAS NFDS(sec) HACER:
			SI (v <> "$") ENTONCES:
				GRABAR(sec_sal,v);
				AVZ(sec,v);
			SINO 
				AVZ(sec,v);
			Fin_si;
		Fin_mientras;

		CERRAR(sec);
		CERRAR(sec_sal);
Fin_accion.

//Ejercicio4
Se dispone de una secuencia de números de socios, y se desea saber la cantidad de socios que están registrados.

ACCION ejercicio4 ES
	AMBIENTE
		sec: secuencia de enteros;
		v: entero;
		conts: entero; 

	PROCESO
		ARR(sec);
		AVZ(sec,v);
		conts:=0;

		MIENTRAS NFDS(sec) HACER:
			conts:= conts + 1;
			AVZ(sec,v);
		Fin_mientras;

		Escribir("Cantidad de socios registrados:",conts);
Fin_accion.

//Ejercicio5 
La secuencia de socios del problema anterior tiene el inconveniente de que los números están ordenados pero no son correlativos. Se desea generar una secuencia que contenga los números de socios que no figuran en la secuencia de socios.

ACCION ejercicio5 ES
	AMBIENTE	
		sec, sec_sal: secuencia de enteros;
		v,res1, res2,cont,i: entero;
		
	
	PROCESO
		ARR(sec);
		AVZ(sec,v);
		CREAR(sec_sal);

		MIENTRAS NFDS(sec) HACER:
			res1:= v;
			AVZ(sec,v);
			res2:= v;

			PARA i:= res1 a res2 HACER:
				 cont:= i + 1;

				 SI (cont < res2) ENTONCES:
				 	GRABAR(sec_sal,cont);  //yo grabo en salida los numeros intermedios entre las dos ventanas, sé cuáles son esos números gracias al contador
				Fin_si;
			Fin_para;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);
		CERRAR(sec_sal);
Fin_accion.

//Ejercicio extra
Se dispone de una secuencia de números de socios, y se desea grabar en una secuencia de salida los socios con números par que esán registrados.

ACCION ejercicioextra ES
	AMBIENTE
			sec, sec_sal: secuencia de caracteres;
			v: caracter;

		PROCESO
			ARR(sec);
			AVZ(sec,v);
			CREAR(sec);

			MIENTRAS NFDS(sec) HACER:
				SI (v MOD 2 = 0) ENTONCES:
					GRABAR(sec_sal,v);
				Fin_si;	
				AVZ(sec,v);
			Fin_mientras;

			CERRAR(sec);
			CERRAR(sec_sal);
Fin_accion.

// |h|o|l|a|1|2|4| sec de caracter  
// |1902|456|372|32892| sec de entero

//Ejercicio6
Dada una secuencia de enteros que almacena la cantidad de habitantes de las ciudades capitales de las 23 provincias de la República Argentina, discriminados 4 categorías: menores de 18 años (varones y mujeres) y mayores de 18 años (varones y mujeres). Se pide determinar la población total y los siguientes porcentajes: masculinos, femeninos, mayores de 18 y menores de 18.

ACCION ejercicio6 ES
	AMBIENTE
		sec: secuencia de enteros;
		v: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);


		PARA i:=1 a 23 HACER:
			v_menor:= v_menor + v;
			AVZ(sec,v);
			m_menor:= m_menor + v;
			AVZ(sec,v);
			v_mayor:= v_mayor + v;
			AVZ(sec,v);
			m_mayor:= m_mayor + v;
			AVZ(sec,v);	
		Fin_para;

		tot:= v_menor + m_menor + v_mayor + m_mayor;
		
		CERRAR(sec);

		Escribir("El total de la población es:", tot);
		Escribir("El porcentaje total de masculinos menores de 18 años es:", (v_menor/tot));
		Escribir("El porcentaje total de femeninas menores de 18 años es:", (m_menor/tot));
		Escribir("El porcentaje total de masculinos mayores de 18 años es:", (v_mayor/tot));
		Escribir("El porcentaje total de femeninas menores de 18 años es:", (m_mayor/tot));

Fin_accion.

//Ejercicio7
Se tiene una secuencia de enteros que contiene todos los CUIT de los empleados de una empresa, la misma termina con el CUIT 0. Para evitar largas esperas en los días de pago, la empresa necesita organizarlos de acuerdo al último dígito de su documento, generar una secuencia con los CUIT de las personas que su número de documento termine con 0, 1, 2 o 3.

ACCION ejercicio7 ES	
	AMBIENTE
		sec, sec_sal: secuencia de caracteres;
		cuit: caracter;
		valid= {1,2,3,0};

	PROCESO
		ARR(sec);
		AVZ(sec,cuit);
		CREAR(sec_sal);

		MIENTRAS (cuit <> 0) HACER:
			SI ((cuit MOD 100) DIV 10) EN valid ENTONCES:  //cuit MOD 10 da los ultimos dos díguitos, div 10 da la decena
				GRABAR(sec_sal,cuit);
			Fin_si;
			AVZ(sec,cuit);
		Fin_mientras;

		CERRAR(sec);
		CERRAR(sec_sal);
Fin_accion.

//Ejercicio8
Teniendo en cuenta el ejercicio anterior, se solicita que la secuencia de salida sea una secuencia de caracteres y los CUIT se separen unos de otros con el caracter "-".

ACCION ejercicio8 ES
	AMBIENTE
		sec: secuencia de enteros;
		cuit, i: entero;
		sec_sal: secuencia de caracteres;

	PROCESO
		ARR(sec);
		AVZ(sec,cuit);
		CREAR(sec_sal);		

		MIENTRAS (cuit <> 0) HACER:
			PARA i:=1 a 11 HACER:
				GRABAR(sec_sal,cuit);
				AVZ(sec,cuit);
			Fin_para;
			GRABAR(sec_sal,"-");


//Ejercicio9
Se dispone de una secuencia de caracteres. Se desea determinar la cantidad de palabras que comienzan con la letra 'I'.

ACCION ejercicio9 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v: caracter;
		cont: entero;
	
	PROCESO
		ARR(sec);
		AVZ(sec,v);
		cont:=0;

		MIENTRAS NFDS(sec) HACER:
			SI (v = "i") ENTONCES:
				cont:= cont + 1;
			Fin_si;
			
			MIENTRAS (v <> ".") HACER:
				MIENTRAS (v <> " ") HACER:
					AVZ(sec,v);	
				Fin_mientras;
				AVZ(sec,v);
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);

		Escribir("Cantidad de palabras que incian con la letra i:", cont);
Fin_accion.

//Ejercicio10
Se dispone de una secuencia de caracteres. Se desea permita contar la cantidad de palabras que comienzan con una letra dada.

ACCION ejercicio10 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v,car: caraacter;
		cont: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);

		Escribir("Ingrese una letra:");
		Leer(car);

		cont:= 0;

		MIENTRAS NFDS(sec) HACER:
			SI (v = car) ENTONCES:
				cont:= cont + 1;
			Fin_si;

			MIENTRAS (v <> ".") HACER: //trato oracion 
				MIENTRAS (v <> " ") HACER: //trato palabra
					AVZ(sec,v);	
				Fin_mientras;
				AVZ(sec,v);
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);

		Escribir("La cantidad de palabras con la letra",car,"es:",cont);
Fin_accion.

//Ejercicio11
Dada una secuencia de caracteres, determinar la cantidad de palabras de 4 caracteres (letras)

ACCION ejercicio11 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v: caraacter;
		contp, contcar: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);
		
		contp:=0;
		contcar:=0;

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:
				MIENTRAS (v <> " ") HACER:
					contcar:= contcar + 1;
					AVZ(sec,v);
				Fin_mientras;
					
				SI (contcar = 4) ENTONCES:
					contp:= contp + 1;
				Fin_si;

				AVZ(sec,v);
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);

		Escribir("La cantidad de palabras con 4 caracteres es de:",contp);
Fin_accion.

//Ejercicio12
Se dispone de una secuencia de caracteres. Se desea listar las palabras que comiencen con "ALG".

ACCION ejercicio12 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v,car: caraacter;
		
	PROCESO
		ARR(sec);
		AVZ(sec,v);
		car:= v;
		
		Escribir("Lista de palabras que comienzan con 'alg':");

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:
				MIENTRAS (v <> " ") HACER:
					SI (v = "a") ENTONCES:
						AVZ(sec,v);
					SINO
						SI (v = "l") ENTONCES:
							AVZ(sec,v);
						SINO	
							SI (v = "g") ENTONCES:
								Escribir("Alg");
								MIENTRAS (v <> " ") HACER:
									Escribir(v);
									AVZ(sec,v);
							Fin_si;
						Fin_si;
					Fin_si;
					AVZ(sec,v);
				Fin_mientras;
				AVZ(sec,v);
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);
Fin_accion. 

//Ejercicio13
A partir del ejercicio anterior, determinar el porcentaje que representan las palabras que comienzan con "ALG" sobre todas las palabras de la secuencia.

ACCION ejercicio13 ES
	AMBIENTE
		sec: secuencia de caracteres;
		c: caracter;
		tot,tot_a: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);	 
		
		tot:= 0;
		tot_a:=0;

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:
				MIENTRAS (v <> " ") HACER:
					SI (v = "a") ENTONCES:
						AVZ(sec,v);
					SINO
						SI (v = "l") ENTONCES:
							AVZ(sec,v);
						SINO	
							SI (v = "g") ENTONCES:
								tot_a:= tot_a + 1;
							Fin_si;
						Fin_si;
					Fin_si;
					AVZ(sec,v);
				Fin_mientras;
				tot:= tot + 1;
				AVZ(sec,v);
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);

		Escribir("El porcentaje de palabras que comienzan con 'ALG' es:", tot_a/tot);
Fin_accion. 
							
//Ejercicio14
Se dispone de una secuencia de caracteres y se desea saber la cantidad de caracteres (incluidos los espacios) que existen entre una coma y la siguiente. Se debe considerar que puede haber más de un par de comas, y que las subsecuencias inicial y final deben descartarse por no cumplir la condición enunciada. Supóngase que las comas son siempre contiguas al último caracter de la palabra.

ACCION ejercicio14 ES
	AMBIENTE
		sec: secuencia de caracteres;
		v: caraacter;
		contcar: entero;

	PROCESO	
		ARR(sec);
		AVZ(sec,v);

		contcar:=0;

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:
				MIENTRAS (v <> ",") HACER:
					AVZ(sec,v);  //avanzo todo hasta la primer coma
				Fin_mientras;
				AVZ(sec,v); //avanzo la coma

				// MIENTRAS (v <> ",") HACER:
				// 	contcar:= contcar + 1;
				// 	MIENTRAS (v <> " ") HACER: //trato palabra
				// 		AVZ(sec,v);
				// 	Fin_mientras;
				// 	AVZ(sec,v);
				// Fin_mientras;

				REPETIR
					contcar:= contcar + 1;
					MIENTRAS (v <> " ") HACER:
						AVZ(sec,v);
					Fin_mientras;
					AVZ(sec,v);
				HASTA QUE (v = ",");
				
				AVZ(sec,v); //avanzo la coma
			Fin_mientras;	
			AVZ(sec,v);
		Fin_mientras;	
		CERRAR(sec);
Fin_accion.

//Ejercicio15
Se desea saber la cantidad promedio de palabras que contienen las oraciones de una secuencia de caracteres. Supóngase que los puntos son siempre contiguos al último caracter de la palabra. La secuencia finaliza con una marca.

ACCION ejercicio15 ES	
	AMBIENTE
		sec: secuencia de caracteres;
		v: caraacter;
		contp, conto: entero;
	
	PROCESO
		ARR(sec);
		AVZ(sec,v);

		contp:= 0;
		conto:= 0;

		MIENTRAS (v <> marca) HACER:
			conto:= conto + 1; //cuento oracion
			MIENTRAS (v <> ".") HACER: //trato oracion
				contp:= contp + 1;	//cuento palabra
				MIENTRAS (v <> " ") HACER: //trato palabra
					AVZ(sec,v);
				Fin_mientras;
				AVZ(sec,v);
			Fin_mientras;
		Fin_mientras;

		CERRAR(sec);

		Escribir("El promedio de palabras por oración es:",contp/conto);
Fin_accion.

//Ejercicio16
Se dispone de una secuencia de números de DNI asignados a un circuito electoral, generar otra secuencia de números que contenga los DNI múltiplos de 3.

ACCION ejercicio16 ES	
	AMBIENTE
		sec, sec2: secuencia de enteros;
		v: entero;

	PROCESO	
		ARR(sec);
		AVZ(sec,v);
		CREAR(sec2);

		MIENTRAS NFDS(sec) HACER:
			SI (v MOD 3 = 0) ENTONCES:
				GRABAR(sec2,v) ENTONCES:
			Fin_si;
			AVZ(sec,v);
		Fin_mientras;
Fin_accion.

//Ejercicio17
Se desea calcular el costo de un telegrama, que se determina en función del número de palabras (que vale V1 cada una), salvo que el promedio de letras por palabra supere las cinco letras, caso en que cada palabra vale V2.

ACCION ejercicio17 ES
	AMBIENTE
		SEC: secuencia de caracteres;
		c: caraacter;
		contp,contcar: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:
				contp:= contp + 1; //cuento las palabras
				MIENTRAS (v <> " ") HACER:
					contcar:= contcar + 1; //cuento las letras de las palabras
					AVZ(sec,v);
				Fin_mientras;

				SI (contcar > 5) ENTONCES:
					valor:= contcar * V2;
				SINO
					valor:= contcar * V1;	
				Fin_si;
				
				tot:= valor * contp;
				AVZ(sec,v);
			Fin_mientras;
		Fin_mientras;

		CERRAR(sec);
		Escribir("El valor del telegrama es:", tot);
Fin_accion.

//Ejercicio18
Escribir un algoritmo que produzca una secuencia de salida que contenga una oración formada por las palabras en posición par de cada oración de una secuencia texto de entrada, que además comienzan con la letra 'M'.

ACCION ejercicio18 ES	
	AMBIENTE
		sec,sec2: secuencia de caracteres;
		v: caraacter;
		cont: entero;

	PROCESO
		ARR(sec);
		AVZ(sec,v);
		CREAR(sec2);

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:
				cont:= cont + 1;  //cuento la posicion
				SI (v = "m") ENTONCES:
					SI (cont MOD 2 = 0) ENTONCES:
						GRABAR(sec2, v); 
						MIENTRAS (v <> " ") HACER:
							GRABAR(sec2,v);
							AVZ(sec,v);
						Fin_mientras;	
					Fin_si;
				Fin_si;
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);
		CERRAR(sec2);

Fin_accion.

//Ejercicio19
Dada una secuencia de caracteres, generar otra secuencia con todas las palabras de tres caracteres. Informar el porcentaje de las mismas, sobre el total de palabras de la secuencia de entrada. Considerar que puede haber más de un blanco entre palabras.

ACCION ejercicio19 ES
	AMBIENTE
		sec, sec2: secuencia de caraacteres;
		v: caraacter;
		contcar, totp, contp3: entero
					
	PROCESO
		ARR(sec);
		AVZ(sec,v);
		CREAR(sec2);

		contcar:= 0; //cuenta caracteres
		totp:= 0; //cuenta todas las palabras de la secuencia
		contp3:= 0; //cuenta el total de palabras de 3 caracteres

		MIENTRAS NFDS(sec) HACER:
			MIENTRAS (v <> ".") HACER:  //trato oracion
				MIENTRAS (v = " ") HACER:
					AVZ(sec,v);
				Fin_mientras;

				totp:= totp + 1; //cuento la palabra de la sec de entrada
				
				MIENTRAS (v <> " ") //trato palabra
					car1:= car2; //u
					car2:= car3; //n
					car3:= v; //a
					contcar:= contcar + 1;
					AVZ(sec,v);
				Fin_mientras;
				
				SI (contcar = 3) ENTONCES:
					contp3:= contp3 + 1;
					GRABAR(sec2,car1);
					GRABAR(sec2,car2);
					GRABAR(sec2,car3);
				Fin_si;
			Fin_mientras;
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);
		CERRAR(sec2);
						
		Escribir("La secuencia contiene",totp,"palabras, de las cuales el porcentaje de palabras que presentan 3 letras es de:",contp3/totp);
Fin_accion.

//Ejercicio20
Se dispone de dos secuencias texto formadas por oraciones bimembres (sujeto y predicado, separados por comas). Se desea una secuencia texto de salida formada por oraciones bimembres, de la siguiente forma:
El sujeto, de la secuencia 1, y el predicado, de la secuencia 2.
Al finalizar informar cuantas oraciones tiene cada secuencia.

ACCION ejercicio20 ES
	AMBIENTE
		sec1,sec2,sec_sal: secuencia de caraacteres;
		v1,v2: caraacter;
		cantor1, cantor2: entero;
	
	PROCESO
		ARR(sec1);
		AVZ(sec1,v1);
		ARR(sec2);
		AVZ(sec2,v2);
		CREAR(sec_sal);

		cantor1:= 0;
		cantor2:= 0;

		MIENTRAS NFDS(sec1) Y NFDS(sec2) HACER:
			cantor1:= cantor1 + 1;
			MIENTRAS (v1 <> ".") HACER:
				REPETIR
					GRABAR(sec_sal,v1);
					AVZ(sec1,v1);
				HASTA QUE (v1 = ","); //copio hasta la copa

				cantor2:= cantor2 + 1;
				MIENTRAS (v2 <> ".") HACER:
					MIENTRAS (v2 <> ",") HACER:
						AVZ(sec2,v2);
					Fin_mientras;
					
					AVZ(sec2,v2); //avanzo la coma

					REPETIR 
						GRABAR(sec_sal,v2);
						AVZ(sec2,v2),
					HASTA QUE (v = ".");
					AVZ(sec2,v2);
				Fin_mientras;
				AVZ(sec1,v1);
			Fin_mientras;
		Fin_mientras;

		CERRAR(sec1);
		CERRAR(sec2);
		CERRAR(sec_sal);

		Escribir("La secuencia 1 tiene:",cantor1,"oraciones");
		Escribir("La secuencia 2 tiene:",cantor2,"oraciones");
Fin_accion.

//Ejercicio21 -> idem anterior pero al revés

//Ejercicio22
La empresa Ideas Argentinas S.A. posee datos de sus empleados en dos secuencias de caracteres; la primera secuencia (Sec1) formada por los nombres (uno por persona) de los empleados separados por blancos y la segunda secuencia (Sec2) posee los números de documento de cada uno de los empleados (palabras de 8 dígitos numéricos). Ambas secuencias poseen la misma cantidad de datos, es decir al primer nombre de la primera secuencia le corresponde el primer número de documento de la segunda secuencia y así sucesivamente. La secuencia de números de documentos no posee espacios en blanco ni ningún otro tipo de caracter que separe un número de documento de otro.
A la empresa le interesa tener en una nueva secuencia (Sec3) los datos de todas aquellas personas que cumplan la condición de que el nombre se encuentre en una posición impar. La nueva secuencia debe generarse de la siguiente manera: el número de documento seguido (sin espacios) por una coma y luego (sin espacios) por el nombre y seguido al nombre un #.

ACCION ejercicio22 ES
	AMBIENTE
		sec1,sec2,sec3: secuencia de caraacteres;
		v1,v2: caracter;
		contpos,i: entero;

	PROCESO
		ARR(sec1);
		AVZ(sec1,v1);
		ARR(sec2);
		AVZ(sec2,v2);
		CREAR(sec3);

		contpos:= 0;

		MIENTRAS NFDS(sec1) HACER:
			contpos:= contpos + 1;
			MIENTRAS (v1 <> " ") HACER:  //los nombres estan separados por blancos	
				SI (contpos MOD 2 <> 0) ENTONCES:
					PARA i:=1 a 8 HACER:  //copio el dni
						GRABAR(sec3,v2);
						AVZ(sec2,v2);
					Fin_para;
					GRABAR(sec3,",");

					MIENTRAS (v1 <> " ") HACER:
						GRABAR(sec3,v1);
						AVZ(sec1,v1);
					Fin_mientras;
					GRABAR(sec3,"#");
				Fin_si; 
				AVZ(sec1,v1):
			Fin_mientras;
		Fin_mientras;

		CERRAR(sec1);
		CERRAR(sec2);
		CERRAR(sec3);
Fin_accion.

//Ejercicio23
Realice un algoritmo para el enunciado del ejercicio 1.22, pero considerando que los datos que se deben copiar en la Sec3 son los de aquellas personas que cumplan la condición de que: el número de documento comienza con un número impar. 

//Ejercicio24
Realice un algoritmo para el enunciado del ejercicio 1.22, pero considerando que los datos que se deben copiar en la Sec3 son los de aquellas personas que cumplan la condición de que: el nombre no comienza con una vocal.

//Ejercicio25 
Dada una secuencia texto de entrada que contiene palabras alfanuméricas, escribir un algoritmo que genere dos secuencias de salida. Una de ellas contendrá solo las palabras de la secuencia original que comienzan con vocal, en las cuales se reemplazarán todas las vocales por ‘#’, por ejemplo: entrada 'avión1', salida '#v##n1' y la otra será una secuencia numérica en la que se almacenarán las cantidades de vocales que se encontraron en cada una de las palabras que cumplieron la condición. Por final de proceso se deberá informar el promedio de palabras por oración.

ACCION ejercicio25 ES
	AMBIENTE
		sec, sec1: secuencia de alfanumericos.
		v: alfanumerico;
		sec2: secuencia de enteros;
		contv,contp,contor: entero;
		vocal= {"a","e","i","o","u"};

	PROCESO
		ARR(sec);
		AVZ(sec,v);

		contv:=  0;
		contp:= 0;
		contor:= 0;

		MIENTRAS NFDS(sec) HACER:
			contor:= contor + 1; //cuento oracion
			MIENTRAS (v <> ".") HACER:
				contp:= contp + 1; //cuento palabra
				MIENTRAS (v <> " ") HACER:
					SI (v en vocal) ENTONCES:
						GRABAR(sec1,"#");  //reemplazo las vocales por # en salida (sec1)
						contv:= contv + 1;  //cuento vocales por palabra
					Fin_si;
					AVZ(sec,v);
				Fin_mientras;
				GRABAR(sec2,contv);  //grabo la cantidad de vocales por palabra en salida (sec2)
				AVZ(sec,v);
			Fin_mientras;	
			AVZ(sec,v);
		Fin_mientras;

		CERRAR(sec);
		CERRAR(sec1);
		CERRAR(sec2);

		Escribir("El promedio de palabras por oración es:", contp/contor);
Fin_accion.

//Ejercicio26
Se posee 2 secuencias (S1 y S2) con las cuales se desea generar una nueva secuencia (SAL) donde se intercalen las palabras de las secuencias de entrada, de la siguiente manera: copiar de S1 aquellas palabras que empiezan y terminan con la misma letra y de S2 aquellas palabras que posean al menos un digito numérico y además estén en posición par.

ACCION ejercicio26 ES
	AMBIENTE
		s1: secuencia de caraacteres;
		s2: secuencia de alfanumericos;
		sal: secuencia de alfanumericos;

		v1,car: caracter;
		v2: alfanumerico;
		
