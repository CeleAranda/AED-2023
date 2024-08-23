//clase 21/05

ACCION ejercicio1 ES
	AMBIENTE
		INSCRIPCIONES = Registro
			clave = Registro	
				facultad: AN(40);
				tema: an(30);
				curso: an(10);
				turno: an(10):
				fecha_in = Registro	
					año
					mes
					dia 
				Fin_reg;
				id
			Fin_reg;
			duracion: entero;
			cupo: N(9);
			cant_insc: N(9);
		Fin_reg;

		Arch: archivo de INSCRIPCIONES ordenado por clave;
		reg: INSCRIPCIONES;

		resg_curso:
		resg_tema:
		resg_facu:
		tot_gen, tot_facu, tot_tema, tot_curso: entero;

		PROCEDIMIENTO corte_curso() ES
			Escribir("La cantidad de inscriptos en el curso:",resg_curso,"es de:",tot_curso);
			tot_tema:= tot_tema + tot_curso;
			tot_curso:=0;
			resg_curso:= reg.clave.curso;
		Fin_proced;

		PROCEDIMIENTO corte_tema() ES	
			corte_curso();
			Escribir("La cantidad de inscriptos del tema:",resg_tema,"es de:",tot_tema);
			tot_facu:= tot_facu + tot_tema;
			tot_tema:= 0;
			resg_tema:= reg.clave.tema;
		Fin_proced;

		PROCEDIMIENTO corte_facu() ES
			corte_tema();
			Escribir("La cantidad de inscriptos de la facultad:",resg_facu,"es de:",tot_facu);
			tot_gen:= tot_gen + tot_facu;
			tot_facu;
			resg_facu:= reg.clave.facultad;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		resg_curso:= reg.clave.curso;
		resg_tema:= reg.clave.tema;
		resg_facu:= reg.clave.facu;

		tot_curso:=0;
		tot_tema:=0;
		tot_facu:=0;
		tot_gen:=0;

		MIENTRAS NFDA(Arch) HACER:
			SI (resg_facu <> reg.clave.facu) ENTONCES:
				corte_facu;
			SINO
				SI (resg_tema <> reg.clave.tema) ENTONCES:
					corte_tema;
				SINO	
					SI (resg_curso <> reg.clave.curso) ENTONCES:
						corte_curso;
					Finsi;
				Finsi;
			Finsi;

			tot_curso:= tot_curso + reg.cant_insc;

			LEER(Arch,reg);
		Fin_mientras;

		corte_facu();
		Escribir("El total general de inscriptos es de:",tot_gen);
		CERRAR(Arch);
	Fin_accion.


ACCION ejercicio2 ES
	AMBIENTE 
		INSCRIPCIONES = Registro
			clave = Registro	
				facultad: AN(40);
				tema: an(30);
				curso: an(10);
				turno: an(10):
				fecha_in = Registro	
					año
					mes
					dia 
				Fin_reg;
				id
			Fin_reg;
			duracion: entero;
			cupo: N(9);
			cant_insc: N(9);
		Fin_reg;

		Arch: archivo de INSCRIPCIONES ordenado por clave;
		reg: INSCRIPCIONES;

		resg_turno:
		resg_curso:
		resg_tema:
		resg_facu:
		tot_gen, tot_facu, tot_tema, tot_curso, tot_turno: entero;

		PROCEDIMIENTO corte_turno() ES	
			Escribir("La cantidad de inscriptos al turno:",resg_turno,"con una duración mayor a 40 horas, es de:",tot_turno);
			tot_curso:= tot_curso + tot_turno;
			tot_turno:= 0;
			resg_turno:= reg.clave.turno;
		Fin_proced;

		PROCEDIMIENTO corte_curso() ES
			corte_turno();
			Escribir("La cantidad de inscriptos al curso:",resg_curso,"con una duración mayor a 40 horas, es de:",tot_curso);
			tot_tema:= tot_tema + tot_curso;
			tot_curso:=0;
			resg_curso:= reg.clave.curso;
		Fin_proced;

		PROCEDIMIENTO corte_tema() ES
			corte_curso();
			
			SI (tot_tema < 500) ENTONCES:
				Escribir("En el tema:",resg_tema,"hay menos de 500 inscriptos");
				Escribir("El total actual es:",tot_tema);
			Finsi;
			
			tot_facu:= tot_facu + tot_tema;
			tot_tema:= 0;
			resg_tema:= reg.clave.tema;
		Fin_proced;

		PROCEDIMIENTO corte_facu() ES
			corte_tema();
			Escribir("La cantidad de inscriptos al curso:",resg_facu,"con una duración mayor a 40 horas, es de:",tot_facu);
			tot_gen:= tot_gen + tot_facu;
			tot_facu;
			resg_facu:= reg.clave.facultad;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);

		resg_turno:= reg.clave.turno;
		resg_curso:= reg.clave.curso;
		resg_tema:= reg.clave.tema;
		resg_facu:= reg.clave.facultad;

		tot_turno:=0;
		tot_curso:=0;
		tot_tema:=0;
		tot_facu:=0;
		tot_gen:=0;

		MIENTRAS NFDA(Arch) HACER:
			SI (resg_facu <> reg.clave.facultad) ENTONCES:
				corte_facu();
			SINO
				SI (resg_tema <> reg.clave.tema) ENTONCES:
					corte_tema();
				SINO
					SI (resg_curso <> reg.clave.curso) ENTONCES:
						corte_curso();
					SINO
						SI (resg_turno <> reg.clave.turno) ENTONCES:	
							corte_turno();
						Finsi;
					Finsi;
				Finsi;
			Finsi;

			SI (reg.duracion > 40) ENTONCES:
				tot_turno:= tot_turno + reg.cant_insc;
			Finsi;

			LEER(Arch,reg);
		Fin_mientras;

		corte_facu();
		Escribir("El total general de inscriptos es de:",tot_gen);
		CERRAR(Arch);
Fin_accion.


ACCION ejercicio3 ES
	AMBIENTE
		INSCRIPCIONES = Registro
			clave = Registro	
				facultad: AN(40);
				tema: AN(30);
				curso: AN(10);
				turno: AN(10):
				fecha_in = Registro	
					año
					mes
					dia 
				Fin_reg;
				id: N(6);
			Fin_reg;
			duracion: entero;
			cupo: N(9);
			cant_insc: N(9);
		Fin_reg;

		Arch: archivo de INSCRIPCIONES ordenado por clave;
		reg: INSCRIPCIONES;

		INSCRIPCIONES_SALIDA = Registro
			tema: AN(30);
			id: N(6);
			cant_insc: N(9);
		Fin_reg;

		Arch_sal: archivo de INSCRIPCIONES_SALIDA;
		reg_sal: INSCRIPCIONES_SALIDA;

		resg_turno: AN(10);
		resg_curso: AN(10);
		resg_tema: AN(30);
		tot_gral, tot_gralcompleto, tot_tema, tot_curso, tot_turno: entero;

		PROCEDIMIENTO corte_turno() ES
			Escribir("La cantidad de cursos que completaron el cupo del turno:",resg_turno,"es de:",tot_cursos_turno);
			tot_cursos_curso:= tot_cursos_curso + tot_cursos_turno; //aumento el total superior de la cantidad de cursos que ya superaron el cupo
			tot_insc_curso:= tot_insc_curso + tot_insc_turno; //aumento el total superior de la cantidad de inscriptos 
			tot_cursos_turno:=0;
			tot_insc_turno:=0;
			resg_turno:= reg.clave.turno;
		Fin_proced;

		PROCEDIMIENTO corte_curso() ES
			corte_turno();
			Escribir("La cantidad de cursos que completaron el cupo del curso:",resg_curso,"es de:",tot_cursos_curso);
			tot_cursos_tema:= tot_cursos_tema + tot_cursos_curso;
			tot_insc_tema:= tot_insc_tema + tot_insc_curso;
			tot_cursos_curso:=0;
			tot_insc_curso:=0;
			resg_curso:= reg.clave.curso;
		Fin_proced;

		PROCEDIMIENTO corte_tema() ES
			corte_curso();
			Escribir("La cantidad de cursos que completaron el cupo del tema:",resg_tema,"es de:",tot_cursos_tema);
			tot_cursos_facu:= tot_cursos_facu + tot_cursos_tema;
			tot_insc_facu:= tot_insc_facu + tot_insc_tema;

			reg_sal.tema:= resg_tema;
			reg_sal.id:= reg.clave.id;
			reg_sal.cant_insc:= tot_insc_tema;
			GRABAR(Arch_sal,reg_sal);

			tot_cursos_tema:=0;
			tot_insc_tema:=0;
			resg_tema:= reg.clave.tema;
		Fin_proced;

		PROCEDIMIENTO corte_facu() ES
			corte_tema();
			tot_cursos_gral:= tot_cursos_gral + tot_cursos_facu;
			tot_insc_gral:= tot_insc_gral + tot_insc_facu;
			tot_cursos_facu:=0;
			tot_insc_facu:=0;
			resg_facu:= reg.clave.facultad;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg):
		ABRIR S/(Arch_sal);

		resg_turno:= reg.clave.turno;
		resg_curso:= reg.clave.curso;
		resg_tema:= reg.clave.tema;

		tot_gral:=0;
		tot_gralcompleto:=0;
		tot_tema:=0;
		tot_curso:=0;
		tot_turno:=0;

		Escribir("|Tema|       |ID_curso|       |Cantidad de inscriptos|")

		MIENTRAS NFDA(Arch) HACER:
			SI (resg_facu <> reg.clave.facultad) ENTONCES:
				corte_facu();
			SINO
				SI (resg_tema <> reg.clave.tema) ENTONCES:
					corte_tema();
				SINO
					SI (resg_curso <> reg.clave.curso) ENTONCES:
						corte_curso();
					SINO
						SI (resg_turno <> reg.clave.turno) ENTONCES:
							corte_turno();
						Finsi;
					Finsi;
				Finsi;
			Finsi;

			SI (reg.cupo = reg.cant_insc) ENTONCES:
				tot_cursos_turno:= tot_cursos_turno + 1;  //cuanto la cantidad de cursos que completaron el cupo
			Finsi;

			tot_insc_turno:= tot_insc_turno + reg.cant_insc;

			LEER(Arch,reg);
		Fin_mientras;

		corte_facu();
		Escribir("El total general de inscriptos es:",tot_insc_gral);
		Escribir("El total general de inscriptos a cursos que ya completaron el cupo es:",tot_cursos_gral);

		CERRAR(Arch);
		CERRAR(Arch_sal);
Fin_accion.


ACCION ejercicio4 ES
	AMBIENTE
		INSCRIPCIONES = Registro
			clave = Registro	
				facultad: AN(40);
				tema: AN(30);
				curso: AN(10);
				turno: AN(10):
				fecha_in = Registro	
					año
					mes
					dia 
				Fin_reg;
				id: N(6);
			Fin_reg;
			duracion: entero;
			cupo: N(9);
			cant_insc: N(9);
		Fin_reg;

		Arch: archivo de INSCRIPCIONES ordenado por clave;
		reg: INSCRIPCIONES;

		INSCRIPCIONES_SAL = Registro
			facultad: AN(40);
			tema: AN(30);
			cant_cursos: entero;
		Fin_reg;

		Arch_sal: archivo de INSCRIPCIONES_SAL;
		reg_sal: INSCRIPCIONES_SAL;

		resg_curso: AN(10);
		resg_tema: AN(30);
		resg_facu: AN(40);

		resg_mayort: AN(30);
		tot_mayort: entero;

		tema_us: AN(30);

		PROCEDIMIENTO corte_curso() ES
			Escribir("La cantidad de alumnos inscriptos al cuso:",resg_curso,"es:",insc_curso_inicio);
			tot_tema:= tot_tema + insc_curso_inicio;
			tot_curso:= 0;
			resg_curso:= reg.clave.curso;
		Fin_proced;

		PROCEDIMIENTO corte_tema() ES
			corte_curso();
			
			SI (tot_tema > tot_mayort) ENTONCES:
				tot_mayort:= tot_mayort + tot_tema;
				resg_mayort:= resg_tema;
			Finsi;

			tot_facu:= tot_facu + tot_tema;
			tot_tema:=0;
			resg_tema:= reg.clave.tema;
		Finsi;

		PROCEDIMIENTO corte_facu() ES
			corte_tema();
			gral_curso_inicio:= gral_curso_inicio + tot_facu;
			tot_facu:=0;
			resg_facu:= reg.clave.facultad;
		Fin_proced;

	PROCESO
		ABRIR E/(Arch);
		LEER (Arch,reg);
		ABRIR S/(Arch_sal);

		cursos_inicio:=0;
		insc_curso_inicio:=0;
		cursos_noinicio:=0;
		insc_curso_no_inicio:=0;
		tot_curso_us:=0;

		resg_curso:= reg.clave.curso;
		resg_tema:= reg.clave.tema;
		resg_facu:= reg.clave.facultad;
		tot_mayort:= LV;

		Escribir("Ingrese unna tema del que quiere saber el porcentaje de cursos:");
		Leer(tema_us);

		Escribir("|Facultad|      |Tema|      |Cantidad de cursos|");

		MIENTRAS NFDA(Arch) HACER:
			SI (resg_facu <> reg.clave.facultad) ENTONCES:
				corte_facu();
			SINO
				SI (resg_tema <> reg.clave.tema) ENTONCES:
					corte_tema();
				SINO	
					SI (resg_curso <> reg.clave.curso) ENTONCES:
						corte_curso();
					Finsi;
				Finsi;
			Finsi;

			SI (fecha_sistema() >= reg.clave.fecha) ENTONCES: //controlo que el mes de inicio sea menor al de la facha del sistema (cursos que ya comenzaron)
				cursos_inicio:= cursos_inicio + 1;
				insc_curso_inicio:= insc_curso_inicio + reg.cant_insc;  //cuento los alumnos de los cursos que ya comenzaron, aumento el contador de menor nivel de corte
			SINO //cursos que aún no comenzaron 
				cursos_noinicio:= cursos_noinicio + 1
				insc_curso_no_inicio:= insc_curso_no_inicio + reg.cant_insc //cuento la cantidad de inscriptos a cursos que no comenzaron

				reg_sal.facultad:= reg.clave.facultad;
				reg_sal.tema:= reg.clave.tema;
				reg_sal.cant_cursos:= cursos_noinicio;
				GRABAR(Arch_sal,reg_sal);
			Finsi;

			SI (reg.clave.tema = tema_us) ENTONCES:
				tot_curso_us:= tot_curso_us + 1; //cuento llos temas que coinciden con los temasa ingresados por el usuario
			Finsi;
			LEER(Arch,reg);
		Fin_mientras;

		corte_facu();
		CERRAR(Arch);
		CERRAR(Arch_sal);

		Escribir("El total de alumnos inscriptos a cursos que ya comenzaron es:",gral_curso_inicio);
		Escribir("El total de alumnos inscriptos a cursos que todavían no han comenzaron es:",insc_curso_no_inicio);
		Escribir("El tema:",resg_mayort,"presenta la mayor cantidad de inscriptos y es:",tot_mayort);
		Escribir("El porcentaje de cursos del tema ingresado por el usuario es:",(tot_curso_us DIV (cursos_inicio + cursos_noinicio)),"%");
Fin_accion.


ACCION ejercicio5 ES	
	AMBIENTE
		PROYECTOS = Registro
			clave = Registro
				suc = {"R", "F", "C", "P"};
				area:
				id:
			Fin_reg;
			cant_trab: N(8);
			estado = {"N","I","P","F"} //N no iniciado, I en curso, P parado, F finalizado
			presup_act: N(8,2);
			presup_inic: N(8,2);
		Fin_reg;
		Arch: archivo de PROYECTOS ordenado por clave;
		reg: PROYECTOS;

		PROYECTOS_SAL = Registro
			suc:
			area:
			cant_proy:
			cant_trab:
		Fin_reg;
		Arch_sal: archivo de PROYECTOS_SAL;
		reg_sal: PROYECTOS_SAL;

		PROCEDIMIENTO corte_area() ES
			Escribir("En el área",resg_area,"la cantidad de proyectos que no finalizaron son:",tot_area_proy,"

	PROCESO
		ABRIR E/(Arch);
		LEER(Arch,reg);
		ABRIR S/(Arch_sal);

		Escribir("|Sucursal|	|Área|		|Cantidad de proyectos|		|Total de trabajadores|");
		
		MIENTRAS NFDA(Arch) HACER:
			SI (resg_suc <> reg.clave.suc) ENTONCES:
				corte_suc();
			SINO
				SI (resg_area <> reg.clave.area) ENTONCES:
					corte_area();			
				Finsi;
			Finsi;

			
