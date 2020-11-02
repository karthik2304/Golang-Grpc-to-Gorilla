package exceptionHandling;

public class CheckedUncecked {

		static void m1() {
		throw new ArithmeticException();
	}
	static void m2() throws ClassNotFoundException {  
		throw new ClassNotFoundException();
	}
	public static void main(String[] args) {
		m1(); 			  
		try {
		  m2();
		}catch(ClassNotFoundException e) {
			System.out.println("Checked Exception handled.");
		}
	}
}
