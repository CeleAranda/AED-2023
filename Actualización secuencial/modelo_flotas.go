//act unitaria
ACCION modelo_flota ES
	flotas = registro
		cod_flota: AN(30);
		vigente: {"si","no"};
		cod_estado: {1,2,3,4};
		km: N(9,2);
	Fin_reg;
	arch_flota, arch_sal: archivo de flotas ordenado por cod_flota;
	f, f_sal: flotas;
	
	flotas_mov = registro;
		cod_flota: AN(30);
		km: N(9,2);
		cod_estado: {1,2,3,4};
	Fin_reg;
	arch_flota_mov: archivo de flotas_mov ordeando por cod_flota;
	m: flotas_mov;

	PROCEDIMIENTO leerA()
		Leer(arch_flota,f);

		SI FDA(arch_flota) ENTONCES:
			f.cod_flota:= HV;
		Fin_si;
	Fin_proced;

	PROCEDIMIENTO leerB()
		Leer(arch_flota_mov,m);

		SI FDA(arch_flota_mov) ENTONCES:
			m.cod_flota:= HV;
		Fin_si;
	Fin_proced;

	cont1,cont2,cont3,cont4,conta: entero;

PROCEDIMIENTO
	ABRIR E/(arch_flota);
	ABRIR E/(arch_flota_mov);
	ABRIR /S(arch_sal);

	leerA();
	leerB();

	cont1:= 0;
	cont2:= 0;
	cont3:= 0;
	cont4:= 0;
	conta:= 0;

	MIENTRAS (f.cod_flota <> HV) O (m.cod_flota <> HV) HACER:
		conta:= conta + 1;
		SI (f.cod_flota < m.cod_flota) ENTONCES:
			SI (m.cod_estado = 1) ENTONCES:
				Esc("¡Error! Alta de avión inválida.");
			SINO
				SI (m.cod_estado = 2) ENTONCES:
					Esc("¡No hay modificaciones en los Km recorridos!");
				SINO	
					SI (m.cod_estado = 3) ENTONCES:
						Esc("¡El avión no se encuentra en mantenimiento!");
					SINO
						Esc("¡Error! Avión operable.");
					Fin_si;
				Fin_si;
			Fin_si;
			f_sal:= f;
			GRABAR(arch_sal;f_sal);
			leerA();
		SINO
			SI (f.cod_flota = m.cod_flota) ENTONCES:
				SI (m.cod_estado = 1) ENTONCES:
					Esc("¡Error! Alta de avión inválida.");
				SINO
					SI (m.cod_estado = 2) ENTONCES:
						Esc("Avión en servicio. Actualizando Km recorridos...");
						f_sal.km:= f_sal.km + m.km;
						cont2:= cont2 + 1;
					SINO	
						SI (m.cod_estado = 3) ENTONCES:
							Esc("¡El avión se encuentra en mantenimiento!");
							cont3:= cont3 + 1;
						SINO
							Esc("¡Error! Avión inoperable.");
							f_sal.vigente:= "no";
							cont4:= cont4 + 1;
						Fin_si;
					Fin_si;
				Fin_si;
				GRABAR(arch_sal,f_sal);
				leerA();
				leerB();
			SINO
				SI (m.cod_estado = 1) ENTONCES:
					f_sal.vigente:= "si";
					f_sal.cod_estado:= m.cod_estado;
					f_sal.km:= 0;
					GRABAR(arch_sal,f_sal);
					cont1:= cont1 + 1;
				SINO
					SI (m.cod_estado = 2) ENTONCES:
						Esc("¡Error! No se pueden modificar los Km recorridos!");
					SINO	
						SI (m.cod_estado = 3) ENTONCES:
							Esc("¡Error!");
						SINO
							Esc("¡Error!");
						Fin_si;
					Fin_si;
				Fin_si;
				GRABAR(arch_sal,f_sal);
				leerB();		
			Fin_si;
		Fin_si;
	Fin_mientras;
	CERRAR(arch_flota);
	CERRAR(arch_flota_mov);
	CERRAR(arch_sal);

	Esc("Total de aviones de estado 1:",cont1);
	Esc("Total de aviones de estado 2:",cont2);
	Esc("Total de aviones de estado 3:",cont3);
	Esc("Total de aviones de estado 4:",cont4);

	Esc("Porcentaje de aviones de estado 1:",(cont1/conta)*100,"%");
	Esc("Porcentaje de aviones de estado 2:",(cont2/conta)*100,"%");
	Esc("Porcentaje de aviones de estado 3:",(cont3/conta)*100,"%");
	Esc("Porcentaje de aviones de estado 4:",(cont4/conta)*100,"%");
Fin_accion.




//act por lotes
ACCION modelo_flota ES
	flotas = registro
		cod_flota: AN(30);
		vigente: {"si","no"};
		cod_estado: {1,2,3,4};
		km: N(9,2);
	Fin_reg;
	arch_flota, arch_sal: archivo de flotas ordenado por cod_flota;
	f, f_sal: flotas;
	
	flotas_mov = registro;
		cod_flota: AN(30);
		km: N(9,2);
		cod_estado: {1,2,3,4};
	Fin_reg;
	arch_flota_mov: archivo de flotas_mov ordeando por cod_flota;
	m: flotas_mov;

	PROCEDIMIENTO leerA()
		Leer(arch_flota,f);

		SI FDA(arch_flota) ENTONCES:
			f.cod_flota:= HV;
		Fin_si;
	Fin_proced;

	PROCEDIMIENTO leerB()
		Leer(arch_flota_mov,m);

		SI FDA(arch_flota_mov) ENTONCES:
			m.cod_flota:= HV;
		Fin_si;
	Fin_proced;

	cont1,cont2,cont3,cont4,conta: entero;

PROCEDIMIENTO
	ABRIR E/(arch_flota);
	ABRIR E/(arch_flota_mov);
	ABRIR /S(arch_sal);

	leerA();
	leerB();

	cont1:= 0;
	cont2:= 0;
	cont3:= 0;
	cont4:= 0;
	conta:= 0;

	MIENTRAS (f.cod_flota <> HV) O (m.cod_flota <> HV) HACER:
		conta:= conta + 1;
		SI (f.cod_flota < m.cod_flota) ENTONCES:
			SI (m.cod_estado = 1) ENTONCES:
				Esc("¡Error! Alta de avión inválida.");
			SINO
				SI (m.cod_estado = 2) ENTONCES:
					Esc("¡No hay modificaciones en los Km recorridos!");
				SINO	
					SI (m.cod_estado = 3) ENTONCES:
						Esc("¡El avión no se encuentra en mantenimiento!");
					SINO
						Esc("¡Error! Avión operable.");
					Fin_si;
				Fin_si;
			Fin_si;
			f_sal:= f;
			GRABAR(arch_sal;f_sal);
			leerA();
		SINO
			SI (f.cod_flota = m.cod_flota) ENTONCES:
				MIENTRAS (f.cod_flota = m.cod_flota) HACER:
					SI (m.cod_estado = 1) ENTONCES:
						Esc("¡Error! Alta de avión inválida.");
					SINO
						SI (m.cod_estado = 2) ENTONCES:
							Esc("Avión en servicio. Actualizando Km recorridos...");
							f_sal.km:= f_sal.km + m.km;
							cont2:= cont2 + 1;
						SINO	
							SI (m.cod_estado = 3) ENTONCES:
								Esc("¡El avión se encuentra en mantenimiento!");
								cont3:= cont3 + 1;
							SINO
								Esc("¡Error! Avión inoperable.");
								f_sal.vigente:= "no";
								cont4:= cont4 + 1;
							Fin_si;
						Fin_si;
					Fin_si;
					leerB();
				Fin_mientras;
				GRABAR(arch_sal,f_sal);
				leerA();
			SINO
				SI (m.cod_estado = 1) ENTONCES:
					f_sal.vigente:= "si";
					f_sal.cod_estado:= m.cod_estado;
					f_sal.km:= 0;
					cont1:= cont1 + 1;

					MIENTRAS (f.cod_flota = m.cod_flota) HACER:
						SI (m.cod_estado = 1) ENTONCES:
							Esc("¡Error! Alta de avión inválida.");
						SINO
							SI (m.cod_estado = 2) ENTONCES:
								Esc("Avión en servicio. Actualizando Km recorridos...");
								f_sal.km:= f_sal.km + m.km;
								cont2:= cont2 + 1;
							SINO	
								SI (m.cod_estado = 3) ENTONCES:
									Esc("¡El avión se encuentra en mantenimiento!");
									cont3:= cont3 + 1;
								SINO 
									Esc("¡Error! Avión inoperable.");
									f_sal.vigente:= "no";
									cont4:= cont4 + 1;
								Fin_si;
							Fin_si;
						Fin_si;
						leerB();
					Fin_mientras;
				SINO
					SI (m.cod_estado = 2) ENTONCES:
						Esc("¡Error! No se pueden modificar los Km recorridos!");
					SINO	
						SI (m.cod_estado = 3) ENTONCES:
							Esc("¡Error!");
						SINO
							Esc("¡Error!");
						Fin_si;
					Fin_si;
				Fin_si;
				GRABAR(arch_sal,f_sal);		
			Fin_si;
		Fin_si;
	Fin_mientras;
	CERRAR(arch_flota);
	CERRAR(arch_flota_mov);
	CERRAR(arch_sal);

	Esc("Total de aviones de estado 1:",cont1);
	Esc("Total de aviones de estado 2:",cont2);
	Esc("Total de aviones de estado 3:",cont3);
	Esc("Total de aviones de estado 4:",cont4);

	Esc("Porcentaje de aviones de estado 1:",(cont1/conta)*100,"%");
	Esc("Porcentaje de aviones de estado 2:",(cont2/conta)*100,"%");
	Esc("Porcentaje de aviones de estado 3:",(cont3/conta)*100,"%");
	Esc("Porcentaje de aviones de estado 4:",(cont4/conta)*100,"%");
Fin_accion.