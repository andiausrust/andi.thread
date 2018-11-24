package andi.com;

public class CalculatorTest {

    public static void main(String[] args) {
        Calculator calculator = new Calculator();

        System.out.println(calculator.add(20, 30));
        System.out.println(calculator.substract(20, 30));
        System.out.println(calculator.multiply(20, 30));
        System.out.println(calculator.divide(20, 30));

    }
}
