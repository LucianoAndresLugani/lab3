package tp2

import "fmt"

type qalculadora struct {
	num1, num2, resultado float32
	op                    int16
}
type resultados struct {
	suma, resta, multiplicacion, division float32
}

func Suma(a float32, b float32) float32 {
	return a + b
}
func Resta(a float32, b float32) float32 {
	return a - b
}
func Multiplicacion(a float32, b float32) float32 {
	return a * b
}
func Division(a float32, b float32) float32 {
	return a / b
}

func Inicio(desv float32) {
	var valores qalculadora
	var resul resultados
	fmt.Println("ingrese numeros a trabajar")
	fmt.Scan(&valores.num1, &valores.num2)
	fmt.Println("seleccione operacion 1//suma  2//resta 3//multiplicacion 4//division")
	fmt.Scan(&valores.op)

	switch valores.op {
	case 1:
		resul.suma = Suma(valores.num1, valores.num2)
		fmt.Println(resul.suma + desv)

	case 2:
		resul.resta = Resta(valores.num1, valores.num2)
		fmt.Println(resul.resta - desv)

	case 3:
		resul.resta = Multiplicacion(valores.num1, valores.num2)
		fmt.Println(resul.multiplicacion * desv)

	case 4:
		resul.resta = Division(valores.num1, valores.num2)
		fmt.Println(resul.division * desv)

	default:
		fmt.Println("Opcion incorrecta")

	}

}
