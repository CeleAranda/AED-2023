ACCION modelo_supermercado ES
	AMBIENTE
		stock = registro
			id_prod: AN(30);
			stock: N(9);
		Fin_reg;
		arch_stock, arch_stock_sal: archivo de stock ordenado por id_prod;
		s, s_sal: stock;

		movimientos_diarios = registro	
			id_prod: AN(30);
			id_cli: AN(30);
			tipo: {"C","D"}; //c: compra, d: devolución
			cant: N(9),
			precio_unit: N(9,2);
			total: N(9,2);
			tipo_env: {"domicilio", "retiro"}
		Fin_reg;
		arch_mov: archivo de movimientos_diarios ordenado por id_prod;
		m: movimientos_diarios;

		productos = registro
			id_prod: AN(30);
			nomb: AN(30);
			desc: AN(30);
			rubro: {"Limpieza","Carnicería","Verdulería","Bazar","Panadería"};
		Fin_reg;
		arch_prod: archivo de productos indexado por id_prod;
		reg: productos;

		PROCEDIMIENTO leerA()
			Leer(arch_stock,s);

			SI FDA(arch_stock) ENTONCES:
				s.id_prod:= HV;
			Fin_reg;
		Fin_proced;

		PROCEDIMIENTO leerB()
			Leer(arch_mov,m);

			SI FDA(arch_mov) ENTONCES:
				m.id_prod:= HV;
			Fin_reg;
		Fin_proced;

		cant_prod, cant_prod_retiro: entero;

	PROCESO
		ABRIR E/(arch_stock);
		ABRIR E/(arch_mov);
		ABRIR E/(arch_prod);
		ABRIR /S(arch_stock_sal);

		leerA();
		leerB();

		cant_prod:=0;
		cant_prod_retiro:=0;

		MIENTRAS (s.id_prod <> HV) o (m.id_prod <> HV) HACER:
			reg.id_prod:= s.id_prod;
			LEER(arch_prod,reg);
			SI EXISTE ENTONCES:
				SI (s.id_prod < m.id_prod) ENTONCES:
					s_sal:= s;
					GRABAR(arch_stock_sal,s_sal);
					leerA();
				SINO
					SI (s.id_prod = m.id_prod) ENTONCES:

						MIENTRAS (s.id_prod = m.id_prod) HACER:
							SI (s.stock > m.cant) ENTONCES:
								SI (m.tipo = "C") ENTONCES:
									s_sal.stock:= s.stock - m.cant; //compra
								SINO	
									s_sal.stock:= s.stock + m.cant; //devolución 
								Fin_si;
							SINO
								cant_prod:= cant_prod + 1; //cuento los productos que no se pueden vender por falta de stock
								Esc("Nombre del producto:",reg.nomb);
								Esc("Rubro:",reg.rubro);
							Fin_si;
							leerB();
						Fin_mientras;
						GRABAR(arch_stock_sal,s_sal);
						leerA();

					SINO
						s_sal.id_prod:= m.id_prod;
						s_sal.stock:= s.stock;

						MIENTRAS (s.id_prod = m.id_prod) HACER:
							SI (s.stock > m.cant) ENTONCES:
								SI (m.tipo = "C") ENTONCES:
									s_sal.stock:= s.stock - m.cant; //compra
								SINO	
									s_sal.stock:= s.stock + m.cant; //devolución 
								Fin_si;
							SINO
								cant_prod:= cant_prod + 1; //cuento los productos que no se pueden vender por falta de stock
								Esc("Nombre del producto:",reg.nomb);
								Esc("Rubro:",reg.rubro);
							Fin_si;
							leerB();
						Fin_mientras;
						GRABAR(arch_stock_sal,s_sal);
					Fin_si;
			SINO
				Esc("¡El producto no existe!");
			Fin_si;
		Fin_mientras;

		Cerrar(arch_stock);
		Cerrar(arch_mov);
		Cerrar(arch_prod);
		Cerrar(arch_stock_sal);
Fin_accion.