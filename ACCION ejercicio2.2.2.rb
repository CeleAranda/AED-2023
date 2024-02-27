ACCION ejercicio2.2.2.17 ES 
    AMBIENTE
        aspirantes_ag = Registro 
            DNI: N(8);
            ApelyNomb: AN(30);
            Carrera: {"ISI", "IQ", "IEM", "LAR"};
            F_Nac: f_fecha;
            Email: AN(30);
            ColegioSec: AN(30);
            Fecha_insc = Registro
                año: N(4);
                mes: 8;
                dia: 1..31;
            Fin_registro;
            Aprobado: {"Sí", "No"};
        Fin_registro;

        aspirantes_feb = Registro 
            DNI: N(8);
            ApelyNomb: AN(30);
            Carrera: {"ISI", "IQ", "IEM", "LAR"};
            F_Nac: f_fecha;
            Email: AN(30);
            ColegioSec: AN(30);
            Fecha_insc = Registro
                año: N(4);
                mes: 2;
                dia: 1..31;
            Fin_registro;
            Aprobado: {"Sí", "No"};
        Fin_registro;

        Arch1: archivo de aspirantes_ag ordenado por DNI;
        reg1: aspirantes_ag;
        Arch2: archivo de aspirantes_feb ordenado por DNI;
        reg2: aspirantes_feb;
        
        seguimiento = Registro
            DNI: N(8);
            ApelyNomb: AN(30);
            Email: AN(30);
            ColegioSec: AN(30);
        Fin_registro;

        Arch_sal: archivo de seguimiento ordenado por DNI;
        reg_sal: seguimiento;

        cont_asp: entero;


        PROCESO
            ABRIR E/(Arch1);
            LEER(Arch1, reg1):
            ABRIR E/(Arch1);
            LEER(Arch2, reg2)
            ABRIR S/(Arch_sal);

            cont_asp:= 0;
            
            MIENTRAS NFDA(Arch1) Y NFDA(Arch2) HACER
                