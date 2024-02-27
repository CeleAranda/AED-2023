# Modelo de parcial

# secuencia GEMAS: NombreDeGema,CompQuimica,Descripcion,%NombreDeGema,CompQuimica,Descripcion,%etc..FDS

# secuancia CODIGOS: 0 1 2 3 4 5 6 7 8 9 10 11 12... 999

ACCION ejercicio_parcial ES
    AMBIENTE
        sec_gem, sec_sal: secuencia de caracter;
        v_gem: caracter
        sec_cod: secuencia de entero;
        v_cod, resg_codImp, resg_codPar, cont_gem: entero;

    PROCESO
        ARR(sec_gem);
        AVZ(sec_gem, v_gem);
        ARR(sec_cod);
        AVZ(sec_cod, v_cod);
        CREAR(sec_sal);


        Escribir("Nombre comercial de las gemas:");

        MIENTRAS NFDS(sec_gem) HACER
            MIENTRAS (v_gem <> "%") HACER
                SI (v_cod MOD 2 <> 0) ENTONCES    
                    MIENTRAS (v_gem <> ",") HACER
                        Escribir(v_gem);
                        AVZ(sec_gem, v_gem);
                    Fin_mientras;
                    resg_codImp:= v_cod;
                SINO 
                    resg_codPar:= v_cod;
                Fin_si;

                AVZ(sec_gem, v_gem);   # avanzo ","

                MIENTRAS (v <> ",") HACER   
                    ESC(sec_sal, v_gem);
                    AVZ(sec_gem, v_gem);
                Fin_mientras;
                ESC(sec_sal, v_gem)    # copio ","

            cont_gem:= cont_gem + 1;   
            Fin_mientras;



