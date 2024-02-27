arreglos

ACCION ejercicio3.4 (arr: arreglo de [1..50] de enteros) ES
	AMBIENTE
		arr_sal: arreglo de [1..50] de entero;
		cont, i, j, num: entero;

	PROCESO
		cont := 0;
		j := 1;
		
		PARA i := 1 a 50 HACER
			SI (arr[i] MOD 3 = 0) ENTONCES
				cont := cont + 1;
			SINO
				res_num: arr[i];
				arr_sal[j] := arr[i] * 3;
				j := j + 1;  //se utiliza la variable j para guardar los números de manera contigua y que no queden espacios vacíos en el arreglo de salida
			fin_si;
		fin_para;

		Escribir("La cantidad de números que cumplieron con la condición es de:", cont);
	Fin_accion.




ACCION ejercicio3.5 (A: arreglo de [1..30] de reales, B: arreglo de [1..30] de reales) ES
	AMBIENTE	
		C: arreglo de [1..60] de reales;
		i, j, k: entero

	PROCESO	
		i:=0;
		j:=0;
		k:=0;

		MIENTRAS (i <= 30 y j <= 30) HACER	
			SI (A[i] <= B[i]) ENTONCES
				C[i] := A[i];
				i := i + 1;
			SINO	
				C[i] := B[i];
				j := j + 1;
			fin_si;

			k:= k + 1;
		fin_mientras;

		MIENTRAS (i <= 30) HACER
			C := A[i];
			i := i + 1;
			k:= k + 1;
		fin_mientras;

		MIENTRAS (j <= 30) HACER	
			C := A[i];
			j := j + 1;
			k:= k + 1;
		fin_mientras;
Fin_accion.
