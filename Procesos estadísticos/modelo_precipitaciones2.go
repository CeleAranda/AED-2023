ACCION ejercicio2 ES
	AMBIENTE
		fecha = registro
			año: N(4);
			mes: 1..12;
			dia: 1..31;
		Fin_reg;

		precip_anuales = registro
			id_pluv: AN(30);
			fecha_reg: fecha;
			precip: N(9,2);
			estado: {"ac","ma"};
		Fin_reg;

		arch: archivo de precip_anuales ordenado por id_pluv y fecha_reg;
		reg: precip_anuales;

		pluviometros = registro
			id_pluv: AN(30);
			modelo: AN(30);
			desc: AN(30);
			dpto: 1..25;
		Fin_reg;

		arch2: archivo de pluviometros indexado por id_pluv;
		reg2: pluviometros;

		arr: arreglo de [1..13,1..26] de reales;
		i,j: entero;
		menor_dpto,menor_mes,mayor_mes: entero;

		PROCEDIMIENTO inicializar()
			PARA i:=1 a 13 HACER:
				PARA j:=1 a 26 HACER:
					arr[i,j]:= 0;
				Fin_para;
			Fin_para;
		Fin_proced;


	PROCESO
		ABRIR E/(arch);
		LEER(arch,reg);
		ABRIR E/(arch2);

		inicializar();

		MIENTRAS NFDA(arch) HACER:
			reg2.id_pluv:= reg.id_pluv;
			LEER(arch2,reg2);

			SI EXISTE
				i:= reg.fecha_reg.fecha.mes;
				j:= reg2.dpto;
				
				arr[i,j]:= arr[i,i] + reg.precip;
					
					SI (arr[i,j] < menor_mes) ENTONCES:
						menor_mes:= i;
						menor_dpto:= j;
					Fin_si;

				arr[i,26]:= arr[i,26] + reg.precip; //total por mes (tot_fila)
				arr[13,j]:= arr[13,j] + reg.precip; //total por dpto (tot_columna)
				arr[13,26]:= arr[13,26] + reg.precip; //total gral
				
				SI (arr[i,26] > mayor_mes) ENTONCES: //mes con mayor precipitaciones
					mayor_mes:= i;
				Fin_si;

				SI (arr[13,j] > 350) ENTONCES: //mes con mayor precipitaciones
					Esc("En el mes:",i,", el departamento:",j,", superó los 350 mm");
				Fin_si;

			SINO
				Esc("¡El pluviómetro no existe!");	
			Fin_si;
		
			LEER(arch,reg);
		Fin_mientras;
		Cerrar(arch);
		Cerrar(arch2);

		Esc("Mes con mayores precipitaciones:",mayor_mes);
		Esc("Menor cantidad de precipitaciones en el mes:",menor_mes,"y en el departamento:",menor_dpto);
Fin_accion.