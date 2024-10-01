ACCION modelo_pokemon (arr: arreglo de [1..151] de alfanumerico) ES
	AMBIENTE
		clave = registro	
			cod_region:
			cod_us:
		Fin_reg;

		CAPTURAS = registro
			clavemov: clave;
			cod_pok:
			puntos_ex:
			fecha_cap: fecha;
			estado_pok: {"E","I","D"}; //entrenándose, incubándose, descansando
			estado_us: {"A","S","B"}; //activo, suspendido, baneado
		Fin_reg;
		arch_c: archivo de CAPTURAS ordenado por clavem y estado_pok;
		c: CAPTURAS;

		USUARIOS = registro
			clavem: clave;
			correo:
			experiencia:
		Fin_reg;
		arch_us, arch_us_sal: archivo de USUARIOS ordenado por clavem;
		u, u_sal: USUARIOS;

		PROCEDIMIENTO leerA() 
			LEER(arch_us,u);

			SI FDA(arch_us) ENTONCES:
				u.clavem.clave.cod_region:= HV;
				u.clavem.clave.cod_us:= HV;
			Fin_si;
		Fin_proced;

		PROCEDIMIENTO leerB() 
			LEER(arch_c,c);

			SI FDA(arch_c) ENTONCES:
				c.clavemov.clave.cod_region:= HV;
				c.clavemov.clave.cod_us:= HV;
				c.cod_pok:= HV;
			Fin_si;
		Fin_proced;

		PROCEDIMIENTO pokemon_d()
			SI (arr[c.cod_pok] = pokemon) ENTONCES:
				pokemon:= arr[c.cod_pok];
				cont_us: cont_us + 1; //cuento los usuarios que tienen ese pokemon descansando 
			Fin_si;
		Fin_proced;

	PROCESO
		ABRIR E/(arch_us);
		ABRIR E/(arch_c);
		ABRIR /S(arch_us_sal);

		leerA();
		leerB();

		pokemon:= arr[1];
		cont_us:=0;

		MIENTRAS (u.clavem <> HV) O (c.clavemov <> HV) HACER:
			SI (u.clavem < c.clavemov) ENTONCES:
				u_sal:= u;
				GRABAR(arch_us_sal,u_sal);
				leerA();
			SINO
				SI (u.clavem = c.clavemov) ENTONCES:
					MIENTRAS (u.clavem = c.clavemov) HACER:
						SI (c.estado_us = "A") ENTONCES:
							SI (c.estado_pok = "E") ENTONCES:
								u_sal.experiencia:= u_sal.experiencia + (c.puntos_ex * 2);
							SINO
								SI (c.estado_us = "I") ENTONCES:
									u_sal.experiencia:= u_sal.experiencia + c.	puntos_ex;
								SINO 
									u_sal.experiencia:= u_sal.experiencia + c.puntos_ex;
									pokemon_d();
								Fin_si;
							Fin_si;
						SINO
							SI (c.estado_us = "S") ENTONCES:
								Esc("¡Usuario eliminado!");
							SINO
								Esc("¡Usuario baneado!");
							Fin_si;
						Fin_si;
						GRABAR(arch_us_sal,u_sal);
						leerB();
					Fin_mientras;
					leerA()
				SINO
					SI (c.estado_us = "A") ENTONCES:
						SI (c.estado_pok = "E") ENTONCES:
							u_sal.correo:= " ";
							u_sal.experiencia:= u_sal.experiencia + (c.puntos_ex * 2);
						SINO
							SI (c.estado_us = "I") ENTONCES:
								u_sal.experiencia:= u_sal.experiencia + c.	puntos_ex;
							SINO 
								u_sal.experiencia:= u_sal.experiencia + c.puntos_ex;
								pokemon_d();
							Fin_si;
						Fin_si;

						MIENTRAS (u.clavem = c.clavemov) HACER:
							SI (c.estado_pok = "E") ENTONCES:
								u_sal.correo:= " ";
								u_sal.experiencia:= u_sal.experiencia + (c.puntos_ex * 2);
							SINO
								SI (c.estado_us = "I") ENTONCES:
									u_sal.experiencia:= u_sal.experiencia + c.	puntos_ex;
								SINO 
									u_sal.experiencia:= u_sal.experiencia + c.puntos_ex;
									pokemon_d();
								Fin_si;
							Fin_si;
							GRABAR(arch_us_sal,u_sal);
							leerB();
					SINO
						SI (c.estado_us = "S") ENTONCES:
							Esc("¡El ususario no está dado de alta!");
						SINO
							Esc("¡El ususario no está dado de alta!");
						Fin_si;
					Fin_si;
				Fin_si;
			Fin_si;
		Fin_mientras;
		CERRAR(arch_c);
		CERRAR(arch_us);
		CERRAR(arch_us_sal);

		Esc("Nombre del pokemon en estado 'Descansando' que lo tienen la mayoría de los usuarios:",pokemon);
		Esc("Cantidad de usuarios que lo presentan:",cont_us);
Fin_accion.
					