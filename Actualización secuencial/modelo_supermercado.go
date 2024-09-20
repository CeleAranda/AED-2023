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
			SI (s.id_prod < m.id_prod) ENTONCES:
				s_sal:= s;
				GRABAR(arch_stock_sal,s_sal);
				leerA();
			SINO
				SI (s.id_prod = m.id_prod) ENTONCES:
					