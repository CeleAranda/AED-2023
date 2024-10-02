ACCION modelo2022 ES
	AMBIENTE
		album = registro
			cod_us: 
			cod_fig:
			cant:
			tipo: {"D","C","V"}; //dorado, común, virtual
		Fin_reg;
		archa: archivo de album ordenado por cod_us;
		rega: album

		amigos = registro
			cod_us:
			apell:
			nomb:
			cel:
		Fin_reg;
		arch_amig: archivo de amigos indexado por cod_us;
		reg_amig: amigos;

		arr: arreglo de [1..11,1..4] de entero;
		i,j: entero;
 
	PROCESO
		ABRIR E/(archa);
		LEER(archa,rega);
		ABRIR E/(arch_amig);

		MIENTRAS NFDA(archa) HACER:
			i:= rega.cod_us;
			SEGUN rega.tipo HACER:
				"D": j:= 1;
				"C": j:=2;
				"V": j:=3;
			Fin_segun;

			arr[i,j]:= arr[i,j] + reg.cant;
			arr[i,4]:= arr[i,4] + reg.cant;
			arr[11,j]:= arr[11,j] + reg.cant;
			arr[11,4]:= arr[11,4] + reg.cant;

			LEER(archa);
		Fin_mientras;

		//totales por usuarios (x fila)
		PARA i:=1 a 10 HACER:
			Esc("Usuario:",i);
			Esc("Cantidad de figuritas:",arr[i,4]);
			Esc("Porcentaje de figuritas descubiertas:",arr[i,4]/arr[11,4],"%");
		Fin_para;

		//totales por tipo (x columna)
		PARA j:=1 a 3 HACER:
			Esc("Tipo de album:",j);
			Esc("Cantidad de figuritas:",arr[11,j]);
		Fin_para;

		//usuario con más figuritas
		PARA i:=1 a 10 HACER:
			SI (arr[i,4] > mayor_us) ENTONCES:
				mayor_us:= arr[i,4];
				us:= rega.cod_us;
			Fin_si;
		Fin_para;

		reg_amig.cod_us:= us;
		LEER(arch_amig,reg_amig);
		SI EXISTE:
			Esc("Usuario que coleccionó más figuritas:",us);
			Esc("Nombre:",reg_amig.nomb);
			Esc("Apellido:",reg_amig.apell);
		Fin_si;
		
		CERRAR(arch_amig);
		CERRAR(archa);
Fin_accion.