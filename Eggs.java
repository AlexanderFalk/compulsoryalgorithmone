
import java.lang.System;

public class Eggs {
	private int noOfFloors;
	private int secretFloor;
	private int flowBuilding;
	private int eggBreakFloor = 0;
	private int low;
	private int high;

	public Eggs(int noOfFloors, int secretFloor) {
		this.noOfFloors = noOfFloors;
		this.secretFloor = secretFloor;
		System.out.println("N = " + noOfFloors + "\nF = " + flowBuilding + "\n");
	}

	// low building
	public void findFloor() {
		low = 1;
		high = noOfFloors;
		while (high - low > 1) {
			flowBuilding = (low + high) / 2;
			if (flowBuilding >= secretFloor) {
				eggBreakFloor++;
				high = flowBuilding;
				System.out.println("HIGH : " + high);
			}
			else {
				low = flowBuilding;
			}
		}

		flowBuilding = high;

		System.out.println("BREAK: " + flowBuilding);
		System.out.println("Eggs used: " + eggBreakFloor);
	}

	public static void main(String[] args) {
		Eggs egg = new Eggs(1000, 5);
		egg.findFloor();
	}
}