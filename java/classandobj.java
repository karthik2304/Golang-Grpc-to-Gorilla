package ClassAndObjects;

public class Profile {

	
	private String name;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		
		this.name = name;
	}

	public static void main(String[] args) {

		Profile karthi = new Profile();
		karthi.setName("karthik");

		}
}