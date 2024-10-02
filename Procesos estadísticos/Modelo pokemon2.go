ACCION modelo_pokemon2 (region: arreglo de [1..31] de alfanumerico; categ: arreglo de [1..16] de alfanumerico) ES
	AMBIENTE
		papers = registro
			dni:
			apellynom:
			mail:
			region:
			categ:
			titulo:
		Fin_reg;
		arch: archivo de papers;
		reg: papers;

		arr: arreglo de [1..32,1..17] de entero;
		j,k,mayor_reg: entero;

		PROCEDIMIENTO iniciar()
			PARA j:=1 a 32 HACER:
				PARA k:=1 a 17 HACER:
					arr[j,k]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;

	PROCESO
		ABRIR E/(arch);
		LEER(arch,reg);

		iniciar();
		mayor_reg:=0;

		MIENTRAS NFDA(arch) HACER:
			j:= reg.region;
			k:= reg.categ;

			arr[j,k]:= arr[j,k] + 1;
			arr[32,k]:= arr[31,j] + 1; //total por regional (fila)
			arr[j,17]:= arr[j,17] + 1; //total por categoría (columna)
			arr[32,17]:= arr[32,17] + 1; //total gral

			LEER(arch,reg);
		Fin_mientras;

		//para los totales por regionales (tot x fila) recorro las filas de la última columna
		Esc("Totales por regionales:");
		PARA j:=1 a 31 HACER:
			Esc("Regional:",region[j]); //uso el arreglo de entrada para mostrar el nombre
			Esc("Papers enviados:",arr[j,17]);
		Fin_para;

		//para los totales por categoría (tot x columna) recorro las columnas de la última fila
		Esc("Totales por categorías:");
		PARA k:=1 a 16 HACER:
			Esc("Categoría:",categ[k]);
			Esc("Papers enviados:",arr[32,k]);
		Fin_para;

		//yo obtengo la regional con más papers enviados una vez que termino de cargar el arreglo de salida
		PARA j:=1 a 31 HACER:
			SI (arr[j,17] > mayor_reg) ENTONCES: //recorro por fila la última columna donde están mis totales por regional
				mayor_reg:= arr[32,j];
			Fin_si;
		Fin_para; 

		//para sacar las otras regionales que tienen igual cantidad de papers que la mayor, comparo y vuelvo a recorrer el arreglo en la última columna
		Esc("Regional/es con la mayor cantidad de papers enviados:");
		PARA j:=1 a 31 HACER:
			SI (mayor_reg = arr[j,17]) ENTONCES:
				Esc(region[j]); //utilizo mi arreglo de entrada para mostrar el nombre de la regional
			Fin_si;
		Fin_para;
		
		Esc("Total general de papers enviados:",arr[32,17]);
		CERRAR(arch);
Fin_accion.



