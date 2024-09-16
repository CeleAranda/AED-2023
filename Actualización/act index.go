//archivos indexados 

ACCION ejercicio231a ES
	AMBIENTE
		Facturas = Registro
			nro_cli:
			nro_fac:
			fecha:
			importe:
		Fin_reg;

		arch_fac: archivo de Facturas ordenado por nro_cli y nro_fac;
		fac: Facturas;

		Clientes = Registro
			nro_cli:
			nomb:
			dni:
			cuit:
			dom:
		Fin_reg;

		arch_cli: archivo de Clientes iindexado por nro_cli;
		cli: Clientes;

	PROCESO
		ABRIR E/(arch_fac);
		LEER(arch_fac,fac);
		ABRIR E/(arch_cli);
		LEER(arch_cli,cli);

		Esc("|Nro Cliente| |Nombre| |Nro Factura|")
		
		MIENTRAS NFDA(arch_fac) HACER:
			cli.nro_cli:= fac.nro_cli; //le asigno al campo del registro del archivo indexado lo que quiero buscar
			Leer(arch_cli,cli);

			SI EXISTE ENTONCES:
				Esc(cli.nro_cli, cli.nomb, cli.nro_fac);
			SINO
				Esc("El cliente no existe en el archivo");
			Fin_si;

			Leer(arch_fac,fac);
		Fin_mientras;






//actualizar un registro
	MIENTRAS NFDA(arch_fac) HACER:
			cli.nro_cli:= fac.nro_cli; //le asigno al campo del registro del archivo indexado lo que quiero buscar
			Leer(arch_cli,cli);

			SI EXISTE ENTONCES:
				//ejemplo de actualizacion de un campo
				Esc("ingrese un importe"); Leer(importe_nuev);
				cli.importe:= importe_nuev;
				REGRABAR(arch_cli,cli);
			SINO
				Esc("El cliente no existe en el archivo");
			Fin_si;

			Leer(arch_fac,fac);
		Fin_mientras;


ACCION ejercicio232 es
	AMBIENTE
		empleados = registro
			nro_suc:
			categ: {"A","B","C"};
			nomb_emp:
			cod_curso:
			tecnico: {"si","no"};
		Fin_reg;

		arch_emp: archivo de empleados ordenado por nro_suc y categ;
		e:empleados;

		curso = registro
			cod_curso;
			nomb:
			fecha:
			cant_horas:
		Fin_reg;

		arch_cur: archivo de curso indexado por cod_curso;
		c: curso;

	PROCESO
		ABRIR E/(arch_emp);
		Leer(arch_emp,e);
		ABRIR E/(arch_cur);
		Leer(arch_cur,c);

		Esc("|Sucursal| |Categoría| |Nombre del empleado| |Nombre del curso|");

		MIENTRAS NFDA(arch_emp) HACER:
			c.cod_curso:= e.cod_curso;
			Leer(arch_emp,e);

			SI EXISTE ENTONCES:
				Esc(e.nro_suc,e.categ,e.nomb_emp,c.nomb);
			SINO
				Esc("El empleado",e.nomb_emp,"no existe en el archivo.");
			Fin_si;




ACCION ejercicio2310 ES
	AMBIENTE
		puntos = registro
			dni:
			cant:
			ult_carga:
		Fin_reg;

		arch_punt: archivo de puntos indexado por dni;
		p:puntos;

		cliente = registro
			dni:
			apeynom:
			edad:
			ciudad:
		Fin_reg;

		arch_cli: archivo de cliente indexado por dni;
		c: cliente;

	PROCESO
		ABRIR E/S(arch_punt);
		Leer(arch_punt,p);
		ABRIR E/(arch_cli);
		Leer(arch_cli,c);

		Esc("Desea ingresar datos? si/no");
		Leer(rta);

		MIENTRAS (rta = "si") HACER:
			Esc("Ingrese su DNI");
			Leer(dni);
			Esc("Ingrese el monto que desea cargar");
			Leer(monto);

			c.dni:= dni; //en archivo de cliente asigno lo que ingresó el ususario para verificar que exista
			SI EXISTE ENTONCES:
				p.dni:= dni;
				SI EXISTE ENTONCES: //existe en clientes y en puntos

					SI (p.cant >= 100) ENTONCES:
						monto:= monto - 100;
					Fin_si;

					p.cant:= (monto div 100); 
					p.ult_carga:= fecha_actual();
					REGRABAR(arch_punt,p);
				
				SINO //alta de cliente en ambos archivos 
					p.dni:= dni;
					p.cant:= (monto div 100); 
					p.ult_carga:= fecha_actual();
					GRABAR(arch_punt,p);
					
					c.dni:= dni;
					Esc("Ingrese su nombre y apellido:");
					Leer(c.apeynom);
					Esc("Ingrese su edad");
					Leer(c.edad);
					Esc("Ingrese su ciudad");
					Leer(c.ciudad);
					GRABAR(arch_cli,c);
				Fin_si;
			SINO
				cont_cli:= cont_cli + 1; //cuento los clientes que no existen
			Fin_si;

			Esc("Desea continuar? si/no");
			Leer(rta);
		Fin_mientras;
		CERRAR(arch_punt);
		CERRAR(arch_cli);

		Esc("La cantidad de clientes que no existen es de:",cont_cli);
	Fin_accion.


	