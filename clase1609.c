#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main() {
    int arreglo[100];
    int i;

    // Inicializar el generador de números aleatorios
    srand(time(NULL));

    // Llenar el arreglo con 100 números aleatorios del 1 al 100 
    for (i = 0; i < 100; i++) {
        arreglo[i] = 1 + rand() % 100;
    }

    // Puedes imprimir el arreglo si deseas verificar los números generados

    for (i = 0; i < 100; i++) {
        printf("%d ", arreglo[i]);
    }
    printf("\n");


    return 0;
}
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

//funcion de carga
void cargar(int a[]) {
    //var local:
    int i;
    srand(time(NULL));

    // Llenar el arreglo con 100 números aleatorios del 1 al 100 
    for (i = 0; i < 100; i++) {
        a[i] = 1 + rand() % 100;
    }
}

int main() {
    //declarar variables
    int arreglo[100];
    int i;

    int maxPos = 0;
    int minPos = 0;

    // Llenar el arreglo con 100 números (debes cargarlo previamente)
    cargar(arreglo);
    int max = arreglo[0];
    int min = arreglo[0];
    for (i = 0; i < 100; i++) {
        printf("%d ", arreglo[i]);
    }
    printf("\n");

    // Encontrar el número mayor y menor y sus posiciones
    for (i = 0; i < 100; i++) {
        if (arreglo[i] > max) {
            max = arreglo[i];
            maxPos = i;
        }
        if (arreglo[i] < min) {
            min = arreglo[i];
            minPos = i;
        }
    }

    // Imprimir el número mayor y su posición
    printf("El número mayor es %d y se encuentra en la posición %d.\n", max, maxPos);

    // Imprimir el número menor y su posición
    printf("El número menor es %d y se encuentra en la posición %d.\n", min, minPos);

    return 0;
}
# include <stdio.h>
#include <stdlib.h>
#include <time.h>

//funcion de carga
void cargar(int a[]) {
    //var local:
    int i;
    srand(time(NULL));

    // Llenar el arreglo con 100 números aleatorios del 1 al 100 
    for (i = 0; i < 100; i++) {
        a[i] = 1 + rand() % 100;
    }
}

int main() {
    //declarar variables
    int arreglo[100];
    int i;
    int sumaPares = 0;
    int contadorPares = 0;

    // Llenar el arreglo con 100 números (debes cargarlo previamente)
    cargar(arreglo);

    for (i = 0; i < 100; i++) {
        printf("%d ", arreglo[i]);
    }
    printf("\n");
    // Contar y sumar los números pares
    for (i = 0; i < 100; i++) {
        if (arreglo[i] % 2 == 0) {
            contadorPares++; 
            sumaPares +=  arreglo[i];
        }
    }

    // Imprimir la cantidad de números pares
    printf("Cantidad de números pares: %d\n", contadorPares);

    // Imprimir la suma de los números pares
    printf("Suma de los números pares: %d\n", sumaPares);

    return 0;
}









#include <stdio.h>
//EJERCICIO 3.4
int main() {
    // Declaración de los arreglos A, B y C
    float A[3] = {1.0,2.0,3.0};
    float B[3] = {1.0, 2.5, 3.5};
    float C[6];

    // Índices para recorrer los arreglos A, B y C
    int i = 0;  // Para el arreglo A
    int j = 0;  // Para el arreglo B
    int k = 0;  // Para el arreglo C

    // Mezclar los arreglos A y B en C
    while (i < 3 && j < 3) {
        if (A[i] <= B[j]) {
            C[k] = A[i];
            i++;
        } else {
            C[k] = B[j];
            j++;
        }
        k++;
    }
    // Si quedan elementos en A, copiarlos a C
    while (i < 3) {
        C[k] = A[i];
        i++;
        k++;
    }

    // Si quedan elementos en B, copiarlos a C
    while (j < 3) {
        C[k] = B[j];
        j++;
        k++;
    }

    // Imprimir el arreglo C resultante (ordenado)
    printf("Arreglo C resultante (ordenado de forma creciente):\n");
    for (k = 0; k < 6; k++) {
        printf("%.2f ", C[k]);
    }
    printf("\n");

    return 0;
}
