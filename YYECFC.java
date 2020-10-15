package cn;

public class YYECFC {
    double x1, x2;
    public double a, b, c;

    public YYECFC(double a, double b, double c) {
        this.a = a;
        this.b = b;
        this.c = c;
    }

    void real_root(double delt) {
        x1 = (-b + Math.sqrt((delt))) / (2 * a);
        x2 = (-b - Math.sqrt((delt))) / (2 * a);
        System.out.println("delt>=0,方程的实根为\nx1=" + x1);
        System.out.println("x2=" + x2);
    }

    void imag_root(double delt) {
        double real, imag;
        real = (-b) / (2 * a);
        imag = Math.sqrt((-delt));
        System.out.println("delt<0,方程的虚根为\nx1=" + real + "+" + imag + "i");
        System.out.println("x2=" + real + "-" + imag + "i");
    }

    void showRoot() {
        double delt = b * b - 4 * a * c;
        if (delt >= 0)
            real_root(delt);
        else
            imag_root(delt);
    }

    public static void main(String[] args) {
        YYECFC y1 ,y2;
        y1 = new YYECFC(1,5,10);
        y1.showRoot();
        y2 = new YYECFC(1,10,9);
        y2.showRoot();
    }
}
