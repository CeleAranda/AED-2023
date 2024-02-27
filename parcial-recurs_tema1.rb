parcial-recurs_tema1

ACCION ejercicio1_pt2 ES
    AMBIENTE
        sec_art, sec_vent; sec_sal: secuencia de caracter;
        v_art, v_vent: caracter;
        resg_m1, resg_m2, m1, m2, resg_FE: caracter;
        UV, stock, tot_d, tot_s: entero

        FUNCION a_entero(v: caracter): entero ES
            SEGUN v HACER   
                "0": a_entero:= 0;
                "1": a_entero:= 1;
                "2": a_entero:= 2;
                "3": a_entero:= 3;
                "4": a_entero:= 4;
                "5": a_entero:= 5;
                "6": a_entero:= 6;
                "7": a_entero:= 7;
                "8": a_entero:= 8;
                "9": a_entero:= 9;
            Fin_segun;
        Fin_funcion;

    PROCESO
        ARR(sec_art);
        AVZ(sec_art, v_art);
        ARR(sec_vent);
        AVZ(sec_vent, v_vent);
        CREAR(sec_sal);

        # totales etc
       
        Escribir("Ingrese un mes, 2 dígitos por separado, para generar un informe de ventas: ");
        Leer(m1, m2);

        Escribir("Informe de ventas");
        Escribir("|Nombre del artículo|        |Cant. unid entregadas en suc.|        |Cant. unida enviadas a domicilio|");

        MIENTRAS NFDS(sec_art) HACER
            MIENTRAS (v_vent <> "#") HACER   # trato una venta
                PARA i:= 1 a 2 HACER
                    AVZ(sec_vent, v_vent);
                Fin_para;

                resg_m1:= v_vent;
                AVZ(sec_vent, v_vent);
                resg_m2:= v_vent;
                AVZ(sec_vent, v_vent);

                AVZ(sec_vent, v_vent);     # avanzo la FP
                resg_FE:= v_vent;
                AVZ(sec_vent, v_vent);

                PARA i:= 1 a 0 HACER
                    UV:= UV + (a_entero(v_vent) * 10**i);   #1 -> 10, 5 -> 5 --> 15
                    AVZ(sec_vent, v_vent);
                Fin_para;
                
                SI (resg_FE = "S") ENTONCES
                    tot_s:= tot_s + UV;     #unidades vendidas con entrega en sucursal
                SINO 
                    tot_d:= tot_d + UV;    
                Fin_si;

            MIENTRAS (v_art <> "&") HACER
                PARA i:= 1 a 6 HACER   
                        AVZ(sec_art,v_art);
                Fin_para;

                PARA i:= 2 a 0 HACER
                    stock:= stock + (a_entero(v_art) * 10**i);
                    AVZ(sec_art,v_art);
                Fin_para;
                











                # SI (m1 = resg_m1 Y m2 = resg_m2) ENTONCES
                #     Escribir(v_art);
                #     AVZ(sec_art,v_art);
                #     Escribir(tot_s);
                #     Escribir(tot_d);
                # SINO
                #     SI (stock - tot_s = 0) ENTONCES
                #         ESC(sec_sal, v_art);
                #         AVZ(sec_art,v_art);
                #     Fin_si;
                # Fin_si
                AVZ(sec_art,v_art);
            Fin_mientras;
        Fin_mientras;
        




        
        
        