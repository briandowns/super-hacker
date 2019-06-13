package java

// https://www.geeksforgeeks.org/rotate-bits-of-an-integer/

const RotateBits = `
class GFG  
{ 
	static final int INT_BITS = 32; 
	
	/*Function to left rotate n by d bits*/
	static int leftRotate(int n, int d) { 
		
		/* In n<<d, last d bits are 0.  
		To put first 3 bits of n at 
		last, do bitwise or of n<<d with 
		n >>(INT_BITS - d) */
		return (n << d) | (n >> (INT_BITS - d)); 
	} 
	
	/*Function to right rotate n by d bits*/
	static int rightRotate(int n, int d) { 
		
		/* In n>>d, first d bits are 0.  
		To put last 3 bits of at 
		first, do bitwise or of n>>d  
		with n <<(INT_BITS - d) */
		return (n >> d) | (n << (INT_BITS - d)); 
	} 
  
	// Driver code 
	public static void main(String arg[])  
	{ 
		int n = 16; 
		int d = 2; 
		System.out.print("Left Rotation of " + n + 
							" by " + d + " is "); 
		System.out.print(leftRotate(n, d)); 
		
		System.out.print("\nRight Rotation of " + n + 
								" by " + d + " is "); 
		System.out.print(rightRotate(n, d)); 
	} 
} `