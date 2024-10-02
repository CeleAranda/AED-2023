ACCION modelo_flota2 ES
	AMBIENTE
		fecha = registro
			año:
			mes:
			dia:
		Fin_reg;

		FLOTAS = registro
			cod_flota:
			vigente: {"si","no"};
			cod_estado:
			km:
		Fin_reg;
		arch: archivo de FLOTAS ordenado por cod_flota;
		reg: FLOTAS;

		AVIONES = registro	
			cod_flota:
			cod_mod:
			descr:
			fecha_compra: fecha;
		Fin_reg;
		arch2: archivo de AVIONES indexado por cod_flota;
		reg2: AVIONES;

		arr: arreglo de [1..3,1..6] de real;
		i,j: entero;

	PROCESO
		ABRIR E/(arch);
		LEER(arch,reg);
		ABRIR E/(arch2);

		MIENTRAS NFDA(arch) HACER:
			SI (reg.km > 25000) ENTONCES:
				reg2.cod_flota:= reg.cod_flota;
				LEER(arch2,reg2);

				SI EXISTE:
					SEGUN reg.vigente HACER:
						"si": i:=1;
						"no": i:=2;
					Fin_segun;
					j:= reg2.cod_mod;
					
					arr[i,j]:= arr[i,j] + 1;
					arr[6,j]:= arr[6,j] + 1; //total por vigente (fila)
					arr[i,3]:= arr[i,3] + 1; //total por modelo (columna)
					arr[6,3]:= arr[6,3] + 1;
				SINO
					Esc("El avión ingresado no se encuentra en el archivo");
				Fin_si;
			Fin_si;
			LEER(arch,reg);
		Fin_mientras;

		//totales x fila
		PARA i:=1 a 2 HACER:
			Esc("Vigente:",i);
			Esc("Cantidad de flotas:",arr[i,6]);
		Fin_para;

		//totales x columna
		PARA j:=1 a 5 HACER:
			Esc("Modelo:",j);
			Esc("Cantidad de flotas:",arr[3,j]);
		Fin_para;

		Esc("Total general de flotas:",arr[3,6]);

		CERRAR(arch);
		CERRAR(arch2);
Fin_accion.

					
			