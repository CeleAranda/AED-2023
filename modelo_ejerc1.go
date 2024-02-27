//Ejercicio1 parcial xd

ACCION ejercicio1 ES
    AMBIENTE
        COMPRAS = registro
            fecha_comp:
            dni_cliente:
            cant_art:
            import_total:
        Fin_registro;

        SOCIOS = registro
            dni:
            apelynom:
            fecha_ad:
            categoria:
        Fin_registro;

        COM: archivo de COMPRAS ordenado por fecha_comp;
        c: COMPRAS
        SOC: archivo SOCIOS indexado por dni;
        s: SOCIOS;

        NODO = registro  //lista doble (se crea)
            apelynom: AN(30)
            chances: N(4)
            prox: puntero a NODO;
            ant: puntero a NODO;
        Fin_registro;

        p, q, prim, ult: puntero a NODO; //q para crear punteros nuevos y p es para recorrer

        CIRCULAR = registro  //solo se recorre, de igual manera se define
            num: entero;
            prox: puntero a CIRCULAR;
        Fin_registro;

        pc, qc, prim_c, ant_c: puntero a CIRCULAR 

        PROCEDIMIENTO insertar_doblenlazada ES
            NUEVO(q);  //insertar por primera vez
                *q.apelynom:= s.apelynom;
                chances_tot:= chances_tot+5;

            SI prim = NILL ENTONCES 
                prim:=q;
                ult:=q;
                *q.prox:= NILL
                *q.ant:= NILL
            SINO
                p:=prim;
                MIENTRAS p <> NILL Y *q.apelynom < *p.apelynom HACER //carga ordenada ascendente
                    p:= *p.prox; //prox es un campo!!!!! por eso el *!!!
                Fin_mientras;

                SI p=prim ENTONCES  //insertar en el primer lugar
                    prim:= q;
                    *q.prox:= p;
                    *q.ant:= NILL;
                    *p.ant:=q;
                SINO
                    SI p = NILL ENTONCES  //insertar en ultimo lugar
                        *p.prox:=q;
                        *q.ant:=p;
                        *q.prox:= NILL;
                        ult:=q; //ult apunta al ultimo elemento
                    SINO
                        *q.prox:= p;
                        *q.ant:=*p.ant; //apuntan al mismo lugar
                        *p.ant:=q;
                        *(*q.ant).prox := q; //voy al anterior de q (no se sabe cual es, pero si la direc de memoria) y modifico el campo proximo, asignando q
                    Fin_si;
                Fin_si;
            Fin_si;
        Fin_procedimiento;

    PROCESO
        ABRIR E/(COMPRAS) //el archivo secuencial controla el mientras
        LEER(COMPRAS,c)
        ABRIR E/(SOCIOS)  //no se modifica
        LEER(SOCIOS,s)

        prim:= NILL //creo lista doble enlazada
        ult:= NILL
        k:=1; //k es la posicion donde queda parada la funcion circular  

        MIENTRAS NDFA(COMPRAS) HACER
            cliente_soc:=c.dni_cliente;  //s.dni:=c.dni_cliente
            s.dni:=cliente_soc;
            LEER(SOCIOS,s)  //lee para buscar

            SI EXISTE ENTONCES
                MIENTRAS p <> NILL Y *p.apelynom <> s.apelynom HACER  //controla que el cliente no este repetido en la lista
                    p:= *p.prox;
                Fin_mientras;
                SI p=NILL ENTONCES  //se llegó al final de la lista, entonces, el cliente no está en la lista, pero hay que agregarlo
                    insertar_doblenlazada();
                SINO
                    operacion:= c.import_total DIV 100 //el puntero ya existe, solo se modifican las chances!!
                    *p.chances:= *p.chances + operacion //se suman un punto (chance) extra
                Fin_si;
                
                SI s.categoria = "black" ENTONCES
                    tiro:=Tirar();
                    PARA i:=k a tiro HACER  
                        ant_c:=pc;
                        pc:=*pc.prox;  //se recorre la lista tantas veces (tiro veces)
                    Fin_para;
                    k:=tiro + k;

                    chances_tot:= chances_tot + *pc.num; //pc es el dato que indica cuantas chances hay que sumar y esta en la posicion que dice el para
                    *q.chances:= chances_tot;
                Fin_si;
            SINO
                Escribir("Error! este cliente no es socio")
            Fin_si;
        Fin_mientras;
    Fin_accion.


