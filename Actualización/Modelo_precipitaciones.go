ACCION modelo_parcial (arr: arreglo de [1..25] de entero) ES
	AMBIENTE
		fecha = registro
			año: N(4);
			mes: 1..12;
			dia: 1..31;
		Fin_reg;

		clave = registro
			id_pluv: N(9);
			fecha_reg: fecha;
		Fin_reg;

		precipitaciones = registro
			clave_mae: clave;
			precip: N(9)
			estado_pluv: {"ac","ma"}; //activo, mantenimiento
		Fin_reg;
		Arch_mae, Arch_mae_sal: archivo de precipitaciones ordenado por clave_mae;
		p,p_sal: precipitaciones;

		mediciones = registro
			clave_mov: clave;
			hora: N(4);
			estado_pluv: {"ac","ma"};
			precip: N(9);
		Fin_reg;
		Arch_mov: archivo de mediciones ordeando por clave_mov;
		m: mediciones;

		pluviometros = registro
			id_pluv: N(9);
			modelo: AN(30);
			desc: AN(30);
			depto: 1..25;
		Fin_reg;
		Arch_pluv: archivo de pluviometros indexado por id_pluv;
		reg: pluviometros;

		i:entero;

		PROCEDIMIENTO leerM() 
			LEER(Arch_mae,p);

			SI FDA(Arch_mae) ENTONCES:
				p.clave_mae:= HV;
			Fin_si;
		Fin_proced;

		PROCEDIMIENTO leerMov() 
			LEER(Arch_mov,m);

			SI FDA(Arch_mov) ENTONCES:
				m.clave_mov:= HV;
			Fin_si;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch_mae);
		ABRIR E/(Arch_mov);
		ABRIR E/(Arch_pluv);

		LeerM();
		LeerMov();
		cont_pluv:=0;

		MIENTRAS (m.clave_mae <> HV) o (p.clave_mov <> HV) HACER:
			reg.id_pluv:= p.clave_mae.clave.id_pluv;
			LEER(Arch_pluv,reg);
			SI EXISTE ENTONCES:
				SI (p.estado_pluv = "ac") ENTONCES:
					SI (p.clave_mae < m.clave_mov) ENTONCES:
						p_sal:= p;
						GRABAR(Arch_mae_sal,p_sal);
						leerM();
						leerMov();
					SINO	
						SI (p.clave_mae = m.clave_mov) ENTONCES:
							MIENTRAS (p.clave_mae = m.clave_mov) HACER:
								//asumo que tengo que modificar la cantidad de precipitaciones y el estado del pluviómetro 
								p_sal.precipitaciones:= m.precipitaciones;
								p_sal.estado_pluv:= m.estado_pluv;
								LeerMov();
							Fin_mientras;
							GRABAR(Arch_mae_sal,p_sal);
							LeerM();

						SINO
							//alta, agrego el pluviómetro
							p_sal.clave_mae.clave.id_pluv:= m.clave_mov.clave.id_pluv;
							p_sal.clave_mae.clave.fecha_reg:= m.clave_mov.clave.fecha_reg;
							p_sal.precip:= m.precip;
							p_sal.estado_pluv:= m.estado_pluv;

							MIENTRAS (p.clave_mae = m.clave_mov) HACER:
								//asumo que tengo que modificar la cantidad de precipitaciones y el estado del pluviómetro 
								p_sal.precipitaciones:= m.precipitaciones;
								p_sal.estado_pluv:= m.estado_pluv;
								LeerMov();
							Fin_mientras;
							GRABAR(Arch_mae_sal,p_sal); //en la alta no se lee el maestro porque no hay nada que leer!!!!!
				SINO
					cont_pluv:= cont_pluv + 1;
					Esc("El pluviómetro:",reg.id_pluv,"está en el departamento:",arr[reg.depto]);
			SINO
				Esc("Este pluviómetro no existe");
			Fin_si;
		Fin_mientras;

		CERRAR(Arch_mae);
		CERRAR(Arch_mov);
		CERRAR(Arch_pluv);
		CERRAR(Arch_mae_sal);

		Esc("La cantidad de pluviómetros en mantenimiento es:",cont_pluv);
Fin_accion.


