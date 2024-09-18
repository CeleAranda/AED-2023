ACCION ejercicio 2219 ES //act unitaria
    AMBIENTE
        fecha = Registro
            año:
            mes:
            dia:
        Fin_reg;

        Clave = Registro
            farm:
            medic:
        Fin_reg
        
        MAE_REMEDIOS = Registro
            clave_mae: Clave
            cant_actual:
            fecha_ven: fecha;
        Fin_reg;

        Mae: archivo de MAE_REMEDIOS ordenado por clave_mae;
        reg_m: MAE_REMEDIOS;

        MOVIMIENTOS = Registro  
            clavem = Registro
                clave: Clave;
                cod_mov: {1,2,3};
            Fin_reg;
            cant_recibida:
        Fin_reg;

        Mov: archivo de MOVIMIENTOS ordenado por clavem;
        reg_mov: MOVIMIENTOS;

        REM_VENC = Registro 
            medic:
            cant_venc:
        Fin_reg;

        Venc: archivo de REM_VENC;
        reg_v: REM_VENC;

        MAE_REMEDIOS_SAL = Registro
            clave_mae: Clave
            cant_actual:
            fecha_ven: fecha;
        Fin_reg;   

        Mae_sal: archivo de MAE_REMEDIOS_SAL;
        reg_m_sal: MAE_REMEDIOS_SAL;


        PROCEDIMIENTO LeerMae() ES
            LEER(Mae, reg_m);

            SI FDA(Mae) ENTONCES:
                reg_m.clave_mae:= HV;
            Fin_si;
        Fin_proced;

        PROCEDIMIENTO LeerMov() ES
            LEER(Mov, reg_mov);

            SI FDA(Mov) ENTONCES:
                reg_mov.clavem:= HV;
            Fin_si;
        Fin_proced;

    PROCESO
        ABRIR E/(Mae);
        ABRIR E/(Mov);
        ABRIR /S(Venc);
        ABRIR /S(Mae_sal);

        LeerMae();
        LeerMov();

        MIENTRAS (reg_m.clave_mae <> HV) O (reg_mov.clavem <> HV) HACER:
            SI (reg_m.clave_mae < reg_mov.clavem) ENTONCES: //si la clave del maestro es menor, solo se lo copia en salida
                reg_m_sal:= reg_m;
                GRABAR(Mae_sal,reg_m_sal);
                LeerMae(); //se lee el archivo menor (maestro)
            SINO
                SI (reg_m.clave_mae = reg_mov.clavem) ENTONCES:
                    SI (reg_mov.cod_mov = 1) ENTONCES:
                        Esc("¡ERROR!"); //error en alta
                    SINO    
                        SI (reg_mov.cod_mov = 2) ENTONCES: //baja
                            reg_v.medic:= reg_m.clave_mae.Clave.medic;
                            reg_v.cant_venc:= reg_m.cant_actual;
                            GRABAR(Venc,reg_v); //si el código de mov es 2, grabo en el archivo de medicamentos vencidos
                        SINO 
                            //modificacion 
                            //asigno campo a campo
                            reg_m_sal.cant_actual:= reg_m_sal.cant_actual + reg_mov.cant_recibida;
                            GRABAR(Mae_sal, reg_m_sal);
                        Fin_si;
                    Fin_si;
                    LeerMae();
                    LeerMov(); //se leen ambos archivos, porque la clave es igual
                SINO
                    //la clave del maestro es mayor
                    SI (reg_mov.cod_mov = 1) ENTONCES: //alta
                        //asigno campo a campo
                        reg_m_sal.cant_actual:= reg_m_sal.cant_actual + reg_mov.cant_recibida;
                        //no entiendo si a la fecha le tengo que sumar 30 o queE
                        LeerMov();
                    SINO    
                        SI (reg_mov.cod_mov = 2) ENTONCES:
                            Esc("¡ERROR!"); //error en baja
                        SINO 
                            Esc("¡ERROR!"); //error en modificacion 
                        Fin_si;
                    Fin_si;
                Fin_si;
            Fin_si;
        Fin_mientras;
        CERRAR(Mae);
        CERRAR(Mov);
        CERRAR(Mae_sal);
Fin_accion.




ACCION ejercicio 2220 ES //act por lotes
	AMBIENTE
		fecha = Registro
			año:
			mes:
			dia:
		Fin_reg;

		Almunos = Registro
			leg:
			apelynom:
			car: {"ISI", "IQ", "IEM"};
			fecha_nac: fecha;
			fecha_ing: fecha;
			dni:
			sexo: {"M","F"};
			fecha_ult_ex: fecha;
			nota:
		Fin_reg;

		ArchAlum: archivo de alum+ ordenado por leg;
		reg_alum, aux: Almunos;

		Examenes = Registro //movimientos
			leg:
			car: {"ISI", "IQ", "IEM"};
			cod_mat:
			fecha_ex: fecha;
			nota:
		Fin_reg;

		ArchEx: archivo de Examenes ordenado por leg;
		reg_ex: Examenes;

		Almunos_sal = Registro
			leg:
			apelynom:
			car: {"ISI", "IQ", "IEM"};
			fecha_nac: fecha;
			fecha_ing: fecha;
			dni:
			sexo: {"M","F"};
			fecha_ult_ex: fecha;
			nota:
		Fin_reg;

		ArchSal: archivo de alum+ ordenado por leg;
		reg_sal: Almunos;

		PROCEDIMIENTO LeerMae() ES
			LEER(ArchAlum,reg_alum);

			SI FDA(ArchAlum) ENTONCES:
				reg_alum.leg:= HV;
			Fin_si;
		Fin_proced;

		PROCEDIMIENTO LeerMov() ES
			LEER(ArchEx,reg_ex);

			SI FDA(ArchEx) ENTONCES:
				reg_ex.leg:= HV;
			Fin_si;
		Fin_proced;

	PROCESO	
		ABRIR E/(ArchAlum);
		ABRIR E/(ArchEx);
		ABRIR /S(ArchSal);

		LeerMae();
		LeerMo();

		MIENTRAS (reg_alum.leg <> HV) O (reg_ex.leg <> HV) HACER:
            SI (reg_alum.leg < reg_ex.leg) ENTONCES: //no hay movimientos
                reg_sal:= reg_alum;
                GRABAR(ArchSal,reg_sal);
                LeerMae();
            SINO
                SI (reg_alum.leg = reg_ex.leg) ENTONCES: //baja o modificacion 
                    aux:= reg_alum; //resguardo el registro maestro
                    MIENTRAS (reg_alum.leg = reg_ex.leg) HACER: //mmietras la clave del maestro y movimiento sean iguales
                        SI (reg_ex.nota > 5) HACER: 
                            aux.nota:= reg_ex.nota; //modifico el auxiliar porque no puedo modificar sobre el maestro
                            aux.fecha_ult_ex:= reg_ex.fecha_ex;
                        Fin_si;
                        LeerMov(); //leo el de movimientos para corroboar que las claves sean iguales
                    Fin_mientras;
                    reg_sal:= aux;
                    GRABAR(ArchSal,reg_sal);
                    LeerMae();
                SINO


Fin_accion.




ACCION ejercicio 2221 ES //act unitaria
    AMBIENTE
        fecha = Registro
            año: N(4);
            mes: 1..12;
            dia: 1..31;
        Fin_reg;

        Clave = Registro
            cod_us: N(8);
            cod_amig: N(8);
        Fin_reg;

        Amigos = Registro
            clavea: Clave;
            fecha_amist: fecha;
            mensaj: AN(50);
        Fin_reg;

        ArchAmig: archivo de Amigos ordeando por clavea;
        ami: Amigos;

        Amigos_sal = Registro
            clave_s: Clave;
            fecha_amist: fecha;
            mensaj: AN(50);
        Fin_reg;

        ArchSal: archivo de Amigos_sal;
        reg_sal: Amigos;

        Notificaciones = Registro
            claven: Clave;
            cod_mov: {"A","B","M"};
            fecha_amist: fecha;
            mensaj: AN(50);
        Fin_reg;

        ArchNot: archivo de Notificaciones ordenado por claven y cod_mov;
        not: Notificaciones;

        PROCEDIMIENTO LeerA() ES
            Leer(ArchAmig,ami);

            SI FDA(ArchAmig) ENTONCES:
                ami.clavea:= HV;
            Fin_si;
        Fin_proced;

        PROCEDIMIENTO LeerM() ES
            Leer(ArchNot,not);

            SI FDA(ArchNot) ENTONCES:
                ami.claven:= HV;
            Fin_si;
        Fin_proced;

    PROCESO 
        ABRIR E/(ArchAmig);
        ABRIR E/(ArchNot);
        ABRIR /S(ArchSal);

        LeerA();
        LeerM();

        MIENTRAS (ami.clavea <> HV) o (not.claven <> HV) HACER:
            SI (ami.clavea = not.claven) ENTONCES: //claves iguales -> hay movimientos para el Maestro (alta=error)
                SI (not.cod_mov = "A") ENTONCES:
                    Esc("¡Error! Ya eres amigo de este usuario.");
                SINO    
                    SI (not.cod_mov = "B") ENTONCES: //baja lógica
                        Esc("Se ha eliminado a:",ami.cod_amig); //en act sec no se puede eliminar de manera física, solo se avanza
                    SINO
                        //igual a "M"              
                        //se asigna campo a campo y se graba en salida           
                        reg_sal.clave_s:= not.claven;
                        reg_sal.fecha_amist:= not.fecha_amist;
                        reg_sal.mensaj:= not.mensaj; //modific
                        GRABAR(ArchSal,reg_sal);
                    Fin_si;
                Fin_si;
                LeerA();
                LeerM();
            SINO    
                SI (ami.clavea < not.claven) ENTONCES: //la clave del maestro es menor -> se copia el maestro en salida y se lo lee (es válido la modificación o la baja lógica)
                    reg_sal:= ami;
                    GRABAR(ArchSal,reg_sal);
                    LeerA();
                SINO
                    //la clave del maestro es mayor -> solo es válido el alta
                    SI (not.cod_mov = "A") ENTONCES:
                        //se asigna campo a campo (reg_sal con movimientos)
                        reg_sal.clave_s:= not.claven;
                        reg_sal.fecha_amist:= not.fecha_amist;
                        reg_sal.mensaj:= not.mensaj; 
                        GRABAR(ArchSal,reg_sal);
                    SINO    
                        SI (not.cod_mov = "B") ENTONCES: //baja lógica
                            Esc("¡Error! No existe este ususario.");
                        SINO
                            //igual a "M"                        
                            Esc("¡Error! No existe este usuario.");
                        Fin_si;
                    Fin_si;
                    LeerM();
                Fin_si;
            Fin_si;
        Fin_mientras;

        CERRAR(ArchAmig);
        CERRAR(ArchNot);
        CERRAR(ArchSal);
Fin_accion.
                
                
                
                
ACCION ejercicio 2222 ES //act unitaria
    AMBIENTE
        fecha = Registro
            año: N(4);
            mes: 1..12;
            dia: 1..31;
        Fin_reg;

        Clave = Registro
            cod_us: N(8);
            cod_amig: N(8);
        Fin_reg;

        Amigos = Registro
            clavea: Clave;
            fecha_amist: fecha;
            mensaj: AN(50);
        Fin_reg;

        ArchAmig: archivo de Amigos ordeando por clavea;
        ami: Amigos;

        Amigos_sal = Registro
            clave_s: Clave;
            fecha_amist: fecha;
            mensaj: AN(50);
        Fin_reg;

        ArchSal: archivo de Amigos_sal;
        reg_sal: Amigos;

        Notificaciones = Registro
            claven: Clave;
            cod_mov: {"A","B","M"};
            fecha_amist: fecha;
            mensaj: AN(50);
        Fin_reg;

        ArchNot: archivo de Notificaciones ordenado por claven y cod_mov;
        not: Notificaciones;

        PROCEDIMIENTO LeerA() ES
            Leer(ArchAmig,ami);

            SI FDA(ArchAmig) ENTONCES:
                ami.clavea:= HV;
            Fin_si;
        Fin_proced;

        PROCEDIMIENTO LeerM() ES
            Leer(ArchNot,not);

            SI FDA(ArchNot) ENTONCES:
                ami.claven:= HV;
            Fin_si;
        Fin_proced;

        PROCEDIMIENTO mayor() ES
            SI (ami.clavea.Clave.cod_us = resg) ENTONCES:
                cont_amig:= cont_amig + 1;
            SINO
                SI (ami.clavea.Clave.cod_us <> HV) ENTONCES:
                    SI (cont_amig > mayor_cont) ENTONCES:
                        mayor_cont:= cont_amig;
                        mayor_us:= resg;
                    Fin_si;
                    resg:= ami.clavea.Clave.cod_us;
                    cont_amig:=0;
                Fin_si;
            Fin_si;
        Fin_proced;


    PROCESO 
        ABRIR E/(ArchAmig);
        ABRIR E/(ArchNot);
        ABRIR /S(ArchSal);

        LeerA();
        LeerM();

        resg:= ami.clavea.Clave.cod_us; //resguardo el ususario 

        MIENTRAS (ami.clavea <> HV) o (not.claven <> HV) HACER:
            mayor();
            SI (ami.clavea = not.claven) ENTONCES: //claves iguales -> hay movimientos para el Maestro (alta=error)
                SI (not.cod_mov = "A") ENTONCES:
                    Esc("¡Error! Ya eres amigo de este usuario.");
                SINO    
                    SI (not.cod_mov = "B") ENTONCES: //baja lógica
                        Esc("Se ha eliminado a:",ami.cod_amig); //en actualiz sec no se puede eliminar de manera física, solo se avanza
                    SINO
                        //igual a "M"              
                        //se asigna campo a campo y se graba en salida           
                        reg_sal.clave_s:= not.claven;
                        reg_sal.fecha_amist:= not.fecha_amist;
                        reg_sals.mensaj:= not.mensaj; //modific
                        GRABAR(ArchSal,reg_sal);
                    Fin_si;
                Fin_si;
                LeerA();
                LeerM();
            SINO    
                SI (ami.clavea < not.claven) ENTONCES: //la clave del maestro es menor -> se copia el maestro en salida y se lo lee (es válido la modificación o la baja lógica)
                    reg_sal:= ami;
                    GRABAR(ArchSal,reg_sal);
                    LeerA();
                SINO
                    //la clave del maestro es mayor -> solo es válido el alta
                    SI (not.cod_mov = "A") ENTONCES:
                        //se asigna campo a campo (reg_sal con movimientos)
                        reg_sal.clave_s:= not.claven;
                        reg_sal.fecha_amist:= not.fecha_amist;
                        reg_sal.mensaj:= not.mensaj; 
                        GRABAR(ArchSal,reg_sal);
                    SINO    
                        SI (not.cod_mov = "B") ENTONCES: //baja lógica
                            Esc("¡Error! No existe este ususario.");
                        SINO
                            //igual a "M"                        
                            Esc("¡Error! No existe este usuario.");
                        Fin_si;
                    Fin_si;
                    LeerM();
                Fin_si;
            Fin_si;
        Fin_mientras;

        CERRAR(ArchAmig);
        CERRAR(ArchNot);
        CERRAR(ArchSal);
Fin_accion.                
               


ACCION ejercicio 2224 ES
    AMBIENTE
        costo_afiliado = registro
            cod_af:
            fecha_a:
            fecha_b:
            cant_serv:
            costo:
        Fin_reg;
        arch_cost, arch_sal: archivo de costo_afiliado ordeando por cod_af;
        c,sal,aux:costo_afiliado;

        servicios = registro
            cod_af:
            fecha_act:
            cod_serv:
            costo:
        Fin_reg;
        arch_serv: archivo de servicios ordeando por cod_af;
        s: servicios

        PROCEDIMIENTO leerA()
            Leer(arch_cost,c);
            
            SI FDA(arch_cost) ENTONCES:
                c.cod_af:= HV;
            Fin_si;
        Fin_proced;

        PROCEDIMIENTO leerB()
            Leer(arch_serv,s);

            SI FDA(arch_serv) ENTONCES:
                s.cod_af:= HV;
            Fin_si;
        Fin_proced;


    PROCESO
        //abrir archivos


        MIENTRAS (c.cod_af <> HV) o (s.cod_af <> HV) HACER:
            SI (c.cod_af = s.cod_af) ENTONCES: //claves iguales -> baja y modificacion 
                MIENTRAS (c.cod_af = s.cod_af) HACER:
                    SI (s.cod_serv = 001) ENTONCES: //alta, es un cliente nuevo -> error
                       Esc("¡Error! Este cliente no está dado de alta")
                    SINO
                        SI (s.cod_serv = 000) ENTONCES: //baja, actualizo la fehca de baja con la fecha actual del archivo de movimientos 
                            sal.fecha_b:= s.fecha_act;
                        SINO //si el cod de servicio es cualquuier otro valor, cuento los servicios y el costo
                            sal.cant_serv:= sal.cant_serv + 1;
                            sal.costo:= s.costo;

