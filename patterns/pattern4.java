package patterns;

public class pattern4 {
    static void patternInc(int n){
        for(int i=0;i<n;i++){
            for(int j=0;j<=i;j++){
                System.out.print(j+1+" ");
            }
            System.out.println();
        }
    }
    public static void main(String[] args) {
        patternInc(5);
    }
    
}
