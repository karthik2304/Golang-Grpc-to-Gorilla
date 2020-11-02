package MethodOverloading;

public class NullTest1 {
	
	
	public static void method(Object obj){
	     System.out.println("method with param type - Object");
	   }
	 
	   public static void method(String obj){
	     System.out.println("method with param type - String");
	   }	   public static void method(StringBuffer strBuf){
		     System.out.println("method with param type - StringBuffer");
		   }
	   
	 
	   public static void main(String [] args){
	     method(null);
	   }
}
