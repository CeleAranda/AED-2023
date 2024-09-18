ACCION modelo2022 ES
	AMBIENTE
		fecha = registro
			año: N(4);
			mes: 1..12;
			dia: 1..31;
		Fin_reg;

		album = registro
			cod_fig: AN(30);
			cant: N(9);
			duplicado: {"si","no"};
		Fin_reg;
		arch_a,arch_sal: archivo de album ordenado por cod_fig;
		a,a_sal: album;

		intercambio = registro
			cod_fig: AN(30);
			cod_ami: AN(30);
			fecha_sol: fecha;
		Fin_reg;
		arch_int: archivo de intercambio ordenado por cod_fig y cod_ami;
		inter: intercambio;

		PROCEDIMIENTO LeerA() 
			Leer(arch_a,a);

			SI FDA(arch_a) ENTONCES:
				a.cod_fig:= HV;
			Fin_si;
		Fin_proced;

		PROCEDIMIENTO LeerB() 
			Leer(arch_int,inter);

			SI FDA(arch_int) ENTONCES:
				inter.cod_fig:= HV;
			Fin_si;
		Fin_proced;

	PROCESO
		ABRIR E/(arch_a);
		ABRIR E/(arch_int);

		leerA();
		leerB();
		
		cant_dupl:= 0;

		MIENTRAS (a.cod_fig <> HV) O (inter.cod_fig <> HV) HACER:
			SI (a.cod_fig < inter.cod_fig) ENTONCES:
				a_sal:= a;
				GRABAR(arch_sal,a_sal);
				leerA();
				leerB();
			SINO
				SI (a.cod_fig = inter.cod_fig) ENTONCES:
					SI (a.duplicado = "si") ENTONCES:		
						cant_dupl:= cant_dupl + 1;
						MIENTRAS (a.cod_fig = inter.cod_fig) HACER:
							a_sal.cant:= a_sal.cant + 1;
							GRABAR(arch_sal,a_sal);
							LeerB();
						Fin_mientras;
						GRABAR(arch_sal,a_sal):
						LeerA();
					SINO
						Esc("¡La figurita no permite duplicados!");
					Fin_si;
				SINO
					verif:= diff_dias(6, inter.fecha_sol); 
					SI (verif = "falso") ENTONCES:
						a_sal.cod_fig:= inter.cod_fig;
						a_sal.cant:= 1;
						a_sal.duplicado:= "no";

						SI (a.duplicado = "si") ENTONCES:		
							cant_dupl:= cant_dupl + 1;
							MIENTRAS (a.cod_fig = inter.cod_fig) HACER:
								a_sal.cant:= a_sal.cant + 1;
								GRABAR(arch_sal,a_sal);
								LeerB();
							Fin_mientras;
							GRABAR(arch_sal,a_sal):
							LeerA();
						Fin_si;
					SINO 
						Esc("¡La fecha de solicitud ha expirado!");
					Fin_si;
				Fin_si;
			Fin_si;
		Fin_mientras;
		CERRAR(arch_a);
		CERRAR(arch_sal);
		CERRAR(arch_int);

		Esc("Cantidad de figuritas duplicadas que se admitieron al álbum:",cant_dupl);
Fin_accion.

